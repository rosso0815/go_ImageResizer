package mygraphics

import (
	"log"
	"testing"
)

func Test_Mock_Single_ReadFile(t *testing.T) {

	log.Println("@@@ Test_Mock_Single_ReadFile")

	var ih ImageHandler

	ih = NewImageHandlerMock()

	ih.ReadFile("test01.jpg")

	ih.GetInfo()

	ih.SaveFileResized()
}
