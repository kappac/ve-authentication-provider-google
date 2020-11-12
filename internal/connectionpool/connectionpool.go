package connectionpool

import (
	"sync"

	"github.com/kappac/ve-authentication-provider-google/internal/grpcclient"
	"github.com/kappac/ve-authentication-provider-google/internal/logger"
)

const (
	defaultMin = 2
	defaultMax = 5
)

// ConnectionConstructor creates new instance of a connection.
type ConnectionConstructor func() (grpcclient.Closer, error)

// ConnectionPool provides an interface to manage connections.
type ConnectionPool interface {
	Push(grpcclient.Closer)
	Pop() grpcclient.Closer
	Run() error
	Stop() error
}

type connections []grpcclient.Closer

type connectionPoper chan grpcclient.Closer

type connectionPusher chan grpcclient.Closer

type closing chan interface{}

type connectionPool struct {
	wg          sync.WaitGroup
	once        sync.Once
	min, max    int
	constructor ConnectionConstructor
	connections connections
	pop         connectionPoper
	push        connectionPusher
	closing     closing
	logger      logger.Logger
	err         error
}

// New creates new ConnectionPool instance
func New(opts ...PoolOption) ConnectionPool {
	pop := make(connectionPoper)
	push := make(connectionPusher)
	cl := make(closing)
	l := logger.New(
		logger.WithEntity("ConnectionPool"),
	)
	p := &connectionPool{
		logger:  l,
		min:     defaultMin,
		max:     defaultMax,
		pop:     pop,
		push:    push,
		closing: cl,
	}

	for _, opt := range opts {
		opt(p)
	}

	p.connections = make(connections, 0)

	return p
}

func (p *connectionPool) Pop() grpcclient.Closer {
	p.logger.Debugm("Poping")
	return <-p.pop
}

func (p *connectionPool) Push(c grpcclient.Closer) {
	p.logger.Debugm("Pushing")
	p.push <- c
}

func (p *connectionPool) Run() error {
	p.wg.Add(1)

	go p.listen()
	p.validateOpts()

	if err := p.initConnections(); err != nil {
		p.stop()
		return err
	}

	p.logger.Infom("Running", "min", p.min, "max", p.max)

	return nil
}

func (p *connectionPool) Stop() error {
	p.stop()
	return p.err
}

func (p *connectionPool) stop() {
	p.stopListen()
	p.wg.Wait()
}

func (p *connectionPool) closeChannels() {
	close(p.pop)
	close(p.push)
}

func (p *connectionPool) validateOpts() {
	if p.min > p.max {
		p.max = p.min
	}
}

func (p *connectionPool) initConnections() error {
	var err error

	if p.constructor == nil {
		return errorNoNoConstructor
	}

	for i := 0; i < p.min; i++ {
		if err = p.initConnection(); err != nil {
			return err
		}
	}

	return nil
}

func (p *connectionPool) closeConnections() {
	for idx, c := range p.connections {
		p.err = c.Close()
		p.logger.Debugm("closing", "idx", idx, "err", p.err)
	}
}

func (p *connectionPool) checkConnections() {
	cc := len(p.connections)
	if cc == 0 && cc < p.max {
		p.initConnection()
	}
}

func (p *connectionPool) initConnection() error {
	c, err := p.constructor()
	p.logger.Debugm("initConnection", "err", err)

	if err == nil {
		p.push <- c
	}

	return err
}

func (p *connectionPool) listen() {
	for {
		var (
			first grpcclient.Closer
			pop   connectionPoper
		)
		cc := len(p.connections)
		if cc > 0 {
			first = p.connections[0]
			pop = p.pop
		}

		p.logger.Debugm("listen", "conn_len", len(p.connections))

		select {
		case c, ok := <-p.push:
			p.logger.Debugm("listen:p.push", "ok", ok, "c", c)
			if ok {
				p.connections = append(p.connections, c)
			}
		case pop <- first:
			p.connections = p.connections[1:]
			p.checkConnections()
		case <-p.closing:
			p.closeChannels()
			p.closeConnections()
			p.wg.Done()
			return
		}
	}
}

func (p *connectionPool) stopListen() {
	p.once.Do(func() {
		close(p.closing)
	})
}
