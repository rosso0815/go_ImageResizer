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
	ReadFile(path string) (err error)

	//SaveFileResized() (err error)
}
