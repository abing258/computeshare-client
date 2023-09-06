// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"computeshare-client/internal/biz"
	"computeshare-client/internal/conf"
	"computeshare-client/internal/data"
	"computeshare-client/internal/server"
	"computeshare-client/internal/service"
	"computeshare-client/third_party/p2p"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	ipfsNode, err := p2p.RunDaemon()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	p2pService := service.NewP2pService(ipfsNode)
	client, err := service.NewDockerCli()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	vmService := service.NewVmService(client, logger)
	computepowerService := service.NewComputepowerService(ipfsNode, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, p2pService, vmService, computepowerService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
