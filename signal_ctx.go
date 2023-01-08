package sigctx

import (
	"context"
	"os/signal"
	"syscall"
)

// NewShutdownContext returns a new context.Context that is canceled when the
// process receives a SIGINT or SIGTERM signal. Call the context.CancelFunc
// to restore default signal handling behavior.
func NewShutdownContext() (context.Context, context.CancelFunc) {
	return DeriveShutdownContext(context.Background())
}

// DeriveShutdownContext returns a context.Context that is canceled when the
// process receives a SIGINT or SIGTERM signal. Call the context.CancelFunc
// to restore default signal handling behavior.
func DeriveShutdownContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
}
