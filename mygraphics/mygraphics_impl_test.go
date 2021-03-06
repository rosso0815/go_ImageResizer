package mygraphics

import (
	"log"
	"testing"
)

func Test_IM_Single_ReadFileHEIC(t *testing.T) {

	log.Println("@@@ Test_IM_Single_ReadFile")

	var ih ImageHandler
	ih = NewImageMagick6Handler()
	ih.ReadFile("test01.HEIC")
	log.Printf("model => %s", ih.GetInfo().model)
	ih.SaveFileResized()

}

func Test_IM_Single_ReadFileJPG(t *testing.T) {

	log.Println("@@@ Test_IM_Single_ReadFile")

	var ih ImageHandler
	ih = NewImageMagick6Handler()
	ih.ReadFile("test01.jpg")
	log.Printf("model => %s", ih.GetInfo().model)
	ih.SaveFileResized()

}
