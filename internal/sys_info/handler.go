package sys_info

type Handler func() interface{}

type Action struct {
	Path    string
	Alias   string
	Handler Handler
}

var actions = make([]Action,0,1<<8)

func GetActions() []Action {
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
