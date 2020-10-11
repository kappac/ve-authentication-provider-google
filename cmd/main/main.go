package main

import (
	"github.com/kappac/ve-authentication-provider-google/internal/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	svr := service.GetService()

	if err := svr.Run(); err != nil {
		logger.Fatal(err)
	}
}
