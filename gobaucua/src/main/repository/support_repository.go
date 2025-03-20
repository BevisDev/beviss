package repository

import "context"

type ISupportRepository interface {
	Get1Postgres(ctx context.Context) bool
}
