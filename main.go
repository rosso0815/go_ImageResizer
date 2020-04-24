package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/rosso0815/go_ImageResizer/mygraphics"
)

func init() {
	log.Println("main -> init")
}

func main() {

	//runtime.GOMAXPROCS(8)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("start")

	app := cli.NewApp()
	app.Name = "go_myImageResizer"
	app.Usage = "path"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name: "convert",
			//Aliases: []string{"a"},
			Usage: "convert a given path the included images",
			Action: func(c *cli.Context) error {
				mygraphics.RunConvert(c.Args())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("@@@ done")
}
