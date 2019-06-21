package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Define the metrics we wish to expose
var fooMetric = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "foo_metric", Help: "Shows whether a foo has occurred in our cluster"})

var barMetric = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "bar_metric", Help: "Shows whether a bar has occurred in our cluster"})

func init() {
	//Register metrics with prometheus
	prometheus.MustRegister(fooMetric)
	prometheus.MustRegister(barMetric)

	//Set fooMetric to 54321
	fooMetric.Set(54321)

	//Set barMetric to 12345
	barMetric.Set(12345)
}

func main() {
	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port :8086")
	log.Fatal(http.ListenAndServe(":8086", nil))
}
