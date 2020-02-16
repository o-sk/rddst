package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/o-sk/rddst"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "rddst",
		Usage:       "rddst <url>",
		Description: "Get redirect destination",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				return cli.Exit("Argument error", 1)
			}

			r := rddst.NewRddst(&http.Client{})
			dst, err := r.GetRedirectDestination(c.Args().Get(0))
			if err != nil {
				return cli.Exit(err, 1)
			}

			fmt.Fprintf(c.App.Writer, "%s\n", dst)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
