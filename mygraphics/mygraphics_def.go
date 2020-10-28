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

// ImageHandler used for different implementations
type ImageHandler interface {
	ReadFile(path string) (err error)
	//GetInfo() (img Image)
	//SaveFileResized() (err error)
}
