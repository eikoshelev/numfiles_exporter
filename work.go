package main

import (
	"io/ioutil"
	"log"
	"path"
	"time"
)

func openDirectory(directories []string) {

	log.Printf("Target directories:\n")

	for _, dir := range directories {
		log.Printf("\"%s\"\n", dir)
		go checkNumFiles(dir)
	}

}

func openSubdirectories(directory string) {

	log.Printf("Target subdirectories in \"%s\":\n", directory)

	dir, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatalf("Failed read directory: %s\n", err)
	}

	for _, item := range dir {
		if item.IsDir() {
			log.Printf("\"%s\"\n", item.Name())
			go checkNumFiles(path.Join(directory, item.Name()))
		}
	}
}

func checkNumFiles(directory string) {

	for {
		dir, err := ioutil.ReadDir(directory)
		if err != nil {
			log.Fatalf("Failed read directory: %s\n", err)
		}
		numberOfFiles.WithLabelValues(directory).Set(float64(len(dir)))
		time.Sleep(time.Duration(*timeout) * time.Second)
	}
}
