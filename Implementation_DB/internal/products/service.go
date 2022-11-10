package products

import (
	"context"

	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/internal/domains"
)

type Service interface {
	GetByName(ctx context.Context, name string) (domains.Product, error)
	Store(ctx context.Context, p domains.Product) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetByName(ctx context.Context, name string) (domains.Product, error) {
	return s.repository.GetByName(ctx, name)
}

func (s *service) Store(ctx context.Context, p domains.Product) (int, error) {
	return s.repository.Store(ctx, p)
}
