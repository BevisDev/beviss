package impl

import (
	"context"
	"github.com/BevisDev/backend-template/redis"
	"goauth/src/main/repository"
	"goauth/src/main/service"
)

type PingServiceImpl struct {
	pingRepository repository.IPingRepository
	redisClient    *redis.RedisClient
}

func NewPingServiceImpl(
	pingRepository repository.IPingRepository,
) service.IPingService {
	return &PingServiceImpl{
		pingRepository: pingRepository,
	}
}

func (impl *PingServiceImpl) PingDB(ctx context.Context) map[string]bool {
	return map[string]bool{
		"Schema1": impl.pingRepository.Get1MSSQL(ctx),
		"Schema2": impl.pingRepository.Get1Orc(ctx),
	}
}

func (impl *PingServiceImpl) PingRedis(ctx context.Context) map[string]bool {
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
