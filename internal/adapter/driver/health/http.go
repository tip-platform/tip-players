// Package health provides lightweight health/readiness endpoints.
package health

import (
	"context"
	"net/http"
	"time"
)

// ReadyFunc returns nil when the service is ready to receive traffic.
// It should be fast and respect the context deadline.
type ReadyFunc func(context.Context) error

// StartHTTP starts an HTTP server exposing /healthz (liveness) and /readyz (readiness).
//
// - /healthz: must be cheap and always returns 200 if the process is alive.
// - /readyz: calls ready(ctx). If it returns error, responds 503.
func StartHTTP(addr string, ready ReadyFunc) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if ready == nil {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("ready"))
			return
		}

		ctx := r.Context()
		if _, hasDeadline := ctx.Deadline(); !hasDeadline {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, 2*time.Second)
			defer cancel()
		}

		if err := ready(ctx); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("not ready"))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	})

	srv := &http.Server{Addr: addr, Handler: mux}

	go func() {
		_ = srv.ListenAndServe()
	}()

	return srv
}
