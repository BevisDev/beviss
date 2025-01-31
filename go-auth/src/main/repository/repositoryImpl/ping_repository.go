package repositoryImpl

import (
	"context"
	"goauth/src/main/global"
	"goauth/src/main/repository"
)

type PingRepositoryImpl struct {
}

func NewPingRepositoryImpl() repository.IPingRepository {
	return &PingRepositoryImpl{}
}

func (p PingRepositoryImpl) Get1MSSQL(ctx context.Context) bool {
	var result int
	global.AuthDB.GetUsingNamed(ctx, &result, "SELECT 1", nil)
	return result == 1
}

func (p PingRepositoryImpl) Get1Orc(ctx context.Context) bool {
	var result int
	global.AuthDB.GetUsingNamed(ctx, &result, "SELECT 1 FROM DUAL", nil)
	return result == 1
}
