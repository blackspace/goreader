package main

import (
	"os"
	"io"
	"encoding/json"
	"log"
	"goreader/internal/sys_info"
)

func main() {
	var paths []string

	if err:=json.NewDecoder(os.Stdin).Decode(&paths);err==io.EOF {
	} else if err!=nil {
		log.Fatal(err)
	}

	result := getKVsByPaths(paths)

	json.NewEncoder(os.Stdout).Encode(result)
}

func getKVsByPaths(paths []string) map[string]interface{} {
	result :=make(map[string]interface{})

	for _,p := range paths {
		log.Print("Search the ",p)
		f:=sys_info.GetFunc(p)

		if f!=nil {
			v:=f()
			log.Print("Get the handler func")
			result[p]=v
		}
	}

	return result
}

