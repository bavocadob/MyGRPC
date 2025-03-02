package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"mygrpcp_project/config"
	"mygrpcp_project/gRPC/paseto"
	auth "mygrpcp_project/gRPC/proto"
	"time"
)

type GRPCClient struct {
	client      *grpc.ClientConn
	authClient  auth.AuthServiceClient
	pasetoMaker *paseto.PasetoMaker
}

func NewClient(cfg *config.Config) (*GRPCClient, error) {
	c := new(GRPCClient)

	if client, err := grpc.Dial(cfg.GRPC.URL, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	} else {
		c.client = client
		c.authClient = auth.NewAuthServiceClient(c.client)
		c.pasetoMaker = paseto.NewPasetoMaker(cfg)
	}
	return c, nil
}

func (g *GRPCClient) CreateAuth(address string) (*auth.AuthData, error) {
	log.Println("CreateAuth {}", address)
	now := time.Now()
	expired := now.Add(time.Hour * 24 * 30)

	a := &auth.AuthData{
		Name:       address,
		CreateDate: now.Unix(),
		ExpireDate: expired.Unix(),
	}

	if token, err := g.pasetoMaker.CreateNewToken(a); err != nil {
		return nil, err
	} else {
		a.Token = token

		if res, err := g.authClient.CreateAuth(context.Background(), &auth.CreateTokenReq{Auth: a}); err != nil {
			return nil, err
		} else {
			return res.Auth, nil
		}

	}

}
func (g *GRPCClient) VerifyAuth(token string) (*auth.Verify, error) {
	if res, err := g.authClient.VerifyAuth(context.Background(), &auth.VerifyTokenReq{Token: token}); err != nil {
		return nil, err
	} else {
		return res.V, nil
	}
}
