package network

import (
	"github.com/gin-gonic/gin"
	"mygrpcp_project/config"
	"mygrpcp_project/gRPC/client"
	"mygrpcp_project/service"
)

type Network struct {
	cfg        *config.Config
	service    *service.Service
	gRPCClient *client.GRPCClient
	engine     *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service, gRPCClient *client.GRPCClient) (*Network, error) {
	r := &Network{cfg: cfg, service: service, engine: gin.New(), gRPCClient: gRPCClient}

	// token 생성
	r.engine.POST("/login", r.login)
	// verify token
	r.engine.GET("/verify", r.verifyLogin(), r.verify)

	return r, nil
}

func (r *Network) Run() error {
	_ = r.engine.Run(":8080")
	return nil
}
