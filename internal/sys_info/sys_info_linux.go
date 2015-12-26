package sys_info


import (
	"io/ioutil"
	"strings"
	. "strconv"
	"time"
)

func RegistHandlers() {
	RegistAction(Action{"/uptime","uptime",uptime})

	RegistAction(Action{"/version","version",version})

	RegistAction(Action{"/now","now",now})
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

func users() interface{} {
	return nil
}
