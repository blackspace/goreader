package subcommands
import (
	"goreader/network"
	"log"
	"os"
	"net/http"
	"flag"
	. "strconv"
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

	log.Print("Be listening on 10443.Go to https://127.0.0.1:",*cmd.flagPort)

	os.Chdir(os.Getenv("HOME"))

	err := http.ListenAndServeTLS(":"+Itoa(*cmd.flagPort), ".goreader/server.crt", ".goreader/server.key",nil)

	if err!=nil {
		log.Fatal(err)
	}
}
