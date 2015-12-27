package main

import (
	"github.com/rakyll/command"
	"github.com/blackspace/goreader/subcommands"
)

func main() {
	subcommands.LoadSubcommand()
	command.Parse()
	command.Run()
}
