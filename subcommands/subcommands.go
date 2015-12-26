package subcommands

import "github.com/rakyll/command"

func LoadSubcommand() {
	command.On("version", "prints the version", &VersionCommand{}, nil)
	command.On("server", "start the server for listenning", &ServerCommand{}, nil)
	command.On("install","install the server to your system",&InstallCommand{},nil)
}

