package config

type ModuleConfig interface {
	Config() Config
}

type module struct {
	config Config
}

func (m *module) Config() Config {
	return m.config
}

func NewModule() ModuleConfig {
	return &module{
		config: Load(),
	}
}
