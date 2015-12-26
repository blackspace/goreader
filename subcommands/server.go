package subcommands
import (
	"goreader/network"
	"flag"
)

type ServerCommand struct {
	flagPort *int
}

func (cmd *ServerCommand) Flags(fs *flag.FlagSet) *flag.FlagSet {
	cmd.flagPort = fs.Int("p", 10443, "the port for listenning.")
	return fs
}

func (cmd *ServerCommand) Run(args []string) {
	network.LoadHttpsHandlers()

	network.Listen(*cmd.flagPort)
}
