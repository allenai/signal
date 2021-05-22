// Package signal offers a utility for handling signal interrupts.
package signal

import (
	"context"
	"os"
	"os/signal"
)

// WithSignal returns a cancelable context which ends on signal interrupt.
//
// The first interrupt cancels the context, allowing callers to terminate
// gracefully. Upon receiving a second interrupt the process is terminated with
// exit code 130 (128 + SIGINT)
func WithSignal(parent context.Context) (context.Context, context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	ctx, cancel := context.WithCancel(parent)

	// In most cases this routine will leak due to the lack of a second signal.
	// That's OK since this is expected to last for the life of the process.
	go func() {
		select {
		case <-sigChan:
			cancel()
		case <-ctx.Done():
			// Do nothing.
		}
		<-sigChan
		os.Exit(130)
	}()

	return ctx, func() {
		signal.Stop(sigChan)
		cancel()
	}
}
