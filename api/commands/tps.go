package commands

import (
	"strings"
)

func (c *CommandClient) GetTps(serverName string, getLast bool) *Command {
	return &Command{
		CommandName: "tps",
		Run: func() (result string, err error) {
			command, err := c.getResultFromCommand(serverName, "tps")
			if err != nil {
				return
			}
			if getLast {
				return strings.Split(strings.Split(command, "TPS from last 1m, 5m, 15m: ")[1], ", ")[0], nil
			}
			return command, nil
		},
	}
}
