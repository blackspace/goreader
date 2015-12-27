package compile

import (
	"bytes"
	"os/exec"
	"log"
	"errors"
)


func CompileDaemon() error {
	var err_buf bytes.Buffer
	cmd := exec.Command("go","install","github.com/blackspace/goreader/internal/goreader_daemon")
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


