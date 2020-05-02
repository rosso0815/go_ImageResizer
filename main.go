package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/rosso0815/go_ImageResizer/myjobscheduler"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	app := &cli.App{
		Name:    "go_ImageResizer",
		Version: "1.0.0",
		Commands: []*cli.Command{
			{
				Name:    "convert",
				Aliases: []string{"c"},
				Usage:   "convert a directory of images",
				Action: func(c *cli.Context) error {
					myjobscheduler.RunConvert(c.Args().First())
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
