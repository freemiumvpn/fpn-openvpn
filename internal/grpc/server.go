package grpc

import (
	"context"
	"net"

	"github.com/freemiumvpn/fpn-openvpn-server/generated/vpn"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type (
	// Options for the grpc server
	Options struct {
		Address    string
		AuthURL    string
		AuthScopes []string
	}

	// Service grpc
	Service struct {
		server  *grpc.Server
		options Options
	}
)

// New instantiates a grpc server
func New(ctx context.Context, options Options) *Service {
	grpc_prometheus.EnableHandlingTimeHistogram()
	server := grpc.NewServer()
	grpc_prometheus.Register(server)

	return &Service{
		server:  server,
		options: options,
	}
}

// Listen starts grp server
func (s *Service) Listen(ctx context.Context, vpnServer vpn.VpnServiceServer) error {
	vpn.RegisterVpnServiceServer(s.server, vpnServer)

	listener, err := net.Listen("tcp", s.options.Address)
	if err != nil {
		return errors.Errorf("failed to listen: %v", err)
	}

	logrus.
		WithField("address", s.options.Address).
		Info("Starting gRPC service")

	go func() {
		select {
		case <-ctx.Done():
			s.server.GracefulStop()
		}
	}()

	return s.server.Serve(listener)
}
