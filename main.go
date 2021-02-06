package main

import (
	"os"

	"github.com/freemiumvpn/fpn-openvpn-server/internal/openvpn/mi"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	appName        = "fpn-openvpn"
	appDescription = "openvpn wrapper"
)

var (
	flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "management-interface-port",
			Usage:   "Management Interface Port",
			EnvVars: []string{"MANAGEMENT_INTERFACE_PORT"},
			Value:   ":5555",
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Description = appDescription
	app.Flags = flags
	app.Action = appAction

	err := app.Run(os.Args)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to run")
	}
}

func appAction(ctx *cli.Context) error {
	port := ctx.String("management-interface-port")
	managementInterface, err := mi.New(port)
	if err != nil {
		return err
	}

	reply, _ := managementInterface.GetHelp()
	println(string(reply))

	reply, _ = managementInterface.GetStats()
	println(string(reply))

	return nil
}
