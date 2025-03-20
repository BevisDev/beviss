package repositoryImpl

import (
	"context"
	"gobaucua/src/main/lib"
	"gobaucua/src/main/repository"
)

type SupportRepositoryImpl struct {
}

func NewSupportRepositoryImpl() repository.ISupportRepository {
	return &SupportRepositoryImpl{}
}

func (p *SupportRepositoryImpl) Get1Postgres(ctx context.Context) bool {
	var result int
	lib.AuthDB.GetOne(ctx, &result, "SELECT 1", nil)
	return result == 1
}
