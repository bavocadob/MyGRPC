package cmd

import (
	"mygrpcp_project/config"
	"mygrpcp_project/gRPC/client"
	"mygrpcp_project/network"
	"mygrpcp_project/repository"
	"mygrpcp_project/service"
)

type App struct {
	cfg *config.Config

	gRPCClient *client.GRPCClient
	service    *service.Service
	repository *repository.Repository
	network    *network.Network
}

func NewApp(cfg *config.Config) *App {
	a := &App{cfg: cfg}

	var err error

	if a.gRPCClient, err = client.NewClient(cfg); err != nil {
		panic(err)
	} else if a.repository, err = repository.NewRepository(cfg, a.gRPCClient); err != nil {
		panic(err)
	} else if a.service, err = service.NewService(cfg, a.repository); err != nil {
		panic(err)
	} else if a.network, err = network.NewNetwork(cfg, a.service, a.gRPCClient); err != nil {
		panic(err)
	} else {
		_ = a.network.Run()
	}

	return a
}
