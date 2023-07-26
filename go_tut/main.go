package main

import (
	"go_tut/hello_world"
	"go_tut/variables"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "My golang Tutorial",
		Usage: "This is a tutorial for golang",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "variables",
				Aliases: []string{"v"},
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.Bool("variables") {
				variables.Exec()
			} else if ctx.Bool("help") {
				cli.ShowAppHelp(ctx)
			} else {
				hello_world.Run()
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
