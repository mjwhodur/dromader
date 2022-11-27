package main

import (
	"fmt"
	"github.com/mjwhodur/dromader/server"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(ctx *cli.Context) error {
			fmt.Println("boom! I say!")
			ctx.Args()
			server.Serve()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
