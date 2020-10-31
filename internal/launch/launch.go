package launch

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/kappac/ve-authentication-provider-google/internal/logger"
	"github.com/kappac/ve-authentication-provider-google/internal/types/runstopper"
)

var (
	log = logger.New(
		logger.WithEntity("Launch"),
	)
)

// Function defines an interface of code to be run in
// a routine.
type Function func(errCh chan<- error, exitCh chan error)

// Launch starts functions provided as the parameters and manages
// sygnals to interrupt its.
func Launch(fns ...Function) chan error {
	log.Debugm("initializing")

	fnsAmount := len(fns)

	var (
		exitChs    = make([]chan error, fnsAmount)
		resCh      = make(chan error)
		funcErrCh  = make(chan error)
		interuptCh = make(chan os.Signal)
		stopCh     = make(chan bool)
		stop       = func() {
			stopCh <- true
		}
		isClosing bool
		err       error
	)

	go interupt(interuptCh)

	for i, fn := range fns {
		exitCh := make(chan error)
		go fn(funcErrCh, exitCh)
		exitChs[i] = exitCh
	}

	go (func() {
		for !isClosing {
			select {
			case err = <-funcErrCh:
				log.Debugm("FunctionError", "err", err)
				go stop()
			case <-interuptCh:
				log.Debugm("Interupting")
				go stop()
			case <-stopCh:
				log.Debugm("Stopping")
				for _, exitCh := range exitChs {
					exitCh <- errors.New("Stop service")
					err := <-exitCh
					log.Infom("FunctionStopped", "err", err)
				}

				isClosing = true

				resCh <- err
			}
		}
	})()

	return resCh
}

// WithRunStopper creates a function suitable for Launch to
// deal with RunStopper.
func WithRunStopper(rs runstopper.RunStopper) Function {
	return func(errCh chan<- error, exitCh chan error) {
		var (
			isClosing bool
		)

		go (func() {
			errCh <- rs.Run()
		})()

		for !isClosing {
			select {
			case <-exitCh:
				isClosing = true
				exitCh <- rs.Stop()
			}
		}
	}
}

func interupt(interuptCh chan<- os.Signal) {
	signal.Notify(interuptCh, syscall.SIGINT, syscall.SIGTERM)
}
