package gophers

import (
	"os"
	"os/signal"
	"syscall"
)

func InstallShutdownReceiver(receiver func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for range c {
			receiver()
		}
	}()
}
