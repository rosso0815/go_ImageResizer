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
type ImageProcess struct {
	fabric string
	mw     *imagick.MagickWand
	img    Image
}

func init() {
	log.Println("mygraphics_impl -> init")
	imagick.Initialize()
	defer imagick.Terminate()
}

// NewProcessImplImages to do the real stuff
func NewProcessImplImages() (*ImageProcess, error) {
	log.Println("@@@ NewProcessImplImages")
	return &ImageProcess{fabric: "real-worker"}, nil
}

// ReadFileFromPath does the thing
func (ip *ImageProcess) ReadFileFromPath(lPath string) (err error) {
	log.Println("@@@ ReadFileFromPath path=", lPath)
	log.Println("path.Ext=", strings.ToLower(path.Ext(lPath)), "xxx")
	myExt := strings.ToLower(path.Ext(lPath))

	// error by not jpg files ...
	if strings.Compare(myExt, ".jpg") != 0 {
		return errors.New("mygraphics: cannot open not-jpg file")
	}

	ip.mw = imagick.NewMagickWand()
	ip.mw.ReadImage(lPath)
	ip.img.width = ip.mw.GetImageWidth()
	ip.img.heigth = ip.mw.GetImageHeight()
	ip.img.make = ip.mw.GetImageProperty("exif:Make")
	ip.img.model = ip.mw.GetImageProperty("exif:Model")
	ip.img.model = strings.ReplaceAll(ip.img.model, " ", "")
	ip.img.created = ip.mw.GetImageProperty("exif:DateTimeOriginal")
	ip.img.path = lPath
	return nil
}

// SaveFileResized reads the file and analyze it
// TODO handle folder of actual file
func (ip *ImageProcess) SaveFileResized() (err error) {

	log.Println("@@@ SaveFileResized")
	//var err error
	if (ip.img.heigth < ip.img.width) && (ip.img.width > 1980) {

		scale := ip.img.width * 1000 / 1980
		newHeigth := ip.img.heigth * 1000 / scale
		log.Println("image width > 1980 scale=", scale, " newHheigth=", newHeigth)

		//err = image.wand.ResizeImage(uint(1980), newHeigth, imagick.FILTER_LANCZOS)
		err = ip.mw.ResizeImage(uint(1980), newHeigth, imagick.FILTER_LANCZOS, 1)
		if err != nil {
			panic(err)
		}

		// convert_date
		layout := "2006:01:02 15:04:05"
		tm, err := time.Parse(layout, ip.img.created)
		if err != nil {
			return errors.New("mygraphics: canno parse date")
		}
		log.Printf("dateTimeOriginal=%q t=%v\n", ip.img.created, tm)

		fixedModel := strings.Replace(ip.img.model, "-", "", 10)
		log.Println("image.model=", ip.img.model, " fixedModel=", fixedModel)
		newFilename := tm.Format("20060102_150405") +
			"_" +
			strings.ToLower(ip.img.make) +
			"_" +
			strings.ToLower(fixedModel) +
			".jpg"

		log.Println("new_filename", newFilename)

		log.Println("dir  = ", path.Dir(""))
		log.Println("base = ", path.Base(ip.img.path))

		newpath := path.Join(path.Dir(ip.img.path), path.Base(newFilename))
		log.Println("newpath=", newpath)

		ip.mw.SetImageCompressionQuality(95)
		ip.mw.WriteImage(newpath)
	}

	return nil
}

// GetInfo reads the file and analyze it
func (ip *ImageProcess) GetInfo() (img Image) {
	log.Println("@@@ GetInfo")
	return ip.img
}
