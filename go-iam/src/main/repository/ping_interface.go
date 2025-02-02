package repository

import "context"

type IPingRepository interface {
	Get1Postgres(ctx context.Context) bool
}
