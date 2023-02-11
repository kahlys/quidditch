package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/justinas/alice"
	"github.com/rs/cors"
	"go.uber.org/zap"

	"github.com/kahlys/quidditch/backend"
	"github.com/kahlys/quidditch/backend/api"
	"github.com/kahlys/quidditch/backend/store"
)

var fAddr = flag.String("addr", ":8080", "listening address")

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		os.Exit(1)
	}

	// TODO: fix origins and remode debug
	c := cors.New(cors.Options{
		AllowCredentials: true,
	})

	db, err := store.NewDatabase("postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
	if err != nil {
		logger.Sugar().Fatalf("connecting to database: %v", err)
	}

	service := backend.NewService(logger, db)

	handler := api.Handler(logger, service)
	chain := alice.New(c.Handler, mwDebug(logger)).Then(handler)
	serv := http.Server{
		Addr:    *fAddr,
		Handler: chain,
	}

	logger.Sugar().Infow("Server listening", "addr", *fAddr)
	err = serv.ListenAndServe()
	logger.Sugar().Fatalw("Server stopped", "err", err.Error())
}

func mwDebug(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			logger.Sugar().Debugw("api", "method", req.Method, "url", req.URL)
			next.ServeHTTP(w, req)
		})
	}
}
