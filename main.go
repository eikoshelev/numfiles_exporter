package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	yaml "gopkg.in/yaml.v2"
)

var (
	directory   = flag.String("directory", "", "directory with target subdirectories")
	targetsFile = flag.String("targets", "", "file with separate target directories")
	timeout     = flag.Int64("timeout", 15, "interval (seconds) of checking the number of files in the target directory")
	port        = flag.String("port", ":9095", "port of return of metrics")
	handler     = flag.String("handler", "/metrics", "handler for which metrics will be available")
	count       = flag.String("count", "all", "counting what to do: only files, only folders, everything (can take: files || dirs || all)")
)

type targetList struct {
	TargetDirectories []string `yaml:"targetDirectories"`
}

func checkFlags() {

	if *directory == "" && *targetsFile == "" {
		log.Fatalln("One or more flags (-directory, -targets) for target directories are not specified")
	} else if *directory != "" {
		openSubdirectories(*directory)
	} else if *targetsFile != "" {
		file, err := ioutil.ReadFile(*targetsFile)
		if err != nil {
			log.Fatalln("Failed read configuration file: ", err)
		}
		var tl targetList
		if err = yaml.Unmarshal(file, &tl); err != nil {
			log.Fatalln("Failed unmarshal targets file: ", err)
		}
		openDirectory(tl.TargetDirectories)
	}
}

func main() {

	flag.Parse()

	// prometheus handler
	http.Handle(*handler, promhttp.Handler())

	go func() {
		if err := http.ListenAndServe(*port, nil); err != nil {
			log.Fatalf("Failed to set http listener: %s", err)
			os.Exit(1)
		}
	}()

	log.Printf("Started work!\n")

	checkFlags()

	sgnl := make(chan os.Signal)
	signal.Notify(sgnl, os.Interrupt, os.Kill)
	<-sgnl
}
