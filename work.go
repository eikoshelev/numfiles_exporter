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
		go checkNumber(dir)
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
			go checkNumber(path.Join(directory, item.Name()))
		}
	}
}

func checkNumber(directory string) {

	switch *count {
	case "all":

		log.Printf("\"%s\"\n", directory)
		for {
			dir, err := ioutil.ReadDir(directory)
			if err != nil {
				log.Fatalf("Failed read directory: %s\n", err)
			}
			numberOfFiles.WithLabelValues(directory, *count).Set(float64(len(dir)))
			time.Sleep(time.Duration(*timeout) * time.Second)
		}
	case "files":

		log.Printf("\"%s\"\n", directory)
		for {
			dir, err := ioutil.ReadDir(directory)
			if err != nil {
				log.Fatalf("Failed read directory: %s\n", err)
			}
			filesCount := 0
			for _, item := range dir {
				if !item.IsDir() {
					filesCount++
				}
			}
			numberOfFiles.WithLabelValues(directory, *count).Set(float64(filesCount))
			time.Sleep(time.Duration(*timeout) * time.Second)
		}
	case "dirs":

		log.Printf("\"%s\"\n", directory)
		for {
			dir, err := ioutil.ReadDir(directory)
			if err != nil {
				log.Fatalf("Failed read directory: %s\n", err)
			}
			dirsCount := 0
			for _, item := range dir {
				if item.IsDir() {
					dirsCount++
				}
			}
			numberOfFiles.WithLabelValues(directory, *count).Set(float64(dirsCount))
			time.Sleep(time.Duration(*timeout) * time.Second)
		}
	default:
		log.Fatalf("Not valid value of the \"-count\" flag - \"%s\": should be \"files\", \"dirs\" or \"all\" - default.", *count)
	}
}
