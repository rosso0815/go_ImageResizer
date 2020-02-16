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
	ReadFileFromPath(path string) (err error)
}

// ImageProcessor handles the execution etc
// type ImageProcessor interface {
// 	ProcessImage(imageHandler ImageHandler) (err error)
// }
