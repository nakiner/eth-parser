package main

import (
	"context"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/nakiner/eth-parser/internal/app/v1"
	ethClient "github.com/nakiner/eth-parser/internal/client/etherium"
	"github.com/nakiner/eth-parser/internal/client/inmemory"
	blockRepo "github.com/nakiner/eth-parser/internal/domain/block/repository"
	subsRepo "github.com/nakiner/eth-parser/internal/domain/subscriber/repository"
	transactionsRepo "github.com/nakiner/eth-parser/internal/domain/transaction/repository"
	"github.com/nakiner/eth-parser/internal/logger"
	"github.com/nakiner/eth-parser/internal/observer"
	"github.com/nakiner/eth-parser/internal/provider/cache"
	ethProvider "github.com/nakiner/eth-parser/internal/provider/etherium"
	"github.com/nakiner/eth-parser/internal/server"
	"github.com/nakiner/eth-parser/internal/service/transaction"
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

	memoryStoreMetadata := inmemory.New()
	transactionStore := inmemory.New()
	cacheStoreProvider := cache.NewCacheProvider(memoryStoreMetadata, memoryStoreMetadata, transactionStore)

	blockRepository := blockRepo.NewRepository(cacheStoreProvider)
	subsRepository := subsRepo.NewRepository(cacheStoreProvider)
	transactionsRepository := transactionsRepo.NewRepository(cacheStoreProvider)

	ethC := ethClient.NewClient()
	ethP := ethProvider.NewProvider(ethC)

	ethService := transaction.NewService(blockRepository, subsRepository, transactionsRepository)

	appImpl := v1.NewService(ethService)
	mustInit(pb.RegisterETHParserServiceV1HandlerServer(ctx, mux, appImpl))

	observerImpl := observer.NewObserver(ethService, ethP)
	observerActor, observerCloseFunc := observerImpl.Serve(ctx, 2)
	app.AddActor(observerActor, observerCloseFunc)

	app.SetHandler(mux)
	app.Use(middleware.Recoverer)
	app.Use(middleware.Logger)
}

func mustInit(err error) {
	if err != nil {
		logger.FatalKV(context.Background(), "init failure", "err", err)
	}
}
