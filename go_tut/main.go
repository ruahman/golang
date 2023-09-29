package main

import (
	"go_tut/cli_args"
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
			&cli.BoolFlag{
				Name:    "cli",
				Aliases: []string{"c"},
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.Bool("variables") {
				variables.Run()
			} else if ctx.Bool("help") {
				cli.ShowAppHelp(ctx)
			} else if ctx.Bool("cli") {
				cli_args.Exec()
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
