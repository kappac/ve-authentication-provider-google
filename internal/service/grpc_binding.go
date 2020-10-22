package service

import (
	"context"

	"github.com/kappac/ve-authentication-provider-google/internal/google"
	"github.com/kappac/ve-authentication-provider-google/internal/pb"
)

type GrpcBinding struct {
	VEAuthenticationProviderGoogle
}

func (b GrpcBinding) ValidateToken(ctx context.Context, req *pb.VEValidateTokenRequest) (*pb.VEValidateTokenResponse, error) {
	var (
		token *google.Token
		err   error
	)

	if token, err = b.VEAuthenticationProviderGoogle.ValidateToken(req.Token); err != nil {
		res := &pb.VEValidateTokenResponse{
			Request: &pb.VEValidateTokenRequest{
				Token: req.Token,
			},
			Error: &pb.VEError{
				Code:        100,
				Description: err.Error(),
			},
		}

		return res, nil
	}

	res := &pb.VEValidateTokenResponse{
		Info: &pb.VEProviderInfo{
			FullName:   token.FullName,
			GivenName:  token.GivenName,
			FamilyName: token.FamilyName,
			Picture:    token.Picture,
			Email:      token.Email,
		},
		Request: &pb.VEValidateTokenRequest{
			Token: req.Token,
		},
	}

	return res, nil
}
