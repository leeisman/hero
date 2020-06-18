package graceful


import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Launch service gracefully, or force exit after two consecutive signals
func Launch(start func() error, shutdown func() error, timeout time.Duration) {
	done := make(chan struct{})
	go func() {
		if err := start(); err != nil {
			fmt.Println(err)
		}
		done <- struct{}{}
	}()
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}...)

	select {
	case <-done:
		if shutdown != nil {
			if err := shutdownWithTimeout(shutdown, timeout); err != nil {
				fmt.Println(err)
			}
			return
		}
	case s := <-ch:
		fmt.Printf("received signal %s: terminating\n", s)
		if shutdown != nil {
			if err := shutdownWithTimeout(shutdown, timeout); err != nil {
				fmt.Println(err)
			}
			return
		}
	}
}

func shutdownWithTimeout(shutdown func() error, timeout time.Duration) error {
	done := make(chan error)
	go func() {
		done <- shutdown()
	}()
	select {
	case <-time.After(timeout):
		return errors.New("timeout reached: terminating")
	case err := <-done:
		return err
	}
}
