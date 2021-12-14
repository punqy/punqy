package http

type ModuleHttpHandlers interface {
}

type module struct {
}

func NewModule() ModuleHttpHandlers {
	var m module
	return &m
}
