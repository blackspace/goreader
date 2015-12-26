package subcommand
import (
	"flag"
	"fmt"
	"goreader/network"
	"log"
	"os"
	"net/http"
	. "strconv"
)

type VersionCommand struct{
}

func (cmd *VersionCommand) Flags(fs *flag.FlagSet) *flag.FlagSet {
	return fs
}

func (cmd *VersionCommand) Run(args []string) {
	fmt.Println("Now it is in the developping")
}

type ServerCommand struct {
	flagPort *int
}

func (cmd *ServerCommand) Flags(fs *flag.FlagSet) *flag.FlagSet {
	cmd.flagPort = fs.Int("p", 10443, "the port for listenning.")
	return fs
}

func (cmd *ServerCommand) Run(args []string) {

	network.InitHttpsHandlers()

	log.Print("Be listening on 10443.Go to https://127.0.0.1:",*cmd.flagPort)

	os.Chdir(os.Getenv("HOME"))

	err := http.ListenAndServeTLS(":"+Itoa(*cmd.flagPort), ".goreader/server.crt", ".goreader/server.key",nil)

	if err!=nil {
		log.Fatal(err)
	}
}
