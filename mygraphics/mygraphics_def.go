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

// ImageHandler used for ?
type ImageHandler interface {
	ReadFileFromPath(path string) (err error)
	GetInfo() (img Image)
	SaveFileResized() (err error)
}
