package network

import (
	"net/http"
	"log"
	. "github.com/blackspace/goreader/internal/goreader_daemon/compile"
	. "github.com/blackspace/goreader/internal/goreader_daemon/pipe"
	"io"
	"os"
	. "strconv"
)

func root_handler(w http.ResponseWriter,r *http.Request)  {
	log.Println("It is a request be coming in")

	w.Header().Set("Content-Type","Content-Type: text/plain; charset=utf-8")

	if err := CompileDaemon(); err!=nil {
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

	response := DealIt(content)

	w.Write(response)
}

func LoadHttpsHandlers() {
	http.HandleFunc("/", root_handler)
}

func Listen(port int) {
	log.Print("Be listening on 10443.Go to https://127.0.0.1:",port)

	os.Chdir(os.Getenv("HOME"))

	err := http.ListenAndServeTLS(":"+Itoa(port), ".goreader/cert.pem", ".goreader/key.pem",nil)

	if err!=nil {
		log.Fatal(err)
	}
}
