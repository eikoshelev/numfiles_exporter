package main

import (
	"io/ioutil"
	"log"
	"path"
	"time"
)

func (p *parameters) openDirectory(directories []string) {

	log.Printf("Target directories:\n")

	for _, dir := range directories {
		go p.checkNumber(dir)
	}
}

func (p *parameters) openSubdirectories() {

	log.Printf("Target subdirectories in '%s':\n", p.Directory)

	dir, err := ioutil.ReadDir(p.Directory)
	if err != nil {
		log.Fatalf("Failed read directory: %s\n", err.Error())
	}

	for _, item := range dir {
		if item.IsDir() {
			go p.checkNumber(path.Join(p.Directory, item.Name()))
		}
	}
}

func (p *parameters) checkNumber(directory string) {

	switch p.Count {
	case "all":

		log.Printf("'%s'\n", directory)
		for {
			dir, err := ioutil.ReadDir(directory)
			if err != nil {
				log.Fatalf("Failed read directory: %s\n", err.Error())
			}
			numberOfFiles.WithLabelValues(directory, p.Count).Set(float64(len(dir)))
			time.Sleep(p.Timeout)
		}
	case "files":

		log.Printf("'%s'\n", directory)
		for {
			dir, err := ioutil.ReadDir(directory)
			if err != nil {
				log.Fatalf("Failed read directory: %s\n", err.Error())
			}
			filesCount := 0
			for _, item := range dir {
				if !item.IsDir() {
					filesCount++
				}
			}
			numberOfFiles.WithLabelValues(directory, p.Count).Set(float64(filesCount))
			time.Sleep(p.Timeout)
		}
	case "dirs":

		log.Printf("'%s'\n", directory)
		for {
			dir, err := ioutil.ReadDir(directory)
			if err != nil {
				log.Fatalf("Failed read directory: %s\n", err.Error())
			}
			dirsCount := 0
			for _, item := range dir {
				if item.IsDir() {
					dirsCount++
				}
			}
			numberOfFiles.WithLabelValues(directory, p.Count).Set(float64(dirsCount))
			time.Sleep(p.Timeout)
		}
	default:
		log.Fatalf("Not valid value of the '-count' flag - '%s': should be 'files', 'dirs' or 'all' - default.", p.Count)
	}
}
