package mygraphics

import (
	"log"
)

// MockImage to play
type MockImage struct {
	fabric string
	image  Image
}

func init() {
	log.Println("@@@ mock init")
}

// NewProcessMockImages handles the execution etc
func NewImageHandlerMock() *MockImage {
	log.Println("@@@ NewImageHandlerMock")
	return &MockImage{fabric: "hellau"}
}

// ReadFileFromPath ddd
func (mi *MockImage) ReadFile(path string) (err error) {
	log.Println("@@@ NewImageHandlerMock ReadFileFromPath", path)
	return nil
}

func (mi *MockImage) GetInfo() (img Image) {
	log.Println("@@@ NewImageHandlerMock GetInfo")
	return mi.image
}

func (mi *MockImage) SaveFileResized() (err error) {
	log.Println("@@@ NewImageHandlerMock SaveFileResized")
	return nil
}
