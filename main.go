package main

import (
	"net/http"
	"log"
	"os/exec"
	"bytes"
	"io"
	"errors"
	"strings"
	"os"
	"flag"
	"github.com/rakyll/command"
	. "strconv"
	"fmt"
)

func handler(w http.ResponseWriter,r *http.Request) {
	log.Println("It is a request be coming in")

	w.Header().Set("Content-Type","Content-Type: text/plain; charset=utf-8")

	if err := _GetNewDaemon(); err!=nil {
		w.Write([]byte(err.Error()))
		return
	}


	content := make([]byte,0,r.ContentLength)

	for {

		temp := make([]byte,r.ContentLength)

		if n,err:=r.Body.Read(temp) ; err==io.EOF {
			content=append(content,temp[:n]...)
			break
		} else {
			content=append(content,temp[:n]...)
		}
	}

	response := _Daemon(content)

	w.Write(response)
}

func _GetNewDaemon() error {
	var err_buf bytes.Buffer
	cmd := exec.Command("go","install","goreader/internal/goreader-daemon")
	cmd.Stderr=&err_buf

	if err := cmd.Run();err != nil {
		log.Print(err)
		log.Print(err_buf.String())

		return errors.New("Compiling the daemon is wrong")

	} else {
		log.Println("It has maked a new goreader-daemon")
		return nil
	}
}

func _Daemon(input []byte) (output []byte) {
	log.Print("The requst is ",string(input))

	var input_buf=bytes.NewBuffer(input)
	var out_buf=new(bytes.Buffer)
	var err_buf=new(bytes.Buffer)

	cmd := exec.Command("goreader-daemon")
	cmd.Stdin = input_buf
	cmd.Stdout = out_buf
	cmd.Stderr = err_buf

	if err :=cmd.Run();err != nil {
		result :=strings.Join(strings.Split(err_buf.String()," ")[2:]," ")

		log.Print(result)

		return []byte(result)
	} else {
		if err_buf.Len()>0 {
			log.Print("The error of output is ",string(err_buf.Bytes()))
		}

		log.Print("The response is ",out_buf.String())

		return out_buf.Bytes()
	}

}


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
	flagDaemon *bool
}

func (cmd *ServerCommand) Flags(fs *flag.FlagSet) *flag.FlagSet {
	cmd.flagPort = fs.Int("p", 10443, "the port for listenning.")
	return fs
}

func (cmd *ServerCommand) Run(args []string) {
	http.HandleFunc("/",handler)

	log.Print("Be listening on 10443.Go to https://127.0.0.1:",*cmd.flagPort)

	os.Chdir(os.Getenv("HOME"))

	err := http.ListenAndServeTLS(":"+Itoa(*cmd.flagPort), ".goreader/server.crt", ".goreader/server.key",nil)

	if err!=nil {
		log.Fatal(err)
	}
}


func main() {

	command.On("version", "prints the version", &VersionCommand{}, nil)
	command.On("server", "start the server for listenning", &ServerCommand{}, nil)

	command.Parse()
	command.Run()

}
