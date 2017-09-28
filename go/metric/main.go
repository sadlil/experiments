package main

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var counter = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace: "application",
	Subsystem: "count",
	Name: "counter_five_second",
	Help: "counter help text",
})

func main() {
	prometheus.MustRegister(counter)

	go run()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8989", nil)
}

func run() {
	for {
		time.Sleep(time.Second*5)
		counter.Inc()
	}
}