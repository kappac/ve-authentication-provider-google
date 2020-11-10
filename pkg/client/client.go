package client

import (
	"context"
	"time"

	"github.com/kappac/ve-authentication-provider-google/internal/grpcclient"
	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/request"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/response"
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
	ValidateToken(c context.Context, r request.VEValidateTokenRequest) (providerinfo.VEProviderInfo, error)
}

type veAuthenticationProviderGoogleClient struct {
	gc      grpcclient.GrpcClient
	service pb.VEAuthProviderGoogleServiceClient
}

func newClient() VEAuthenticationProviderGoogleClient {
	return &veAuthenticationProviderGoogleClient{
		gc: grpcclient.New(),
	}
}

func (c *veAuthenticationProviderGoogleClient) Dial(addr string, opts ...grpc.DialOption) error {
	if err := c.gc.Dial(addr, opts...); err != nil {
		return err
	}
	c.service = pb.NewVEAuthProviderGoogleServiceClient(c.gc.GetClientConn())

	return nil
}

func (c *veAuthenticationProviderGoogleClient) Close() error {
	return c.gc.Close()
}

func (c *veAuthenticationProviderGoogleClient) ValidateToken(ctx context.Context, r request.VEValidateTokenRequest) (providerinfo.VEProviderInfo, error) {
	if err := r.Verify(); err != nil {
		return nil, err
	}

	req, err := r.Marshal()
	if err != nil {
		return nil, err
	}

	resp, err := c.service.ValidateToken(ctx, req.(*pb.VEValidateTokenRequest))
	if err != nil {
		return nil, err
	}

	veresp := response.New()
	if err := veresp.Unmarshal(resp); err != nil {
		return nil, err
	}

	return veresp.GetInfo(), veresp.GetError()
}
