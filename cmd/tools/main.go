package main

import (
	"fmt"
	"github.com/tarxzfv/go-deepl/pkg/deepl"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var app *cli.App

func init() {
	app = cli.NewApp()

	app.Version = "0.1.0"
	app.Name = "go-deepl"
	app.Usage = "Client for DeepL API (https://www.deepl.com/docs-api/)"
	app.EnableBashCompletion = true

	app.Commands = []*cli.Command{
		{
			Name:  "translate",
			Usage: "Translate texts",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Required: false, // If this is omitted, the API will auto-detect the language of the text
					Name:     "source_lang",
					Aliases:  []string{"s"},
					Value:    "en",
					Usage:    "Language of the text to be translated",
				},
				&cli.StringFlag{
					Required: true,
					Name:     "target_lang",
					Aliases:  []string{"t"},
					Value:    "ja",
					Usage:    "The language into which the text should be translated",
				},
			},
			Action: func(ctx *cli.Context) error {
				t, err := InitializeTranslator(&deepl.Config{
					URL:     "",
					AuthKey: "auth",
				})
				if err != nil {
					return err
				}

				v, err := t.ProcessWithCLI(ctx)
				if err != nil {
					return err
				}

				fmt.Println(v)

				return nil
			},
		},
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
