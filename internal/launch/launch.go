package launch

import (
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/kappac/ve-authentication-provider-google/internal/logger"
	"github.com/kappac/ve-authentication-provider-google/internal/veservice"
)

var (
	log = logger.New(
		logger.WithEntity("Launch"),
	)
)

// Function defines an interface of code to be run in
// a routine.
type Function func(errCh chan<- error, exitCh chan error)

type closing chan interface{}

// Launch starts functions provided as the parameters and manages
// sygnals to interrupt its.
func Launch(fns ...Function) chan error {
	log.Debugm("initializing")

	fnsAmount := len(fns)

	var (
		once       = sync.Once{}
		exitChs    = make([]chan error, fnsAmount)
		resCh      = make(chan error)
		funcErrCh  = make(chan error)
		interuptCh = make(chan os.Signal)
		closing    = make(closing)
		stop       = func() {
			once.Do(func() {
				close(closing)
			})
		}
		err error
	)

	go interupt(interuptCh)

	for i, fn := range fns {
		exitCh := make(chan error)
		go fn(funcErrCh, exitCh)
		exitChs[i] = exitCh
	}

	go (func() {
		for {
			select {
			case err = <-funcErrCh:
				log.Debugm("FunctionError", "err", err)
				stop()
			case <-interuptCh:
				log.Debugm("Interupting")
				stop()
			case <-closing:
				log.Debugm("Stopping")

				for _, exitCh := range exitChs {
					exitCh <- errors.New("Stop service")
					err = <-exitCh
					log.Infom("FunctionStopped", "err", err)
				}

				resCh <- err
				return
			}
		}
	})()

	return resCh
}

// WithRunStopper creates a function suitable for Launch to
// deal with RunStopper.
func WithRunStopper(rs veservice.RunStopper) Function {
	return func(errCh chan<- error, exitCh chan error) {
		go (func() {
			errCh <- rs.Run()
		})()

		for {
			select {
			case <-exitCh:
				exitCh <- rs.Stop()
				return
			}
		}
	}
}

func interupt(interuptCh chan<- os.Signal) {
	signal.Notify(interuptCh, syscall.SIGINT, syscall.SIGTERM)
}
