package subcommands
import (
	"flag"
	"fmt"
)

type VersionCommand struct{}

func (cmd *VersionCommand) Flags(fs *flag.FlagSet) *flag.FlagSet {
	return fs
}

func (cmd *VersionCommand) Run(args []string) {
	fmt.Println("Now it is in the developping")
}
