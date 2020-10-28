package service

import (
	"context"

	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"github.com/kappac/ve-authentication-provider-google/internal/types/runstopper"
	"github.com/kappac/ve-authentication-provider-google/pkg/request"
)

type grpcBinding struct {
	runstopper.RunStopper

	svc authProviderGoogle
}

func (b grpcBinding) ValidateToken(ctx context.Context, req *pb.VEValidateTokenRequest) (*pb.VEValidateTokenResponse, error) {
	veRequest := request.New()
	if err := veRequest.Unmarshal(req); err != nil {
		return nil, err
	}

	if err := veRequest.Verify(); err != nil {
		return nil, err
	}

	resp, _ := b.svc.ValidateToken(veRequest)
	resppb, pberr := resp.Marshal()
	if pberr != nil {
		return nil, pberr
	}

	return resppb.(*pb.VEValidateTokenResponse), nil
}

func (b grpcBinding) Run() error {
	return b.svc.Run()
}

func (b grpcBinding) Stop() error {
	return b.svc.Stop()
}
