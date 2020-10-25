package launch

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kappac/ve-authentication-provider-google/internal/logger"
)

var (
	log = logger.New(
		logger.WithEntity("Launch"),
	)
)

// Function defines an interface of code to be run in
// a routine.
type Function func(errCh chan<- error, exitCh chan bool)

// Launch starts functions provided as the parameters and manages
// sygnals to interrupt its.
func Launch(fns ...Function) chan error {
	log.Debugm("initializing")

	var (
		exitChs []chan bool

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

	for _, fn := range fns {
		exitCh := make(chan bool)
		go fn(funcErrCh, exitCh)
		exitChs = append(exitChs, exitCh)
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
					exitCh <- true

					select {
					case <-exitCh:
					}
				}

				isClosing = true

				resCh <- err
			}
		}
	})()

	return resCh
}

func interupt(interuptCh chan<- os.Signal) {
	signal.Notify(interuptCh, syscall.SIGINT, syscall.SIGTERM)
}
