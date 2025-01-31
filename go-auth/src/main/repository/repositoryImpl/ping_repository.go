package repositoryImpl

import (
	"context"
	"goauth/src/main/infrastructure/database"
	"goauth/src/main/repository"
)

type PingRepositoryImpl struct {
}

func NewPingRepositoryImpl() repository.IPingRepository {
	return &PingRepositoryImpl{}
}

func (p PingRepositoryImpl) Get1MSSQL(ctx context.Context, schema string) bool {
	var result int
	database.GetUsingNamed(ctx, &result, schema, "SELECT 1", nil)
	return result == 1
}

func (p PingRepositoryImpl) Get1Orc(ctx context.Context, schema string) bool {
	var result int
	database.GetUsingNamed(ctx, &result, schema, "SELECT 1 FROM DUAL", nil)
	return result == 1
}
