package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/urfave/cli"

	"github.com/rosso0815/go_ImageResizer/mygraphics"
)

func workerConvert(id int, jobs <-chan string, results chan<- string) {
	for j := range jobs {
		log.Println("worker", id, "started  job", j)
		img1, _ := mygraphics.ReadMetaInfo(j)
		mygraphics.WriteResizedImages(img1)

		results <- j + " done"
	}
}

func runConvert(paths []string) error {

	log.Println("@@@ runConvert path=", paths)

	if len(paths) == 0 {
		log.Println("paths.length == 0")
		paths = append(paths, ".")
	}

	for _, path := range paths {
		log.Println("path=", path)

		if stat, err := os.Stat(path); err == nil && stat.IsDir() {
			log.Println(path, "is a directory")
		} else {
			log.Fatal("dir ", os.Args[1], " => this is not a directory , exit")
		}

		// go worker stuff
		jobs := make(chan string, 1000)
		results := make(chan string, 1000)
		for w := 1; w <= 8; w++ {
			go workerConvert(w, jobs, results)
		}

		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}
		maxJobs := 0
		for _, f := range files {
			if f.IsDir() == false {
				abs, _ := filepath.Abs(filepath.Join(path, f.Name()))
				log.Println("abs=", abs)
				//images = append(images, abs)

				// add to worker - queue
				jobs <- abs
				maxJobs++
			}
		}
		close(jobs)

		for a := 1; a <= maxJobs; a++ {
			log.Println("result ", a, " ", <-results)
		}
	}
	return nil
}

func main() {

	runtime.GOMAXPROCS(8)

	log.Println("@@@ start")

	app := cli.NewApp()
	app.Name = "go_myImageResizer"
	app.Usage = "path"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name: "convert",
			//Aliases: []string{"a"},
			Usage: "convert a given path the included images",
			Action: func(c *cli.Context) error {
				runConvert(c.Args())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("@@@ done")
}
