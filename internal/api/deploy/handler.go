package deploy

import (
	"fmt"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/vanvanni/virtuit/internal/spark"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Oink!")
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("config")
	if err != nil {
		http.Error(w, "Missing TOML file in 'config' field", http.StatusBadRequest)
		return
	}
	defer file.Close()

	var parsedCfg spark.Spark
	if _, err := toml.NewDecoder(file).Decode(&parsedCfg); err != nil {
		http.Error(w, "Invalid TOML: "+err.Error(), http.StatusBadRequest)
		return
	}

	cfg, err := spark.CreateSpark(&parsedCfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = spark.Deploy(*cfg)
	if err != nil {
		http.Error(w, "Failed to deploy Spark: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Spark deployed successfully"))
}
