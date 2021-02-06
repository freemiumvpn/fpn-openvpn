package api

import (
	"context"

	"github.com/freemiumvpn/fpn-openvpn-server/generated/vpn"
	"github.com/freemiumvpn/fpn-openvpn-server/internal/openvpn/mi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	// Options for the api
	Options struct {
		MiAddress string
	}

	// API structure
	API struct {
		mi *mi.ManagementInterface
	}
)

// New instantiates the API
func New(ctx context.Context, options Options) (*API, error) {
	managementInterface, err := mi.New(options.MiAddress)

	if err != nil {
		return nil, err
	}

	return &API{
		mi: managementInterface,
	}, nil
}

// Create a new vpn client
func (*API) Create(context.Context, *vpn.CreateRequest) (*vpn.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

// Delete a vpn client
func (*API) Delete(context.Context, *vpn.DeleteRequest) (*vpn.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// Connect a client
func (*API) Connect(context.Context, *vpn.ConnectRequest) (*vpn.ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

// Disconnect a client
func (*API) Disconnect(context.Context, *vpn.DisconnectRequest) (*vpn.DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}

// GetSession from vpn
func (*API) GetSession(context.Context, *vpn.GetSessionRequest) (*vpn.GetSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}

// SubscribeToSession id
func (*API) SubscribeToSession(*vpn.SubsribeToSessionRequest, vpn.VpnService_SubscribeToSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToSession not implemented")
}
