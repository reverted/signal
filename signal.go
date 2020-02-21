package signal

import (
	"os"
	"os/signal"
	"syscall"
)

func Wait() chan os.Signal {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	return sigs
}
