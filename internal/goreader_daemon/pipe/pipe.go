package pipe

import (
	"os/exec"
	"log"
	"bytes"
	"strings"
)

func DealIt(input []byte) (output []byte) {
	log.Print("The requst is ",string(input))

	var input_buf=bytes.NewBuffer(input)
	var out_buf=new(bytes.Buffer)
	var err_buf=new(bytes.Buffer)

	cmd := exec.Command("goreader_daemon")
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
