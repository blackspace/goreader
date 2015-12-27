package sys_info

type Handler func() interface{}

type Action struct {
	Path    string
	Alias   string
	Handler Handler
	Descript string
}

var actions = make([]Action,0,1<<8)

func GetActions() []Action {
	RegistAction(Action{Path: "/actions",
		Alias:"actions",
		Handler:Actions,
		Descript:"Get all status actionses"})
	return actions
}

func RegistAction(a Action) {
	actions =append(actions, a)
}

func GetHandler(path string) Handler {
	for _, a :=range GetActions() {
		if a.Alias==path || a.Path==path {
			return a.Handler
		}
	}

	return nil
}

func Actions() interface{} {
	result :=make([]interface{},0,1<<8)
	for _,a:=range actions {
		result=append(result, struct{Path string;Alias string;Descript string}{a.Path,a.Alias,a.Descript})
	}
	return result
}
