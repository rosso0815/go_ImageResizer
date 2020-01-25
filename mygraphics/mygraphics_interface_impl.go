package mygraphics

import "log"

// Interface01 play
type Interface01 struct {
	Path string
}

// ReadFile for test interface
func (if01 Interface01) ReadFile(path string) (err error) {
	log.Println("@@@ if01 ReadData path       = " + path)
	log.Println("@@@ if01 ReadData local_path = " + if01.Path)
	return nil
}
