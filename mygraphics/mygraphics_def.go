package mygraphics

// Image structur to handle images
type Image struct {
	width   uint
	heigth  uint
	make    string
	model   string
	created string
	path    string
}

// ImageHandler used for test
type ImageHandler interface {
	ReadFileFromPath(path string) error

	//SaveFileResized() (err error)
}

// ImageProcessor hadles the execution etc
type ImageProcessor interface {
	ProcessImages(imageHandler ImageHandler) error
}
