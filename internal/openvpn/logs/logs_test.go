package logs_test

import (
	"testing"

	"github.com/freemiumvpn/fpn-openvpn-server/internal/openvpn/logs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLogInput(t *testing.T) {
	input := `OpenVPN CLIENT LIST
Updated,Mon Apr 19 21:17:41 2021
Common Name,Real Address,Bytes Received,Bytes Sent,Connected Since
foo-bar,000.00.0.0:00000,9700028,653546,Mon Apr 19 21:16:18 2021
foo-gizmo,000.00.0.0:00000,9700028,653546,Mon Apr 19 21:16:18 2021
ROUTING TABLE
Virtual Address,Common Name,Real Address,Last Ref
00.000.0.0,foo-bar,000.00.0.0:00000,Mon Apr 19 21:17:40 2021
GLOBAL STATS
Max bcast/mcast queue length,0
END
`

	output, err := logs.ParseLogInput(input)
	require.NoError(t, err)

	expected := []logs.Client{
		{
			Name:           "foo-bar",
			ConnectedSince: "Mon Apr 19 21:16:18 2021",
			BytesReceived:  "9700028",
			BytesSent:      "653546",
		},
		{
			Name:           "foo-gizmo",
			ConnectedSince: "Mon Apr 19 21:16:18 2021",
			BytesReceived:  "9700028",
			BytesSent:      "653546",
		},
	}

	assert.Equal(t, output.UpdatedAt, "Mon Apr 19 21:17:41 2021")
	assert.Equal(t, output.Clients, expected)
}

func TestParseLogFile(t *testing.T) {
	_, err := logs.ParseLogFile("./dummy.log")
	require.NoError(t, err)
}

func TestParseLogs(t *testing.T) {
	output, err := logs.ParseLogs("./dummy.log")
	require.NoError(t, err)

	expected := []logs.Client{
		{
			Name:           "foo-bar",
			ConnectedSince: "Mon Apr 19 21:16:18 2021",
			BytesReceived:  "9700028",
			BytesSent:      "653546",
		},
		{
			Name:           "foo-gizmo",
			ConnectedSince: "Mon Apr 19 21:16:18 2021",
			BytesReceived:  "9700028",
			BytesSent:      "653546",
		},
	}

	assert.Equal(t, output.UpdatedAt, "Mon Apr 19 21:17:41 2021")
	assert.Equal(t, output.Clients, expected)
}
