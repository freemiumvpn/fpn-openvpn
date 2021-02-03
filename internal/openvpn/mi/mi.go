package mi

import (
	"fmt"
	"net"
)

/**
mi stands for Management Interface
*/

const (
	outputSuccess = "SUCCESS: "
	outputError   = "ERROR: "
	outputEnd     = "END"
)

type (
	// ManagementInterface structure
	ManagementInterface struct {
		client  net.Conn
		replies <-chan []byte
		events  <-chan []byte
	}
)

// New instantiates a management interface
func New(port string) (*ManagementInterface, error) {
	conn, err := net.Dial("tcp", port) // tcp vs unix
	if err != nil {
		return nil, err
	}

	// conn.SetWriteDeadline(time.Time{})

	replies := make(chan []byte, 2)
	events := make(chan []byte, 2)

	go CreateConnectionListener(conn, events, replies)

	return &ManagementInterface{
		client:  conn,
		events:  events,
		replies: replies,
	}, nil
}

// GetHelp maps to help
func (mi *ManagementInterface) GetHelp() ([]byte, error) {
	return mi.write([]byte("help \n"))
}

func (mi *ManagementInterface) write(cmd []byte) ([]byte, error) {
	_, err := mi.client.Write(cmd)
	if err != nil {
		return nil, err
	}

	return mi.read()
}

func (mi *ManagementInterface) read() ([]byte, error) {
	reply, ok := <-mi.replies
	if !ok {
		return nil, fmt.Errorf("Reply channel closed while awaiting")
	}

	return reply, nil
}
