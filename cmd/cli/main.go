package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/kappac/ve-authentication-provider-google/internal/google"
	"github.com/kappac/ve-authentication-provider-google/internal/pb"
	"google.golang.org/grpc"
)

type client struct {
	context.Context
	pb.VEAuthProviderGoogleServiceClient
	log.Logger
}

// new returns an AddService that's backed by the provided ClientConn.
func new(ctx context.Context, cc *grpc.ClientConn, logger log.Logger) client {
	return client{ctx, pb.NewVEAuthProviderGoogleServiceClient(cc), logger}
}

func (c client) validateToken(t string) *google.Token {
	req := &pb.VEValidateTokenRequest{
		Token: t,
	}

	reply, err := c.VEAuthProviderGoogleServiceClient.ValidateToken(c.Context, req)
	if err != nil {
		c.Logger.Log(err)
		return nil
	}

	if reply.Error != nil {
		c.Logger.Log(reply.Error)
		return nil
	}

	return &google.Token{
		FullName:   reply.Info.FullName,
		GivenName:  reply.Info.GivenName,
		FamilyName: reply.Info.FamilyName,
		Picture:    reply.Info.Picture,
		Email:      reply.Info.Email,
	}
}

func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		addr  = fs.String("addr", ":8000", "Address for gRPC server")
		token = fs.String("token", "token", "Token to be validated")
	)
	flag.Usage = fs.Usage // only show our flags
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	// logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	// logger = log.NewContext(logger).With("transport", *transport)

	root := context.Background()

	cc, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		_ = logger.Log("err", err)
		os.Exit(1)
	}
	defer cc.Close()
	svc := new(root, cc, logger)

	req := &pb.VEValidateTokenRequest{
		Token: *token,
	}
	begin := time.Now()
	_, err = svc.VEAuthProviderGoogleServiceClient.ValidateToken(root, req)

	_ = logger.Log("method", "ValidateToken", "took", time.Since(begin))
}
