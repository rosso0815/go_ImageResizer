package mygraphics

import (
	"errors"
	"gopkg.in/gographics/imagick.v2/imagick"
	"log"
	"path"
	"strings"
	"time"
)

// Image structur to handle our images
type Image struct {
	width   uint
	heigth  uint
	make    string
	model   string
	created string
	path    string
	wand    *imagick.MagickWand
}

func init() {
	imagick.Initialize()
	defer imagick.Terminate()
}

// WriteResizedImages :
func WriteResizedImages(image Image) error {
	log.Println("@@@ WriteResizedImages")
	var err error
	if (image.heigth < image.width) && (image.width > 1980) {

		scale := image.width * 1000 / 1980
		newHeigth := image.heigth * 1000 / scale
		log.Println("image width > 1980 scale=", scale, " newHheigth=", newHeigth)

		err = image.wand.ResizeImage(1980, newHeigth, imagick.FILTER_LANCZOS, 1)
		if err != nil {
			panic(err)
		}

		// convert_date
		layout := "2006:01:02 15:04:05"
		tm, err := time.Parse(layout, image.created)
		if err != nil {
			return errors.New("mygraphics: canno parse date")
		}
		log.Printf("dateTimeOriginal=%q t=%v\n", image.created, tm)

		fixedModel := strings.Replace(image.model, "-", "", 10)
		log.Println("image.model=", image.model, " fixedModel=", fixedModel)
		newFilename := tm.Format("20060102_150405") +
			"_" +
			strings.ToLower(image.make) +
			"_" +
			strings.ToLower(fixedModel) +
			".jpg"

		log.Println("new_filename", newFilename)

		log.Println("dir  = ", path.Dir(image.path))
		log.Println("base = ", path.Base(image.path))

		newpath := path.Join(path.Dir(image.path), path.Base(newFilename))
		log.Println("newpath=", newpath)
		image.wand.WriteImage(newpath)
	}
	return nil
}

// ReadMetaInfo read the info from path
// wand can be nil
func ReadMetaInfo(myPath string) (Image, error) {

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

	return Image{
		width:   w,
		heigth:  h,
		make:    make,
		model:   model,
		created: created,
		path:    myPath,
		wand:    mw}, nil

}
