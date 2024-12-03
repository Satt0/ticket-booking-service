package outboxrelay

import (
	"context"
	"http-server/internal/shared"
	"time"

	"go.uber.org/fx"
)

type OutboxRelay struct {
	deps *shared.SharedDeps
	ctx  *context.Context
}

func NewOutboxRelay(deps *shared.SharedDeps, lc fx.Lifecycle) *OutboxRelay {
	ctx, cancel := context.WithCancel(context.Background())
	lc.Append(fx.StopHook(func() {
		cancel()
	}))
	return &OutboxRelay{
		deps: deps,
		ctx:  &ctx,
	}
}

func (s *OutboxRelay) Run() {
	ticket := time.Tick(100 * time.Millisecond)
	for {
		<-ticket
		s.sendRelayedKafkaMessage()
	}
}
