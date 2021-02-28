package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	contador = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minha_aplicacao",
		Help: "Número de eventos processados",
	})
)

func main() {
	testeMetricas()
	http.Handle("/metricas", promhttp.Handler())
	http.ListenAndServe(":9099", nil)
}

func testeMetricas() {
	go func() {
		for {
			contador.Inc()
			time.Sleep(1 * time.Second)
		}
	}()
}
