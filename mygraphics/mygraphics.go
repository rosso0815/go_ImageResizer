package mygraphics

import (
	"errors"
	"log"
	"path"
	"strings"
	"time"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func init() {

	imagick.Initialize()

	defer imagick.Terminate()
}

// ImageMagickInternals play
type ImageMagickInternals struct {
	wand  *imagick.MagickWand
	image Image
}

// ReadFile reads the file and analyze it
func (im ImageMagickInternals) SaveFileResized(filePath string) (img Image, err error) {

	log.Println("@@@ SaveFileResized")
	//var err error
	if (im.image.heigth < im.image.width) && (im.image.width > 1980) {

		scale := im.image.width * 1000 / 1980
		newHeigth := im.image.heigth * 1000 / scale
		log.Println("image width > 1980 scale=", scale, " newHheigth=", newHeigth)

		//err = image.wand.ResizeImage(uint(1980), newHeigth, imagick.FILTER_LANCZOS)
		err = im.wand.ResizeImage(uint(1980), newHeigth, imagick.FILTER_LANCZOS, 1)
		if err != nil {
			panic(err)
		}

		// convert_date
		layout := "2006:01:02 15:04:05"
		tm, err := time.Parse(layout, im.image.created)
		if err != nil {
			return im.image, errors.New("mygraphics: canno parse date")
		}
		log.Printf("dateTimeOriginal=%q t=%v\n", im.image.created, tm)

		fixedModel := strings.Replace(im.image.model, "-", "", 10)
		log.Println("image.model=", im.image.model, " fixedModel=", fixedModel)
		newFilename := tm.Format("20060102_150405") +
			"_" +
			strings.ToLower(im.image.make) +
			"_" +
			strings.ToLower(fixedModel) +
			".jpg"

		log.Println("new_filename", newFilename)

		log.Println("dir  = ", path.Dir(""))
		log.Println("base = ", path.Base(im.image.path))

		newpath := path.Join(path.Dir(im.image.path), path.Base(newFilename))
		log.Println("newpath=", newpath)

		im.wand.SetImageCompressionQuality(95)
		im.wand.WriteImage(newpath)
	}

	return im.image, nil
}

// ReadFile read the info from path
// wand can be nil
func (im ImageMagickInternals) ReadFile(myPath string) (Image, error) {

	log.Println("@@@ ReadMetaInfo path=", myPath)
	log.Println("path.Ext=", strings.ToLower(path.Ext(myPath)), "xxx")

	myExt := strings.ToLower(path.Ext(myPath))

	// error by not jpg files ...
	if strings.Compare(myExt, ".jpg") != 0 {
		return Image{}, errors.New("mygraphics: cannot open not-jpg file")
	}

	mw := imagick.NewMagickWand()

	mw.ReadImage(myPath)

	w := mw.GetImageWidth()
	h := mw.GetImageHeight()
	make := mw.GetImageProperty("exif:Make")
	model := mw.GetImageProperty("exif:Model")
	model = strings.ReplaceAll(model, " ", "")
	created := mw.GetImageProperty("exif:DateTimeOriginal")

	log.Println("width           = ", w)
	log.Println("height          = ", h)
	log.Println("attribute make  = ", make)
	log.Println("model           = ", model)
	log.Println("created         = ", created)

	return im.image, nil

}
