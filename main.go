package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "HealthCheck",
		Usage: "Get the health status of a website",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check health",
				Required: false,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}

			status := Check(c.String("domain"), port) //Check function present in check.go
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
