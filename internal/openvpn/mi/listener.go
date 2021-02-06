package mi

import (
	"bufio"
	"net"

	"github.com/sirupsen/logrus"
)

// CreateConnectionListener dispatches events and replies
func CreateConnectionListener(conn net.Conn, events chan []byte, replies chan []byte) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		buffer := scanner.Bytes()

		if buffer[0] == '>' {
			events <- buffer
		} else {
			replies <- buffer
		}
	}

	if err := scanner.Err(); err != nil {
		errorMessage := "Scanner Error"
		logrus.Fatal(errorMessage)
		events <- []byte(errorMessage)
	}

	close(events)
	close(replies)
}
