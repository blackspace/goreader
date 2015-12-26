package daemon

import (
	"bytes"
	"os/exec"
	"log"
	"errors"
	"strings"
)


func CompileDaemon() error {
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


func DealIt(input []byte) (output []byte) {
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