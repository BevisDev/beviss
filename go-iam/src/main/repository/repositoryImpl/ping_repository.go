package repositoryImpl

import (
	"context"
	"goiam/src/main/global"
	"goiam/src/main/repository"
)

type PingRepositoryImpl struct {
}

func NewPingRepositoryImpl() repository.IPingRepository {
	return &PingRepositoryImpl{}
}

func (p *PingRepositoryImpl) Get1Postgres(ctx context.Context) bool {
	var result int
	global.AuthDB.GetUsingNamed(ctx, &result, "SELECT 1", nil)
	return result == 1
}
