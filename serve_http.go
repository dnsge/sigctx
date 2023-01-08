package sigctx

import (
	"context"
	"log"
	"net/http"
	"time"
)

var (
	// ShutdownGracePeriod specifies the duration to wait while shutting down
	// the HTTP server gracefully before forcibly exiting
	ShutdownGracePeriod = time.Second * 10
)

// ListenAndServe starts a http.Server and waits until a SIGINT or SIGTERM is
// received. Then, the http.Server will be shutdown gracefully.
func ListenAndServe(server *http.Server) {
	ctx, done := NewShutdownContext()
	defer done()

	failed := false
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			failed = true
			log.Printf("Error while listening and serving: %s\n", err.Error())
			done()
		}
	}()

	// Wait until we receive a signal
	<-ctx.Done()

	if failed {
		// We failed to listen, bail out and don't do any shutdown tasks
		return
	}

	// Restore default signal handling
	done()
	log.Println("Shutting down gracefully, press Ctrl+C to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), ShutdownGracePeriod)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Printf("Error while shutting down server: %s\n", err.Error())
	}
}
