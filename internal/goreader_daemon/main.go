package main

import (
	"os"
	"io"
	"encoding/json"
	"log"
	"github.com/blackspace/goreader/internal/sys_info"
)

func main() {
	json.NewEncoder(os.Stdout).Encode(getResultOut(getInputPaths()))
}

func getInputPaths() []interface{} {

	var paths []interface{}

	if err:=json.NewDecoder(os.Stdin).Decode(&paths);err==io.EOF {
	} else if err!=nil {
		log.Fatal(err)
	}

	return paths
}

func getResultOut(paths []interface{}) map[string]interface{} {
	result :=make(map[string]interface{})

	for _,p := range paths {
		switch lp :=p.(type) {
		case string:
			f:=sys_info.GetHandler(lp)
			if f!=nil {
				v:=f()
				result[lp]=v
			}
		case map[string]interface{}:
			for lk,lv:=range lp {
				f:=sys_info.GetHandler(lk)

				log.Printf("%#v",lv)

				if f!=nil {
					params:=make([]string,0,1<<3)

					for _,llv:=range lv.([]interface{}) {
						params=append(params,llv.(string))
					}

					v:=f(params...)

					result[lk]=v
				} else {
					continue
				}

			}

		}

	}

	return result
}

