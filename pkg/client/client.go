package client

import (
	"context"
	"time"

	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/providerinfo"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/request"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/response"
	"github.com/kappac/ve-back-end-utils/pkg/grpcclient"
	"google.golang.org/grpc"
)

const (
	dialTimeout = 5 * time.Second
)

// VEAuthenticationProviderGoogleClient is a client interface
// for VEAuthenticationProviderGoogle service
type VEAuthenticationProviderGoogleClient interface {
	Dial() error
	Close() error
	ValidateToken(c context.Context, r request.VEValidateTokenRequest) (providerinfo.VEProviderInfo, error)
}

type veAuthenticationProviderGoogleClient struct {
	gc       grpcclient.GrpcClient
	service  pb.VEServiceClient
	addr     string
	grpcOpts []grpc.DialOption
}

func newClient(addr string, opts []grpc.DialOption) VEAuthenticationProviderGoogleClient {
	return &veAuthenticationProviderGoogleClient{
		gc:       grpcclient.New(),
		grpcOpts: opts,
		addr:     addr,
	}
}

func (c *veAuthenticationProviderGoogleClient) Dial() error {
	if err := c.gc.Dial(c.addr, c.grpcOpts...); err != nil {
		return err
	}
	c.service = pb.NewVEServiceClient(c.gc.GetClientConn())

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
