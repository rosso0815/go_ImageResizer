package mygraphics

import (
	"errors"
	"log"
	"path"
	"strings"
	"time"

	"gopkg.in/gographics/imagick.v2/imagick"
)

// ImageProcess based implementation
type IM6Image struct {
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
func NewImageMagick6Handler() *IM6Image {
	log.Println("@@@ NewImageMagick6Handler")
	// return nil
	//return &IM6_ReadFile
	//GetInfo() (img Image)
	return &IM6Image{fabric: "real-worker"}
}

// ReadFileFromPath does the thing
func (im6 *IM6Image) ReadFile(lPath string) (err error) {
	log.Println("@@@ ReadFileFromPath path=", lPath)
	log.Println("path.Ext=", strings.ToLower(path.Ext(lPath)))
	myExt := strings.ToLower(path.Ext(lPath))

	// error by not jpg files ...
	// DONE accept HEIC also
	if !strings.Contains(".jpg .heic", myExt) {
		log.Println("NOT .jpg .heic", myExt)
		return errors.New("mygraphics: cannot open not-jpg file")
	}

	im6.mw = imagick.NewMagickWand()
	im6.mw.ReadImage(lPath)
	im6.img.width = im6.mw.GetImageWidth()
	im6.img.heigth = im6.mw.GetImageHeight()
	im6.img.make = im6.mw.GetImageProperty("exif:Make")
	im6.img.model = im6.mw.GetImageProperty("exif:Model")
	im6.img.model = strings.ReplaceAll(im6.img.model, " ", "")
	im6.img.created = im6.mw.GetImageProperty("exif:DateTimeOriginal")
	im6.img.path = lPath
	log.Println("found", im6.img)
	return nil
}

// SaveFileResized reads the file and analyze it
// TODO handle folder of actual file
// TODO handle correct extension
func (im6 *IM6Image) SaveFileResized() (err error) {

	log.Println("@@@ SaveFileResized, heigth:", im6.img.heigth, " width:", im6.img.width)

	// TODO : possibel as cmd args defined ?
	var (
		newWidth  uint = 1980
		newHeigth uint = 960
	)

	// calculate newWidth or newHeigth based on existing values
	if (im6.img.heigth < im6.img.width) && (im6.img.width > newWidth) {
		log.Println("resize width > ", newWidth)
		scale := im6.img.width * 1000 / newWidth
		newHeigth = im6.img.heigth * 1000 / scale
	} else if im6.img.heigth > 960 {
		log.Println("resize width heigth > 960")
		scale := im6.img.heigth * 1000 / newHeigth
		newWidth = im6.img.width * 1000 / scale
	} else {
		log.Println("no resize possible")
		return nil
	}
	log.Println("newWidth=", newWidth, "newHeigth", newHeigth)

	// resize the picture
	err = im6.mw.ResizeImage(newWidth, newHeigth, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		panic(err)
	}

	// convert_date
	layout := "2006:01:02 15:04:05"
	tm, err := time.Parse(layout, im6.img.created)
	if err != nil {
		return errors.New("mygraphics: canno parse date")
	}

	// calculate the new filename
	fixedModel := strings.Replace(im6.img.model, "-", "", 10)
	log.Println("image.model=", im6.img.model, " fixedModel=", fixedModel)
	newFilename := tm.Format("20060102_150405") +
		"_" +
		strings.ToLower(im6.img.make) +
		"_" +
		strings.ToLower(fixedModel) +
		".jpg"
	newpath := path.Join(path.Dir(im6.img.path), path.Base(newFilename))
	im6.mw.SetImageCompressionQuality(95)
	im6.mw.WriteImage(newpath)

	return nil
}

// GetInfo reads the file and analyze it
func (im6 *IM6Image) GetInfo() (img Image) {
	log.Println("@@@ GetInfo", im6.img)
	return im6.img
}
