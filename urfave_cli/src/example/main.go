package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "example"
	app.Usage = `this is sample.
	try some option.`
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "str, s",
			Usage: "input some text.",
		},
	}

	app.Action = func(context *cli.Context) error {
		if context.String("str") != "" {
			fmt.Println("option -s is " + context.String("str"))
		}

		for i := 0; i < context.NArg(); i++ {
			fmt.Println(context.Args().Get(i))
		}

		return nil
	}

	app.Run(os.Args)
}
