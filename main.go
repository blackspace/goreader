package main

import (
	"github.com/rakyll/command"
	. "goreader/subcommand"
)




func main() {
	command.On("version", "prints the version", &VersionCommand{}, nil)
	command.On("server", "start the server for listenning", &ServerCommand{}, nil)
	command.Parse()
	command.Run()
}
