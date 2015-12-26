package handler

type Handler struct {
	Path string
	Alias string
	Func func() interface{}
}

var handler_set = make([]Handler,0,1<<8)

func GetHandlerSet() []Handler {
	return handler_set
}

func RegistHandler(h Handler) {
	handler_set =append(handler_set,h)
}

