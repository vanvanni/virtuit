package api

import (
	"net/http"

	"github.com/vanvanni/virtuit/internal/api/deploy"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Spark is ready"))
	})
	mux.HandleFunc("/deploy", deploy.Handler)
}
