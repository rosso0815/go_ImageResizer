package myjobscheduler

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/rosso0815/go_ImageResizer/mygraphics"
)

// RunConvert does the job
func RunConvert(path string) error {
	log.Println("@@@ runConvert path=", path)

	if len(path) == 0 {
		path = "."
	}
	log.Println("path=", path)

	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		log.Println(path, "is a directory")
	} else {
		log.Fatal("dir ", os.Args[1], " => this is not a directory , exit")
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() == false {
			abs, _ := filepath.Abs(filepath.Join(path, f.Name()))
			log.Println("abs=", abs)
			convertFile(abs)
		}
	}
	log.Println("close(jobs)")
	return nil
}

func convertFile(lPath string) {
	log.Println("@@@ convertFile", lPath)
	mImg, _ := mygraphics.NewProcessImplImages()
	imageHandler := mImg
	imageHandler.ReadFileFromPath(lPath)
	log.Println("GetInfo =", imageHandler.GetInfo())
	imageHandler.SaveFileResized()
}
