package repository

import "context"

type IPingRepository interface {
	Get1MSSQL(ctx context.Context) bool
	Get1Orc(ctx context.Context) bool
}
