package network

import (
	"net/http"
	"log"
	. "goreader/internal/goreader_daemon/compile"
	. "goreader/internal/goreader_daemon/pipe"
	"io"
)

func list_handler(w http.ResponseWriter,r *http.Request){

}

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
	http.HandleFunc("/list",list_handler)
	http.HandleFunc("/", root_handler)
}
