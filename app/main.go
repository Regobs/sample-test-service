// sample-test-service — a deliberately tiny HTTP service used to exercise the
// DevLift deployment pipeline end to end. It has no dependencies, no config and
// no secrets: the point is to produce a real service repo that the deployment
// workflow can open a workflow-YAML PR against.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"service": "sample-test-service",
			"env":     os.Getenv("ENVIRONMENT"),
			"time":    time.Now().UTC().Format(time.RFC3339),
		})
	})

	log.Printf("sample-test-service listening on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
