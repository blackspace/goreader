package sys_info


import (
	"io/ioutil"
	"strings"
	. "strconv"
	"time"
)

func RegistHandlers() {
	RegistHandler(Handler{"/uptime","uptime",uptime})

	RegistHandler(Handler{"/cpu","cpu",func() interface{} {
		return 12.3
	}})

	RegistHandler(Handler{"/version","version",version})

	RegistHandler(Handler{"/now","now",now})
}


func uptime() interface{} {
	bs,_:=ioutil.ReadFile("/proc/uptime")
	content := string(bs)

	f,_:=ParseFloat(strings.Split(content," ")[0],64)

	return f
}

func version() interface{} {
	v,_:=ioutil.ReadFile("/proc/version")
	return string(v)
}

func now() interface{} {
	return time.Now().Unix()
}
