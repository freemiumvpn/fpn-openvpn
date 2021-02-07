package main

import (
	"os"

	"github.com/freemiumvpn/fpn-openvpn-server/internal/api"
	"github.com/freemiumvpn/fpn-openvpn-server/internal/grpc"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

const (
	appName        = "fpn-openvpn"
	appDescription = "openvpn wrapper"
)

var (
	flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "management-interface-port",
			Usage:    "Management Interface Port",
			EnvVars:  []string{"MANAGEMENT_INTERFACE_PORT"},
			Required: true,
			Value:    ":5555",
		},
		&cli.StringFlag{
			Name:     "grpc-port",
			Usage:    "GRPC port",
			EnvVars:  []string{"GRPC_PORT"},
			Required: true,
			Value:    ":8989",
		},
		&cli.StringFlag{
			Name:     "vpn-remote-ip",
			Usage:    "VPN remote IP",
			EnvVars:  []string{"VPN_REMOTE_IP"},
			Required: true,
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

func appAction(cliCtx *cli.Context) error {
	ctx := cliCtx.Context
	eg, ctx := errgroup.WithContext(ctx)

	grpcServer := grpc.New(ctx, grpc.Options{
		Address: cliCtx.String("grpc-port"),
	})
	vpnAPI, err := api.New(ctx, api.Options{
		MiAddress: cliCtx.String("management-interface-port"),
		RemoteIP:  cliCtx.String("vpn-remote-ip"),
	})

	if err != nil {
		return err
	}

	eg.Go(func() error {
		return grpcServer.Listen(ctx, vpnAPI)
	})

	return eg.Wait()
}
