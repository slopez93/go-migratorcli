package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Analytics Migrator CLI",
		Usage: "CLI for running and generating migrations for the analytics ClickHouse database",
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "Generate migration files",
				Flags: []cli.Flag{},
				Action: func(c *cli.Context) error {
					fmt.Printf("Hello from generate command")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("%v", err)
	}
}
