package api

import (
	"os"
	"os/exec"
	"path"

	"github.com/sirupsen/logrus"
)

// DeleteClient revokes a client key
func DeleteClient(clientID string) error {
	command := path.Join(setupDir, "/revoke-client-cert.sh")
	args := []string{command, clientID}

	logrus.WithFields(logrus.Fields{
		"command":  command,
		"clientID": clientID,
	}).Info("Deleting Client")

	output, err := exec.Command("/bin/bash", args...).Output()
	if err != nil {
		logrus.
			WithError(err).
			Error("Failed to delete client")

		return err
	}

	logrus.
		WithFields(logrus.Fields{
			"output": string(output),
		}).
		Debug("client revoked")

	clientConfigPath := path.Join(easyRsaDir, "/pki/", clientID+".ovpn")
	err = os.Remove(clientConfigPath)

	if err != nil {
		logrus.
			WithFields(logrus.Fields{
				"clientID": clientID,
			}).
			WithError(err).
			Error("Failed to delete ovpn file")

		return err
	}

	logrus.
		WithFields(logrus.Fields{
			"output": string(output),
		}).
		Debug("client ovpn file delete")
	return nil
}
