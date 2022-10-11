package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/edgedb/edgedb-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	ctx := context.Background()

	log := mustInitLogger(true)
	defer closerCleanup(log)

	client := mustCreateDBClient(ctx, log)
	closerAdd(client.Close)

}

func mustCreateDBClient(ctx context.Context, log *zap.SugaredLogger) *edgedb.Client {
	user := mustGetEnv(log, "EDGEDB_USER")
	password := mustGetEnv(log, "EDGEDB_PASSWORD")
	host := mustGetEnv(log, "EDGEDB_HOST")
	port := mustGetEnv(log, "PORT")
	db := mustGetEnv(log, "EDGEDB_DATABASE")
	certPath := mustGetEnv(log, "EDGEDB_CERT_PATH")

	dsn := fmt.Sprintf("edgedb://%s:%s@%s:%s/%s?tls_ca_file=%s",
		user,
		password,
		host,
		port,
		db,
		certPath,
	)

	client, err := edgedb.CreateClientDSN(ctx, dsn, edgedb.Options{
		ConnectTimeout: time.Second,
	})
	if err != nil {
		log.Fatalf("connecting to db: %v", err)
	}

	return client
}

func mustGetEnv(log *zap.SugaredLogger, key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("env '%s' is not set", key)
	}

	return val
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
