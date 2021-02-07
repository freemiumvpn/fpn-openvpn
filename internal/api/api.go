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
		RemoteIP  string
	}

	// API structure
	API struct {
		mi      *mi.ManagementInterface
		options Options
	}
)

// New instantiates the API
func New(ctx context.Context, options Options) (*API, error) {
	managementInterface, err := mi.New(options.MiAddress)

	if err != nil {
		return nil, err
	}

	return &API{
		mi:      managementInterface,
		options: options,
	}, nil
}

// Create a new vpn client
func (a *API) Create(ctx context.Context, createRequest *vpn.CreateRequest) (*vpn.CreateResponse, error) {

	clientConfig, err := CreateClient(createRequest.UserId, a.options.RemoteIP)
	if err != nil {
		return nil, err
	}

	return &vpn.CreateResponse{
		Credentials: string(clientConfig),
		Status:      vpn.VpnSessionStatus_CREATE_REQUEST_APPROVED,
	}, nil
}

// Delete a vpn client
func (a *API) Delete(ctx context.Context, deleteRequest *vpn.DeleteRequest) (*vpn.DeleteResponse, error) {
	err := DeleteClient(deleteRequest.UserId)
	if err != nil {
		return nil, err
	}

	return &vpn.DeleteResponse{
		Status: vpn.VpnSessionStatus_DELETE_REQUEST_APPROVED,
	}, nil
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
