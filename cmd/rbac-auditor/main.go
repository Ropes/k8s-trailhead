package main

import (
	"os"
	"runtime"
	"time"

	log "github.com/Sirupsen/logrus"
)

func main() {
	log.Info("starting GOMAXPROCS smoke test")

	e, ok := os.LookupEnv("GOMAXPROCS")
	if !ok {
		log.Warn("GOMAXPROCS env not defined")
	} else {
		log.Infof("GOMAXPROCS defined: %q", e)
	}

	t := time.NewTicker(time.Second * 1)

	for {
		select {
		case <-t.C:
			log.WithFields(log.Fields{
				"NumCPU":          runtime.NumCPU(),
				"GOMAXPROCS(-1)":  runtime.GOMAXPROCS(-1),
				"GOMAXPROCS(env)": os.Getenv("GOMAXPROCS"),
			}).Info("runtime stats")
		}
	}
}
