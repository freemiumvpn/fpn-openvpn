package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	appName        = "fpn-openvpn"
	appDescription = "openvpn wrapper"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Description = appDescription
	app.Action = appAction

	err := app.Run(os.Args)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to run")
	}
}

func appAction(ctx *cli.Context) error {
	return nil
}
