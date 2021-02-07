package api

import (
	"io/ioutil"
	"os/exec"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	easyRsaDir = "/etc/openvpn/certs"
	setupDir   = "/etc/openvpn/setup"
)

// CreateClient from setup files
func CreateClient(clientID string, remoteIP string) ([]byte, error) {
	command := path.Join(setupDir, "/new-client-cert.sh")
	args := []string{command, clientID, remoteIP}

	logrus.WithFields(logrus.Fields{
		"command":  command,
		"clientID": clientID,
		"remoteIP": remoteIP,
	}).Info("Creating Client")

	_, err := exec.Command("/bin/bash", args...).Output()
	if err != nil {
		logrus.
			WithError(err).
			Error("Failed to execute easyrsa command")

		return nil, err
	}

	clientConfigPath := path.Join(easyRsaDir, "/pki/", clientID+".ovpn")
	return ioutil.ReadFile(clientConfigPath)
}
