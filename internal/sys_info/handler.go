package sys_info

type Func func() interface{}

type Handler struct {
	Path string
	Alias string
	Func Func
}

var handler_set = make([]Handler,0,1<<8)

func GetHandlerSet() []Handler {
	return handler_set
}

func RegistHandler(h Handler) {
	handler_set =append(handler_set,h)
}

func GetFunc(path string) Func {
	for _,h :=range GetHandlerSet() {
		if h.Alias==path || h.Path==path {
			return h.Func
		}
	}

	return nil
}
