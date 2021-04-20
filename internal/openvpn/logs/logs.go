package logs

import (
	"io/ioutil"
	"strings"
)

type (
	Client struct {
		Name           string
		BytesReceived  string
		BytesSent      string
		ConnectedSince string
	}

	Logs struct {
		UpdatedAt string
		Clients   []Client
	}
)

func ParseLogInput(input string) (Logs, error) {
	fields := strings.Split(input, "\n")

	clientIndexStart := -1
	clientIndexEnd := -1
	updatedAt := ""

	for index, field := range fields {
		switch true {
		case strings.Contains(field, "Updated"):
			updatedAt = strings.Split(fields[index], ",")[1]
		case strings.Contains(field, "Common Name,Real Address,Bytes Received,Bytes Sent,Connected Since"):
			clientIndexStart = index + 1
		case strings.Contains(field, "ROUTING TABLE"):
			clientIndexEnd = index
		}
	}

	vpnClientsList := fields[clientIndexStart:clientIndexEnd]
	vpnClients := []Client{}

	for _, vpnClient := range vpnClientsList {
		clientSplit := strings.Split(vpnClient, ",")
		parsedClient := Client{
			Name:           clientSplit[0],
			BytesReceived:  clientSplit[2],
			BytesSent:      clientSplit[3],
			ConnectedSince: clientSplit[4],
		}
		vpnClients = append(vpnClients, parsedClient)
	}

	return Logs{
		UpdatedAt: updatedAt,
		Clients:   vpnClients,
	}, nil
}

func ParseLogFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ParseLogs(path string) (Logs, error) {
	data, err := ParseLogFile(path)
	if err != nil {
		return Logs{}, err
	}
	return ParseLogInput(data)
}
