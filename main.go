package main

import (
	"os"

	"github.com/freemiumvpn/fpn-openvpn-server/internal/api"
	"github.com/freemiumvpn/fpn-openvpn-server/internal/grpc"
	"github.com/freemiumvpn/fpn-openvpn-server/internal/observability"

	"github.com/pkg/errors"
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
		&cli.StringFlag{
			Name:    "observability-address",
			Usage:   "address on which to expose '/__/health' '/__/metrics' '/__/ready'",
			EnvVars: []string{"OBSERVABILITY_ADDRESS"},
			Value:   ":8081",
		},
		&cli.StringFlag{
			Name:    "vpn-remote-ip",
			Usage:   "VPN remote IP",
			EnvVars: []string{"VPN_REMOTE_IP"},
		},
		&cli.StringFlag{
			Name:    "log-level",
			Usage:   "Log level",
			EnvVars: []string{"LOG_LEVEL"},
			Value:   "DEBUG",
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
	if err := logger(cliCtx); err != nil {
		return err
	}

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

	eg.Go(func() error {
		return observability.New(ctx, cliCtx.String("observability-address"))
	})

	return eg.Wait()
}

func logger(cliCtx *cli.Context) error {
	level := cliCtx.String("log-level")
	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return errors.Wrap(err, "failed to parse log-level")
	}

	logrus.SetLevel(parsedLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return nil
}
