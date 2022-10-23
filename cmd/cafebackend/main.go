package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/edgedb/edgedb-go"
	"github.com/gin-gonic/gin"
	userhandler "github.com/relipocere/cafebackend/internal/business/user-handler"
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

	log := mustInitLogger(true)
	defer closerCleanup(log)

	mustInitViper(log)

	edgeClient := mustCreateDBClient(ctx, log)
	closerAdd(edgeClient.Close)

	userRepo := user.NewRepo()

	userHandler := userhandler.NewHandler(edgeClient, userRepo)

	resolver := graph.NewResolver(userHandler)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(resolver))
	srv.SetErrorPresenter(middleware.ErrorHandlerMw(log))

	router := gin.Default()
	router.Use(middleware.AuthenticationMW(edgeClient, userRepo))

	router.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	port := viper.GetString("server.port")
	log.Infof("starting server at port %s", port)

	err := router.Run(port)
	if err != nil {
		log.Errorf("can't start server: %v", err)
		return
	}
}

func mustCreateDBClient(ctx context.Context, log *zap.SugaredLogger) *edgedb.Client {
	usr := viper.Get("edgedb.user")
	password := viper.Get("edgedb.password")
	host := viper.Get("edgedb.host")
	port := viper.Get("edgedb.port")
	db := viper.Get("edgedb.database")

	dsn := fmt.Sprintf("edgedb://%s:%s@%s:%s/%s?tls_security=insecure",
		usr,
		password,
		host,
		port,
		db,
	)

	client, err := edgedb.CreateClientDSN(ctx, dsn, edgedb.Options{
		ConnectTimeout: time.Second,
	})
	if err != nil {
		log.Fatalf("connecting to db: %v", err)
	}

	return client
}

func mustInitViper(log *zap.SugaredLogger) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.SetDefault("server.port", ":9000")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("can't read secrets: %v", err)
	}
}

var closers []func() error

func closerAdd(fn func() error) {
	closers = append(closers, fn)
}

func closerCleanup(log *zap.SugaredLogger) {
	log.Debug("Cleaning up")
	for _, fn := range closers {
		err := fn()
		if err != nil {
			log.Errorf("Clean up: %v", err)
		}
	}
	log.Sync()
}

func mustInitLogger(dev bool) *zap.SugaredLogger {
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

	return logger.Sugar()
}
