package checkpayment

import (
	"context"
	"fmt"
	"http-server/internal/shared"
	"sync"

	"go.uber.org/fx"
)

type CheckOutdatedPayment struct {
	deps *shared.SharedDeps
	wg   *sync.WaitGroup
	ctx  context.Context
}

func NewCheckOutdatedPayment(deps *shared.SharedDeps, lc fx.Lifecycle) *CheckOutdatedPayment {
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
	return &CheckOutdatedPayment{
		deps: deps,
		wg:   &wg,
		ctx:  ctx,
	}
}

var CheckOutDatePaymentModule = fx.Options(shared.SharedModuleFx, fx.Provide(NewCheckOutdatedPayment))
