package service

import "context"

type ISupportService interface {
	PingDB(ctx context.Context) map[string]bool
	PingRedis(ctx context.Context) map[string]bool
}
