package main

import (
	"context"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/nakiner/eth-parser/internal/app/v1"
	"github.com/nakiner/eth-parser/internal/logger"
	"github.com/nakiner/eth-parser/internal/server"
	pb "github.com/nakiner/eth-parser/pkg/pb/eth_parser/v1"
)

func main() {
	app := server.New()
	initApp(app)
	mustInit(app.Run())
}

func initApp(app *server.App) {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	appImpl := v1.NewService()
	mustInit(pb.RegisterETHParserServiceV1HandlerServer(ctx, mux, appImpl))

	app.SetHandler(mux)
	app.Use(middleware.Recoverer)
	app.Use(middleware.Logger)
}

func mustInit(err error) {
	if err != nil {
		logger.FatalKV(context.Background(), "init failure", "err", err)
	}
}
