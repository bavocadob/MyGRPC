package service

import (
	"mygrpcp_project/config"
	auth "mygrpcp_project/gRPC/proto"
	"mygrpcp_project/repository"
)

type Service struct {
	cfg *config.Config

	repo *repository.Repository
}

func NewService(cfg *config.Config, repository *repository.Repository) (*Service, error) {
	r := &Service{cfg: cfg, repo: repository}
	return r, nil
}

func (s *Service) CreateAuth(name string) (*auth.AuthData, error) {
	return s.repo.CreateAuth(name)
}
