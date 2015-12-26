package subcommands

import "github.com/rakyll/command"

func LoadSubcommand() {
	command.On("version", "prints the version", &VersionCommand{}, nil)
	command.On("server", "start the server for listenning", &ServerCommand{}, nil)
}

