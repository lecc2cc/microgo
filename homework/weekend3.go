package homework

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
	xerrors "golang.org/x/xerrors"
)

/**
 *
 */
func main() {
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	shutdownCh := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("shutdown ..."))
		shutdownCh <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8081",
	}
	// server group
	g.Go(func() error {
		return server.ListenAndServe()
	})

	// shutdown channel group
	g.Go(func() error {
		var ee error
		select {
		case <-ctx.Done():
			if ctx.Err() != nil {
				ee = xerrors.Errorf("errgroup done err: %v", ee)
			}
		case <-shutdownCh:
			ee = fmt.Errorf("shutdown ch")
		}
		msg := ""
		if ee != nil {
			msg = ee.Error()
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Println("errgroup done/shutdown ch...", msg)
		return server.Shutdown(timeoutCtx)
	})

	// signal group
	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return xerrors.Errorf("signal: %v", sig)
		}
	})

	fmt.Printf("errgroup wait: %+v\n", g.Wait())
}
