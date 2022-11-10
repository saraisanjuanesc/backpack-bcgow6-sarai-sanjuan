package products

import (
	"context"

	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/internal/domains"
)

type Service interface {
	GetByName(ctx context.Context, name string) (domains.Product, error)
	Store(ctx context.Context, p domains.Product) (int, error)
	GetAll(ctx context.Context) ([]domains.Product, error)
	DeleteS(ctx context.Context, id int64) error
	Update(ctx context.Context, id int, name, ptype string, count int, price float64) (domains.Product, error)
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

func (s *service) GetAll(ctx context.Context) ([]domains.Product, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) DeleteS(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}

func (s *service) Update(ctx context.Context, id int, name, ptype string, count int, price float64) (domains.Product, error) {
	return s.repository.Update(ctx, id, name, ptype, count, price)
}
