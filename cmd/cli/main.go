package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kappac/ve-authentication-provider-google/pkg/client"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/request"
	"github.com/kappac/ve-back-end-utils/pkg/logger"
	vecontext "github.com/kappac/ve-back-end-utils/pkg/proto/context"
	"google.golang.org/grpc"
)

func main() {
	log := logger.New(logger.WithEntity("Client"))
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		addr  = fs.String("addr", ":8000", "Address for gRPC server")
		token = fs.String("token", "token", "Token to be validated")
	)

	flag.Usage = fs.Usage
	if err := fs.Parse(os.Args[1:]); err != nil {
		_ = log.Errorm("ParsingArgs", "err", err)
		os.Exit(1)
	}

	cl := client.New(
		client.WithAddress(*addr),
		client.WithGRPCDialOptions(
			grpc.WithInsecure(),
		),
	)
	if err := cl.Dial(); err != nil {
		_ = log.Errorm("DialingFail", "err", err)
		os.Exit(1)
	}
	defer cl.Close()

	begin := time.Now()
	ctx := vecontext.New(context.Background(), "", "", "")
	req := request.New(request.WithToken(*token))
	info, err := cl.ValidateToken(ctx, req)

	fmt.Printf("Info: %v\n", info)
	_ = log.Infom("ValidateToken", "err", err, "took", time.Since(begin))
}
