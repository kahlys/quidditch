package main

import (
	"flag"
	"net/http"
	"os"

	"go.uber.org/zap"

	"github.com/kahlys/quidditch/backend/api"
)

var fAddr = flag.String("addr", ":80", "listening address")

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		os.Exit(1)
	}

	serv := http.Server{
		Addr:    *fAddr,
		Handler: api.Handler(),
	}

	logger.Sugar().Infow("Server listening", "addr", *fAddr)
	err = serv.ListenAndServe()
	logger.Sugar().Fatalw("Server stopped", "err", err.Error())
}
