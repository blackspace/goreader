package main

import (
	"os"
	"io"
	"encoding/json"
	"log"
	. "goreader/internal/handler"
	. "goreader/internal/sys_info"
)

func init() {
	RegistHandlers()
}

func main() {
	var paths []string

	if err:=json.NewDecoder(os.Stdin).Decode(&paths);err==io.EOF {
	} else if err!=nil {
		log.Fatal(err)
	}

	result := getKVsByPaths(paths)

	json.NewEncoder(os.Stdout).Encode(result)
}


func RegistHandlers() {
	RegistHandler(Handler{"/uptime","uptime",func() interface{} {
		return Uptime()
	}})

	RegistHandler(Handler{"/cpu","cpu",func() interface{} {
		return 12.3
	}})
}


func getKVsByPaths(paths []string) map[string]interface{} {
	result :=make(map[string]interface{})

	for _,p := range paths {
		 for _,h :=range GetHandlerSet() {
			 if h.Alias==p || h.Path==p {
				 v := h.Func()

				result[p]=v
			 }
		 }
	}

	return result
}

