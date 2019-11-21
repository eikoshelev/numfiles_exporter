package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	yaml "gopkg.in/yaml.v2"
)

type parameters struct {
	Directory, TargetsFile, Port, Handler, Count string
	Timeout                                      time.Duration
}

type targetList struct {
	TargetDirectories []string `yaml:"targetDirectories"`
}

func (p *parameters) checkFlags(timer string) {

	var err error

	if timer == "" {
		log.Printf("Flag '-timeout' is empty: set default timeout - 15s\n")
		p.Timeout = time.Duration(15 * time.Second)
	} else {
		p.Timeout, err = time.ParseDuration(timer)
		if err != nil {
			log.Printf("Failed parse value of flag '-timeout': %s, set default timeout - 15s\n", err.Error())
			p.Timeout = time.Duration(15 * time.Second)
		}
	}

	if p.Directory == "" && p.TargetsFile == "" {
		log.Fatalln("One or more flags (-directory, -targets) for target directories are not specified")
	} else if p.Directory != "" {
		p.openSubdirectories()
	} else if p.TargetsFile != "" {
		file, err := ioutil.ReadFile(p.TargetsFile)
		if err != nil {
			log.Fatalf("Failed read configuration file: %s", err.Error())
		}
		var tl targetList
		if err = yaml.Unmarshal(file, &tl); err != nil {
			log.Fatalf("Failed unmarshal targets file: %s", err.Error())
		}
		p.openDirectory(tl.TargetDirectories)
	}
}

func main() {

	var (
		params parameters
		timer  string
	)

	flag.StringVar(&params.Directory, "directory", "", "directory with target subdirectories")
	flag.StringVar(&params.TargetsFile, "targets", "", "file with separate target directories")
	flag.StringVar(&timer, "timeout", "", "interval (seconds) of checking the number of files in the target directory")
	flag.StringVar(&params.Port, "port", ":9095", "port of return of metrics")
	flag.StringVar(&params.Handler, "handler", "/metrics", "handler for which metrics will be available")
	flag.StringVar(&params.Count, "count", "all", "counting what to do: only files, only folders, everything (can take: files || dirs || all)")

	flag.Parse()

	// prometheus handler
	http.Handle(params.Handler, promhttp.Handler())

	go func() {
		if err := http.ListenAndServe(params.Port, nil); err != nil {
			log.Fatalf("Failed to set http listener: %s\n", err.Error())
		}
	}()

	log.Printf("Started work!\n")

	params.checkFlags(timer)

	sgnl := make(chan os.Signal)
	signal.Notify(sgnl, os.Interrupt, os.Kill)
	<-sgnl
}
