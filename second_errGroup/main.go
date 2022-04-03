package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	egroup, errCtx := errgroup.WithContext(ctx)
	egroup.Go(func() error {
		return http.ListenAndServe(":8080", nil)
	})
	egroup.Go(func() error {
		<-ctx.Done()
		fmt.Println("server shutdown")
		return errors.New("server shutdown")
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	egroup.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-c:
				cancel()
			}
		}
	})
	if err := egroup.Wait(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("main exit")
}
