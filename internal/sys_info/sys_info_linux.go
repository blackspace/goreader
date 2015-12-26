package sys_info


import (
	"io/ioutil"
	"strings"
)

func Uptime() float64 {
	bs,_:=ioutil.ReadFile("/proc/uptime")
	content := string(bs)

	strings.s(content," ")

}
