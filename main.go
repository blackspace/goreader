package main

import (
	"github.com/rakyll/command"
	"goreader/subcommands"
)

func main() {
	subcommands.LoadSubcommand()
	command.Parse()
	command.Run()
}
