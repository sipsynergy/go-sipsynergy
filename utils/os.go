package utils

import (
	"os"
	"os/signal"
)

// OnInterrupt returns a channel which receives the SIGINT system signal.
func OnInterrupt() chan os.Signal {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	return sigint
}
