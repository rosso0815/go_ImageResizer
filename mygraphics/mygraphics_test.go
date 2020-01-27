package mygraphics

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestInterface(t *testing.T) {
	log.Println("@@@ TestInterface")

	var i1 ImageHandler

	i1 = Interface01{"undefined"}

	i1.ReadFile("test01.jpg")

	i1 = ImageMagickInternals{}

	i1.ReadFile("test01.jpg")

}

func deleteOldFiles() {
	log.Println("@@@ deleteOldFiles")
	files, err := filepath.Glob("20*jpg")
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

	// Interface Handler
	//var img ImageHandler

	//r := ImageMagickInternals{}

	//img = r

	//img.ReadFile("test01.jpg")
	//log.Println("r=", r)

	// if img1.width != 5472 {
	// 	log.Panic("image wrong height")
	// }

	// if img1.heigth != 3080 {
	// 	log.Panic("image wrong heigth")
	// }

	// if img1.make != "SONY" {
	// 	log.Panic("image wrong make")
	// }

	// if img1.model != "DSC-RX100M3" {
	// 	log.Panic("image wrong model")
	// }

	// if img1.created != "2018:04:20 03:42:56" {
	// 	log.Panic("image wrong date")
	// }

	log.Println("@@@ done")
}

func TestResizedImage(t *testing.T) {

	log.Println("@@@ TestResizedImage")

	deleteOldFiles()

	//img1, _ := ReadMetaInfo("test01.jpg")
	//WriteResizedImages(img1)

	//img2, _ := ReadMetaInfo("test02.jpg")
	//WriteResizedImages(img2)

	//deleteOldFiles()
}
