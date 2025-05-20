package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vanvanni/virtuit/internal/api"
	"github.com/vanvanni/virtuit/internal/out"
)

var directories = []string{
	"/etc/virtuit",
	"/var/log/virtuit",
	"/var/virtuit",
	"/var/virtuit/cells",
	"/var/virtuit/kernels",
}

func ensureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// >> Create Dirs
	for _, dir := range directories {
		if err := ensureDir(dir); err != nil {
			log.Fatal(err)
		} else {
			out.Logger.Info("Ready " + dir)
		}
	}

	// >> Start API
	mux := http.NewServeMux()
	api.RegisterRoutes(mux)
	out.Logger.Info("VirtuIT controller running on :1338")
	log.Fatal(http.ListenAndServe(":1338", mux))
}
