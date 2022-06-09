package config

type Environment struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type App struct {
	Owner string `yaml:"owner"`
	Repo  string `yaml:"repo"`
	Tag   string `yaml:"tag"`

	Environments []Environment `yaml:"environments"`
	Dependencies []App         `yaml:"dependencies"`
}
