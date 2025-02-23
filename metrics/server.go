package metrics

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Metrics scraped at %s", time.Now().Format(time.RFC3339))
		promhttp.Handler().ServeHTTP(w, r)
	})
	log.Println("Starting Prometheus metrics server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
