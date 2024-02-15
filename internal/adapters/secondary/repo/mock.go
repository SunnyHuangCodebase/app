package repo

import "context"

type MockRepo struct {
	store map[string]string
}

func (mock MockRepo) GetAllData(ctx context.Context) map[string]string {
	return mock.store
}
func (mock MockRepo) GetDataByID(ctx context.Context, id string) string {
	return mock.store[id]
}

func NewMockRepo() MockRepo {
	return MockRepo{
		store: map[string]string{
			"1": "data1string",
			"2": "data2string",
			"3": "data3string",
		},
	}
}
