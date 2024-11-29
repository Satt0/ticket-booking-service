package crawlerblockchain

import (
	"context"
	"fmt"
	"http-server/internal/shared"
	"sync"

	"go.uber.org/fx"
)

type CrawlerBlockchain struct {
	deps             *shared.SharedDeps
	wg               *sync.WaitGroup
	ctx              context.Context
	producerNotiChan chan int64
}

func NewCrawlerBlockchain(deps *shared.SharedDeps, lc fx.Lifecycle) *CrawlerBlockchain {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("All service started")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Waiting for all service to end")
			cancel()
			wg.Wait()
			fmt.Println("Waiting for all service to end-DONE")
			return nil
		},
	})
	return &CrawlerBlockchain{
		deps:             deps,
		wg:               &wg,
		ctx:              ctx,
		producerNotiChan: make(chan int64),
	}
}

var CrawlerBlockchainModule = fx.Options(shared.SharedModuleFx, fx.Provide(NewCrawlerBlockchain))
