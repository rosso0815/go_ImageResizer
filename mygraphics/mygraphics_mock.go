package mygraphics

import (
	"log"
)

// MockImage to play
type MockImage struct {
	fabric string
}

// func init() {
// 	log.Println("@@@ mock init")
// }

// NewProcessMockImages handles the execution etc
func NewProcessMockImages() (*MockImage, error) {
	log.Println("@@@ NewProcessMockImages")
	return &MockImage{fabric: "hellau"}, nil
}

// ReadFileFromPath ddd
func (mi *MockImage) ReadFileFromPath(path string) (err error) {
	log.Println("@@@ ReadFileFromPath", path)
	return nil
}

// func getImageProcessor() (*ImageProcess, error) {
// 	return &ImageProcess{fabric: "hellau"}, nil
// }

// ImageMagickInternals play
// type ImageMagickInternals struct {
// 	wand  *imagick.MagickWand
// 	image Image
// }

// Interface01 play
// type Interface01 struct {
// 	Path string
// }

// ReadFile for test interface
// func (if01 Interface01) ReadFile(path string) (err error) {
// 	log.Println("@@@ if01 ReadData path       = " + path)
// 	log.Println("@@@ if01 ReadData local_path = " + if01.Path)
// 	return nil
// }

// ReadFile reads the file and analyze it
// func SaveFileResized(filePath string) (img Image, err error) {

// 	log.Println("@@@ SaveFileResized")
// 	//var err error
// 	if (im.image.heigth < im.image.width) && (im.image.width > 1980) {

// 		scale := im.image.width * 1000 / 1980
// 		newHeigth := im.image.heigth * 1000 / scale
// 		log.Println("image width > 1980 scale=", scale, " newHheigth=", newHeigth)

// 		//err = image.wand.ResizeImage(uint(1980), newHeigth, imagick.FILTER_LANCZOS)
// 		err = im.wand.ResizeImage(uint(1980), newHeigth, imagick.FILTER_LANCZOS, 1)
// 		if err != nil {
// 			panic(err)
// 		}

// 		// convert_date
// 		layout := "2006:01:02 15:04:05"
// 		tm, err := time.Parse(layout, im.image.created)
// 		if err != nil {
// 			return im.image, errors.New("mygraphics: canno parse date")
// 		}
// 		log.Printf("dateTimeOriginal=%q t=%v\n", im.image.created, tm)

// 		fixedModel := strings.Replace(im.image.model, "-", "", 10)
// 		log.Println("image.model=", im.image.model, " fixedModel=", fixedModel)
// 		newFilename := tm.Format("20060102_150405") +
// 			"_" +
// 			strings.ToLower(im.image.make) +
// 			"_" +
// 			strings.ToLower(fixedModel) +
// 			".jpg"

// 		log.Println("new_filename", newFilename)

// 		log.Println("dir  = ", path.Dir(""))
// 		log.Println("base = ", path.Base(im.image.path))

// 		newpath := path.Join(path.Dir(im.image.path), path.Base(newFilename))
// 		log.Println("newpath=", newpath)

// 		im.wand.SetImageCompressionQuality(95)
// 		im.wand.WriteImage(newpath)
// 	}

// 	return im.image, nil
// }
