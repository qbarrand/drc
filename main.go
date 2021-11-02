package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/qbarrand/drc/pkg/bt"
	"github.com/qbarrand/drc/pkg/idasen"
	"github.com/urfave/cli/v2"
)

var (
	commit string
	version string
)

type config struct {
	Address string
}

func main() {
	cfg := config{}

	app := cli.NewApp()
	//app.ExtraInfo = func() map[string]string {
	//	return map[string]string{"commit": commit}
	//}
	//app.Metadata = map[string]interface{}{"commit": commit}
	app.Version = version
	app.Authors = []*cli.Author{
		{
			Name:  "Quentin Barrand",
			Email: "quentin@quba.fr",
		},
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "address",
			Aliases:     []string{"addr"},
			Usage:       "Connect to the desk at `ADDRESS`",
			Destination: &cfg.Address,
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:                   "get-current-height",
			Aliases:                []string{"curr"},
			Usage:                  "Returns the current height in centimeters",
			Action: func(ctx *cli.Context) error {
				height, err := getCurrentHeight(ctx.Context, cfg)
				if err != nil {
					return fmt.Errorf("could not get the current height: %v", err)
				}

				if _, err = fmt.Printf("%d\n", height); err != nil {
					return fmt.Errorf("could not print the current height: %v", err)
				}

				return nil
			},
		},
		{
			Name:                   "move-to",
			Aliases:                []string{"mv"},
			Usage:                  "Moves the desk to the specified height in centimeters",
			Description:            "",
			ArgsUsage:              "HEIGHT",
			Action: func(ctx *cli.Context) error {
				heightStr := ctx.Args().Get(0)

				height, err := strconv.Atoi(heightStr)
				if err != nil {
					return fmt.Errorf("%q: not a valid integer", heightStr)
				}

				log.Printf("Now trying to move to %dcm...", height)

				return moveTo(ctx.Context, cfg, height)
			},
		},
	}
	app.Before = func(c *cli.Context) error {
		if cfg.Address == "" {
			return errors.New("address undefined; please specify --name or --address")
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func getCurrentHeight(ctx context.Context, cfg config) (int, error) {
	log.Printf("Getting the height of %s", cfg.Address)

	dev, err := bt.GetDevice(cfg.Address)
	if err != nil {
		return 0, fmt.Errorf("could not connect to the device at %s: %v", cfg.Address, err)
	}

	d := idasen.NewIdasen(dev)

	return d.GetCurrentHeight(ctx)
}

func moveTo(ctx context.Context, cfg config, height int) error {
	return errors.New("not implemented")
}