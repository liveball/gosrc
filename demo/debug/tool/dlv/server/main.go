package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	addr := flag.String("addr", ":10080", "address to listen to")

	flag.Parse()

	if err := run(*addr); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run(addr string) error {
	ctx := context.Background()

	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	metrics := NewMetrics()

	h := &MainHandler{
		metrics: metrics,
	}
	mux := http.NewServeMux()
	mux.Handle("/", h)

	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	errs := make(chan error, 1)
	go func() {
		metrics.CountS("server.started")
		log.Printf("server listening addr=%s\n", srv.Addr)
		errs <- srv.ListenAndServe()
	}()

	var err error
	select {
	case sig := <-sigs:
		log.Printf("caught signal %s, shutting down\n", sig)

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		err = srv.Shutdown(ctx)
	case err = <-errs:
	}

	return err
}

// MainHandler is a fake service handler.
type MainHandler struct {
	metrics *Metrics
}

func (h *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.metrics.CountS("server.req-incoming")
	defer func(start time.Time) {
		log.Printf("request processed in %s\n", time.Now().Sub(start))
	}(time.Now())

	fmt.Fprintf(w, "OK")
}
