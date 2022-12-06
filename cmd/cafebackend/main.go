package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	producthandler "github.com/relipocere/cafebackend/internal/business/product-handler"
	reviewhandler "github.com/relipocere/cafebackend/internal/business/review-handler"
	storehandler "github.com/relipocere/cafebackend/internal/business/store-handler"
	userhandler "github.com/relipocere/cafebackend/internal/business/user-handler"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/database/image"
	"github.com/relipocere/cafebackend/internal/database/product"
	"github.com/relipocere/cafebackend/internal/database/review"
	"github.com/relipocere/cafebackend/internal/database/store"
	"github.com/relipocere/cafebackend/internal/database/user"
	"github.com/relipocere/cafebackend/internal/graph"
	"github.com/relipocere/cafebackend/internal/graph/generated"
	"github.com/relipocere/cafebackend/internal/graph/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	ctx := context.Background()
	defer closerCleanup()

	mustInitLogger(true)
	mustInitViper()
	mustInitFilesDirecotry()

	pgxPool := mustCreateDBClient(ctx)
	closerAdd(func() error {
		pgxPool.Close()
		return nil
	})

	filesDir := viper.GetString("server.files_dir")

	userRepo := user.NewRepo()
	storeRepo := store.NewRepo()
	imageRepo := image.NewRepo()
	productRepo := product.NewRepo()
	reviewRepo := review.NewRepo()

	masterNode := database.NewPGX(pgxPool)

	userHandler := userhandler.NewHandler(masterNode, userRepo)
	storeHandler := storehandler.NewHandler(masterNode, storeRepo)
	productHandler := producthandler.NewHandler(masterNode, productRepo, storeRepo)
	reviewHandler := reviewhandler.NewHandler(masterNode, storeRepo, reviewRepo)

	resolver := graph.NewResolver(
		filesDir,
		masterNode,
		imageRepo,
		userHandler,
		storeHandler,
		productHandler,
		reviewHandler,
	)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(resolver))
	srv.SetErrorPresenter(middleware.ErrorHandlerMw())

	router := gin.Default()
	router.Static("/assets", filesDir)
	router.Use(middleware.AuthenticationMW(masterNode, userRepo))

	router.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	port := viper.GetString("server.port")
	zap.S().Infof("starting server at port %s", port)

	err := router.Run(port)
	if err != nil {
		zap.S().Errorf("can't start server: %v", err)
		return
	}
}

func mustCreateDBClient(ctx context.Context) *pgxpool.Pool {
	usr := viper.Get("psql.user")
	password := viper.Get("psql.password")
	host := viper.Get("psql.host")
	port := viper.Get("psql.port")
	db := viper.Get("psql.database")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		usr,
		password,
		host,
		port,
		db,
	)

	zap.S().Debugw("establishing db connection", "dsn", dsn)

	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		zap.S().Fatalf("connecting to db: %v", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		zap.S().Fatalf("db ping: %v", err)
	}

	return conn
}

func mustInitViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.SetDefault("server.port", ":9000")

	err := viper.ReadInConfig()
	if err != nil {
		zap.S().Fatalf("can't read secrets: %v", err)
	}
}

var closers []func() error

func closerAdd(fn func() error) {
	closers = append(closers, fn)
}

func closerCleanup() {
	for _, fn := range closers {
		err := fn()
		if err != nil {
			fmt.Printf("clean up: %v", err)
		}
	}
}

func mustInitLogger(dev bool) {
	cfg := &zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "msg",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalColorLevelEncoder,
			TimeKey:     "tt",
			EncodeTime:  zapcore.RFC3339TimeEncoder,
		},
	}

	if dev {
		cfg.Encoding = "console"
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		cfg.Development = true
	}

	logger, err := cfg.Build()
	if err != nil {
		fmt.Printf("can't initialize logger: %v\n", err)
		os.Exit(1)
	}

	zap.ReplaceGlobals(logger)

	closerAdd(func() error {
		return logger.Sync()
	})
}

func mustInitFilesDirecotry() {
	err := os.MkdirAll(viper.GetString("server.files_dir"), fs.ModePerm)
	if err != nil {
		zap.S().Fatalf("can't create files directory: %v", err)
	}
}
