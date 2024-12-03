package bookingjob

import (
	"context"
	"fmt"
	"http-server/internal/shared"
	"sync"

	"go.uber.org/fx"
)

type BookingJob struct {
	deps *shared.SharedDeps
	wg   *sync.WaitGroup
	ctx  context.Context
}

func NewBookingJob(deps *shared.SharedDeps, lc fx.Lifecycle) *BookingJob {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			cancel()
			wg.Wait()
			fmt.Println("all service is stopped")
			return nil
		},
	})
	return &BookingJob{
		deps: deps,
		wg:   &wg,
		ctx:  ctx,
	}
}

var BookingJobModule = fx.Options(shared.SharedModuleFx, fx.Provide(NewBookingJob))
