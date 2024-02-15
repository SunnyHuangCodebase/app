package api

import (
	"app/internal/core/services/data"
	"fmt"
	"io"
	"net/http"
)

type App struct {
	server  *http.Server
	service data.Service
}

func (a App) Start() {
	err := a.server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}

func NewApp(service data.Service) App {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("GET /data", GetData(service))
	mux.HandleFunc("GET /data/{id}", GetDataByID(service))

	server := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	return App{
		server:  server,
		service: service,
	}

}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, HTTP!\n")
}

func GetData(service data.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		apiData := fmt.Sprintf("%s", service.Repo.GetAllData(r.Context()))
		io.WriteString(w, apiData)
	}
}

func GetDataByID(service data.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		io.WriteString(w, service.Repo.GetDataByID(r.Context(), id))
	}
}
