package sys_info


import (
	"io/ioutil"
	"strings"
	. "strconv"
	"time"
	"os"
	"bufio"
	"log"
	"io"
)

func RegistHandlers() {
	RegistAction(Action{"/uptime","uptime",uptime,"the uptime"})

	RegistAction(Action{"/version","version",version,"the version"})

	RegistAction(Action{"/now","now",now,"the now"})

	RegistAction(Action{"/users","users",users,"The users of system"})

	RegistAction(Action{"/memory","memory",memory,"The memory information of system,its unit is K"})
}


func uptime() interface{} {
	bs,_:=ioutil.ReadFile("/proc/uptime")

	f,_:=ParseFloat(strings.Split(string(bs)," ")[0],64)

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
	result := make([]string,0,1<<8)

	if f,err:=os.Open("/etc/passwd"); err!=nil {
		log.Fatal("Failing to open /etc/passwd")
	} else {
		r:=bufio.NewReader(f)

		for {
			l,_,err:=r.ReadLine()
			if err==io.EOF {
				break
			}
			result=append(result,strings.Split(string(l),":")[0])
		}

	}

	return result
}

func memory() interface{} {
	result := make(map[string]int)

	if f,err:=os.Open("/proc/meminfo");err!=nil {
		log.Fatal("Failing to open /proc/meminfo")
	} else {
		r:=bufio.NewReader(f)

		for {
			l,_,err:=r.ReadLine()

			if err==io.EOF {
				break
			}

			k:=strings.Split(string(l),":")[0]
			r := strings.TrimSpace(strings.Split(string(l),":")[1])
			v,_:=Atoi(strings.Split(r," ")[0])


			result[k]=v
		}
	}

	return result
}
