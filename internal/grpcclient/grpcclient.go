package grpcclient

import (
	"time"

	"google.golang.org/grpc"
)

const (
	dialTimeout = 5 * time.Second
)

// Dialer provides an interface to Dial
type Dialer interface {
	Dial(addr string, opts ...grpc.DialOption) error
}

// Closer provides an interface to Close
type Closer interface {
	Close() error
}

// DialCloser combines Dialer and Closer
type DialCloser interface {
	Dialer
	Closer
}

// GrpcClient is a wrapper for grpc.ClientConn
type GrpcClient interface {
	Dial(addr string, opts ...grpc.DialOption) error
	Close() error
	GetClientConn() *grpc.ClientConn
}

type grpcClient struct {
	client *grpc.ClientConn
}

// New constructs new GrpcClient instance
func New() GrpcClient {
	return &grpcClient{}
}

func (c *grpcClient) Dial(addr string, opts ...grpc.DialOption) error {
	opts = append(
		opts,
		grpc.WithBlock(),
		grpc.WithTimeout(dialTimeout),
		grpc.WithDefaultCallOptions(
			grpc.WaitForReady(false),
		),
	)
	cc, err := grpc.Dial(addr, opts...)

	if err != nil {
		return err
	}

	c.client = cc

	return nil
}

func (c *grpcClient) Close() error {
	return c.client.Close()
}

func (c *grpcClient) GetClientConn() *grpc.ClientConn {
	return c.client
}
