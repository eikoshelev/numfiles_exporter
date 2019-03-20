package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	numberOfFiles = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "number_of_files",
		Help: "Number of files in the target directory",
	}, []string{"directory", "counting"})
)

func init() {
	prometheus.Register(numberOfFiles)
}
