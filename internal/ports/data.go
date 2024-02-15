package ports

import "context"

type DataRepository interface {
	GetAllData(ctx context.Context) map[string]string
	GetDataByID(ctx context.Context, id string) string
}
