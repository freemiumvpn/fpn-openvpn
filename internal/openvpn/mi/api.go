package mi

type (
	// Command for management interface
	Command struct {
		Help      string
		LoadStats string
	}
)

var (
	command = Command{
		Help:      "help \n",
		LoadStats: "load-stats \n",
	}
)

// GetHelp maps to help
func (mi *ManagementInterface) GetHelp() (string, error) {
	err := mi.write([]byte(command.Help))
	if err != nil {
		return "", err
	}

	response, err := mi.readMultiline()
	if err != nil {
		return "", err
	}

	return ParseLines(response), nil
}

// GetStats Shows global server load stats.
func (mi *ManagementInterface) GetStats() (string, error) {
	err := mi.write([]byte(command.LoadStats))
	if err != nil {
		return "", err
	}

	response, err := mi.read()
	if err != nil {
		return "", err
	}

	return ParseLine(response), nil
}
