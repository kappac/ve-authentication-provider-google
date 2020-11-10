package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kappac/ve-authentication-provider-google/internal/connectionpool"
	"github.com/kappac/ve-authentication-provider-google/internal/grpcclient"
	"github.com/kappac/ve-authentication-provider-google/internal/logger"
	"github.com/kappac/ve-authentication-provider-google/pkg/client"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/context"
	"github.com/kappac/ve-authentication-provider-google/pkg/proto/request"
	"google.golang.org/grpc"
)

func createConnection(addr string, log logger.Logger) (client.VEAuthenticationProviderGoogleClient, error) {
	var err error

	client := client.New()
	if err = client.Dial(addr, grpc.WithInsecure()); err != nil {
		_ = log.Errorm("DialingFail", "err", err)
	}

	return client, err
}

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

	pool := connectionpool.New(
		connectionpool.WithConstructor(func() (grpcclient.Closer, error) {
			return createConnection(*addr, log)
		}),
	)
	if err := pool.Run(); err != nil {
		_ = log.Errorm("StartingPool", "err", err)
		os.Exit(1)
	}
	defer pool.Stop()

	con := pool.Pop().(client.VEAuthenticationProviderGoogleClient)
	begin := time.Now()
	ctx := context.New("", "", "")
	req := request.New(request.WithToken(*token))
	info, err := con.ValidateToken(ctx, req)
	pool.Push(con)

	fmt.Printf("Info: %v\n", info)
	_ = log.Infom("ValidateToken", "err", err, "took", time.Since(begin))
}
