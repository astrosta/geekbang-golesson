package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"

	"golang.org/x/sync/errgroup"
)

func main() {
	group, ctx := errgroup.WithContext(context.Background())
	mux := new(http.ServeMux)
	mux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "pong")
	})
	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	group.Go(func() error {
		exit := make(chan os.Signal)
		signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-exit:
			return errors.Errorf("received signal: %v", sig)
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "ctx canceled")
		}
	})

	group.Go(func() error {
		select {
		case <-ctx.Done():
			//设置http服务关闭超时，15s
			timeoutCtx, _ := context.WithTimeout(context.Background(), 15*time.Second)
			return srv.Shutdown(timeoutCtx)
		}
	})

	group.Go(func() error {
		return srv.ListenAndServe()
	})

	if err := group.Wait(); err != nil {
		fmt.Printf("%+v", err)
	}

	return
}
