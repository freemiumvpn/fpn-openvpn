package mi

import (
	"bytes"
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
	conn, err := net.Dial("tcp", port) // TODO: tcp vs unix
	if err != nil {
		return nil, err
	}

	// TODO: conn.SetWriteDeadline(time.Time{})

	replies := make(chan []byte, 2)
	events := make(chan []byte, 2)

	go CreateConnectionListener(conn, events, replies)

	return &ManagementInterface{
		client:  conn,
		events:  events,
		replies: replies,
	}, nil
}

func (mi *ManagementInterface) write(cmd []byte) error {
	_, err := mi.client.Write(cmd)
	return err
}

func (mi *ManagementInterface) readMultiline() ([][]byte, error) {
	output := make([][]byte, 0)

	for {
		reply, ok := <-mi.replies
		if !ok {
			return nil, fmt.Errorf("Reply channel closed while awaiting")
		}

		if bytes.Equal(reply, []byte(outputEnd)) {
			break
		}

		output = append(output, reply)
	}

	return output, nil
}

func (mi *ManagementInterface) read() ([]byte, error) {
	reply, ok := <-mi.replies
	if !ok {
		return nil, fmt.Errorf("Reply channel closed while awaiting")
	}

	return reply, nil
}

// ParseLine converts byte to string
func ParseLine(input []byte) string {
	return string(input)
}

// ParseLines converts bytes to string
func ParseLines(input [][]byte) string {
	output := ""
	for _, i := range input {
		output = output + string(i) + "\n"
	}
	return output
}
