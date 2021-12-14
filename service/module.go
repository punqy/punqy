package service

type ModuleService interface {
}

type module struct {
}

func NewModule() ModuleService {
	var m module
	return &m
}
