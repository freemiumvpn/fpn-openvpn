package main

import (
	"os"

	"github.com/freemiumvpn/fpn-openvpn-server/generated/vpn"
	"github.com/freemiumvpn/fpn-openvpn-server/internal/grpc"
	"github.com/freemiumvpn/fpn-openvpn-server/internal/openvpn/mi"
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
			Name:    "management-interface-port",
			Usage:   "Management Interface Port",
			EnvVars: []string{"MANAGEMENT_INTERFACE_PORT"},
			Value:   ":5555",
		},
		&cli.StringFlag{
			Name:    "grpc-port",
			Usage:   "GRPC port",
			EnvVars: []string{"GRPC_PORT"},
			Value:   ":8989",
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
	port := cliCtx.String("management-interface-port")

	ctx := cliCtx.Context
	eg, ctx := errgroup.WithContext(ctx)

	grpcServer := grpc.New(ctx, grpc.Options{
		Address: cliCtx.String("grpc-port"),
	})

	eg.Go(func() error {
		return grpcServer.Listen(ctx, &vpn.UnimplementedVpnServiceServer{})
	})

	eg.Go(func() error {
		managementInterface, err := mi.New(port)
		if err != nil {
			return err
		}
		reply, _ := managementInterface.GetHelp()
		println(string(reply))

		reply, _ = managementInterface.GetStats()
		println(string(reply))
		return nil
	})

	return eg.Wait()
}
