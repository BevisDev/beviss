package impl

import (
	"context"
	"github.com/BevisDev/backend-template/redis"
	"gobaucua/src/main/repository"
	"gobaucua/src/main/service"
)

type SupportServiceImpl struct {
	repo        repository.ISupportRepository
	redisClient *redis.RedisCache
}

func NewSupportServiceImpl(
	repo repository.ISupportRepository,
) service.ISupportService {
	return &SupportServiceImpl{
		repo: repo,
	}
}

func (impl *SupportServiceImpl) PingDB(ctx context.Context) map[string]bool {
	return map[string]bool{
		"Database": impl.repo.Get1Postgres(ctx),
	}
}

func (impl *SupportServiceImpl) PingRedis(ctx context.Context) map[string]bool {
	var resp = make(map[string]bool)
	if err := impl.redisClient.Set(ctx, "key1", 1, 10); err != nil {
		resp["Redis"] = false
		return resp
	}
	var rs int
	if err := impl.redisClient.Get(ctx, "key1", &rs); err != nil {
		resp["Redis"] = false
		return resp
	}
	resp["Redis"] = true
	return resp
}
