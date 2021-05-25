package mygraphics

import (
	"log"
	"testing"
)

func Test_IM6_Single_ReadFile(t *testing.T) {

	log.Println("@@@ Test_IM6_Single_ReadFile")

	var ih ImageHandler

	ih = NewImageMagick6Handler()

	ih.ReadFile("test01.HEIC")

	log.Printf("model => %s", ih.GetInfo().model)

	ih.SaveFileResized()
}
