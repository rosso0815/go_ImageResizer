package mygraphics

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func deleteOldFiles() {
	log.Println("@@@ deleteOldFiles")
	files, err := filepath.Glob("conv*jpg")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		log.Println("delete file ", f)
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}

// TestReadmetadata just to verify the actual pictures
func TestReadmetadata(t *testing.T) {
	log.Println("@@@ start")

	img1, _ := ReadMetaInfo("test01.jpg")
	if img1.width != 5472 {
		log.Panic("image wrong height")
	}

	if img1.heigth != 3080 {
		log.Panic("image wrong heigth")
	}

	if img1.make != "SONY" {
		log.Panic("image wrong make")
	}

	if img1.model != "DSC-RX100M3" {
		log.Panic("image wrong model")
	}

	if img1.created != "2018:04:20 03:42:56" {
		log.Panic("image wrong date")
	}

	log.Println("@@@ done")
}

func TestResizedImage(t *testing.T) {
	log.Println("@@@ TestResizedImage")

	deleteOldFiles()

	img1, _ := ReadMetaInfo("test01.jpg")
	WriteResizedImages(img1)

	img2, _ := ReadMetaInfo("test02.jpg")
	WriteResizedImages(img2)

	deleteOldFiles()
}
