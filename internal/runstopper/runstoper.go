package runstopper

// Runner ...
type Runner interface {
	Run() error
}

// Stopper ...
type Stopper interface {
	Stop() error
}

// RunStopper ...
type RunStopper interface {
	Runner
	Stopper
}
