package main

import (
	"app/internal/adapters/primary/api"
	"app/internal/adapters/secondary/repo"
	"app/internal/core/services/data"
)

func main() {
	dataAPI := repo.NewMockRepo()
	service := data.New(dataAPI)
	app := api.NewApp(service)
	app.Start()
}
