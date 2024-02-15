package data

import (
	"app/internal/ports"
	"context"
)

type API interface {
	GetAllData(ctx context.Context) map[string]string
	GetDataByID(ctx context.Context, id string) string
}

type Service struct {
	Repo ports.DataRepository
}

func New(repo ports.DataRepository) Service {
	return Service{
		Repo: repo,
	}
}
