package subcommands
import (
	"flag"
	"os"
	"os/exec"
)

type InstallCommand struct{}

func (cmd *InstallCommand) Flags(fs *flag.FlagSet) *flag.FlagSet {
	return fs
}

func (cmd *InstallCommand) Run(args []string) {
	home:=os.Getenv("HOME")
	path := home+"/.goreader"

	if fileinfo,_:=os.Stat(path); fileinfo==nil {
		os.Mkdir(path,0755)
	}


	os.Chdir(path)

	exec.Command("openssl","genrsa","-out","key.pem","2048").Run()

	exec.Command("openssl","req","-new","-x509","-key","key.pem","-out","cert.pem","-days","3650","-batch").Run()

	os.Chdir(home)

}
