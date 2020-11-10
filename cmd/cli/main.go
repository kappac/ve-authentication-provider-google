package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kappac/ve-authentication-provider-google/internal/logger"
	"github.com/kappac/ve-authentication-provider-google/pkg/client"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/context"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/request"
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

	client := client.New()
	if err := client.Dial(*addr, grpc.WithInsecure()); err != nil {
		_ = log.Errorm("DialingFail", "err", err)
		os.Exit(1)
	}
	defer client.Close()

	begin := time.Now()
	ctx := context.New("", "", "")
	req := request.New(request.WithToken(*token))
	info, err := client.ValidateToken(ctx, req)

	fmt.Printf("Info: %v\n", info)
	_ = log.Infom("ValidateToken", "err", err, "took", time.Since(begin))
}
