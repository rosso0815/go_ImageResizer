package mygraphics

import (
	"errors"
	"log"
	"path"
	"strings"
	"time"

	"gopkg.in/gographics/imagick.v3/imagick"
)

// ImageProcess based implementation
type ImImage struct {
	fabric string
	mw     *imagick.MagickWand
	img    Image
}

func init() {
	log.Println("mygraphics_impl -> init")
	imagick.Initialize()
	defer imagick.Terminate()
}

// NewImageMagick6Handler to do the real stuff
func NewImageMagick6Handler() *ImImage {
	log.Println("@@@ NewImageMagick6Handler")
	return &ImImage{fabric: "real-worker"}
}

// ReadFileFromPath does the thing
func (im *ImImage) ReadFile(lPath string) (err error) {

	log.Println("@@@ ReadFileFromPath path=", lPath)
	log.Println("path.Ext=", strings.ToLower(path.Ext(lPath)))

	myExt := strings.ToLower(path.Ext(lPath))

	if !strings.Contains(".jpg .heic", myExt) {
		log.Println("NOT .jpg .heic", myExt)
		return errors.New("mygraphics: cannot open not-jpg file")
	}

	im.mw = imagick.NewMagickWand()
	im.mw.ReadImage(lPath)
	im.img.width = im.mw.GetImageWidth()
	im.img.heigth = im.mw.GetImageHeight()
	im.img.make = im.mw.GetImageProperty("exif:Make")
	im.img.model = im.mw.GetImageProperty("exif:Model")
	im.img.model = strings.ReplaceAll(im.img.model, " ", "")
	im.img.created = im.mw.GetImageProperty("exif:DateTimeOriginal")
	im.img.path = lPath
	log.Println("found", im.img)
	return nil
}

// SaveFileResized reads the file and analyze it
// TODO handle folder of actual file
// TODO handle correct extension
func (im *ImImage) SaveFileResized() (err error) {

	log.Println("@@@ SaveFileResized, heigth:", im.img.heigth, " width:", im.img.width)

	// TODO : possibel as cmd args defined ?
	var (
		newWidth  uint = 1980
		newHeigth uint = 960
	)

	// calculate newWidth or newHeigth based on existing values
	if (im.img.heigth < im.img.width) && (im.img.width > newWidth) {
		log.Println("resize width > ", newWidth)
		scale := im.img.width * 1000 / newWidth
		newHeigth = im.img.heigth * 1000 / scale
	} else if im.img.heigth > 960 {
		log.Println("resize width heigth > 960")
		scale := im.img.heigth * 1000 / newHeigth
		newWidth = im.img.width * 1000 / scale
	} else {
		log.Println("no resize possible")
		return nil
	}
	log.Println("newWidth=", newWidth, "newHeigth", newHeigth)

	// resize the picture
	err = im.mw.ResizeImage(newWidth, newHeigth, imagick.FILTER_LANCZOS)
	if err != nil {
		panic(err)
	}

	// convert_date
	layout := "2006:01:02 15:04:05"
	tm, err := time.Parse(layout, im.img.created)
	if err != nil {
		return errors.New("mygraphics: canno parse date")
	}

	// calculate the new filename
	fixedModel := strings.Replace(im.img.model, "-", "", -1)
	fixedModel = strings.Replace(fixedModel, "(", "", -1)
	fixedModel = strings.Replace(fixedModel, ")", "", -1)
	log.Println("image.model=", im.img.model, " fixedModel=", fixedModel)
	newFilename := tm.Format("20060102_150405") +
		"_" +
		strings.ToLower(im.img.make) +
		"_" +
		strings.ToLower(fixedModel) +
		".jpg"
	newpath := path.Join(path.Dir(im.img.path), path.Base(newFilename))
	log.Println("newpath", newpath)
	im.mw.SetImageCompressionQuality(95)
	im.mw.WriteImage(newpath)

	return nil
}

// GetInfo reads the file and analyze it
func (im *ImImage) GetInfo() (img Image) {
	log.Println("@@@ GetInfo", im.img)
	return im.img
}
