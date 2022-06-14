package main

import (
	"net/http"

	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	mq "gitlab.com/oyerickshaw/site-reliability/mqtt-exporter/mqtt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Load env variables from .env in DEV environment
	if os.Getenv("GO_ENV") == "DEV" {
		err := godotenv.Load(".env")

		if err != nil {
			fmt.Println("Error loading .env file", err)
		}
	}

	// Register Gauge
	for _, topic := range mq.Topics {
		if topic.Gauge != nil {
			prometheus.MustRegister(topic.Gauge)
			fmt.Printf("\n Registered gauge %s", topic.Topic)
		}
	}

	// MQTT Connections
	mq.Init()

	// log.Fatal(http.ListenAndServe(":8026", nil))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to mqtt-exporter. Made by Oye! Rickshaw energy team."))
	})
	r.Get("/v2/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "#HELP The version of the broker.\n broker_version %s \n", mq.Version)
	})
	// Enable prometheus metrics endpoint
	r.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8026", r)
}
