package client

import (
	"context"
	"time"

	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-authentication-provider-google/pkg/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/pkg/request"
	"github.com/kappac/ve-authentication-provider-google/pkg/response"
	"google.golang.org/grpc"
)

const (
	dialTimeout = 5 * time.Second
)

// VEAuthenticationProviderGoogleClient is a client interface
// for VEAuthenticationProviderGoogle service
type VEAuthenticationProviderGoogleClient interface {
	Dial(addr string, opts ...grpc.DialOption) error
	Close() error
	ValidateToken(request.VEValidateTokenRequest) (providerinfo.VEProviderInfo, error)
}

type veAuthenticationProviderGoogleClient struct {
	client  *grpc.ClientConn
	service pb.VEAuthProviderGoogleServiceClient
	context context.Context
}

// New constructs new VEAuthenticationProviderGoogleClient instance
func New() VEAuthenticationProviderGoogleClient {
	return &veAuthenticationProviderGoogleClient{
		context: context.TODO(),
	}
}

func (c *veAuthenticationProviderGoogleClient) Dial(addr string, opts ...grpc.DialOption) error {
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
	c.service = pb.NewVEAuthProviderGoogleServiceClient(c.client)

	return nil
}

func (c *veAuthenticationProviderGoogleClient) Close() error {
	return c.client.Close()
}

func (c *veAuthenticationProviderGoogleClient) ValidateToken(r request.VEValidateTokenRequest) (providerinfo.VEProviderInfo, error) {
	if err := r.Verify(); err != nil {
		return nil, err
	}

	req, err := r.Marshal()
	if err != nil {
		return nil, err
	}

	resp, err := c.service.ValidateToken(c.context, req.(*pb.VEValidateTokenRequest))
	if err != nil {
		return nil, err
	}

	veresp := response.New()
	if err := veresp.Unmarshal(resp); err != nil {
		return nil, err
	}

	return veresp.GetInfo(), veresp.GetError()
}
