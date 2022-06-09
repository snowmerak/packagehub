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

func NewApp(owner, repo, tag string) *App {
	return &App{
		Owner: owner,
		Repo:  repo,
		Tag:   tag,
	}
}

func (a *App) AddEnvironment(key, value string) {
	a.Environments = append(a.Environments, Environment{
		Key:   key,
		Value: value,
	})
}

func (a *App) AddDependency(owner, repo, tag string) {
	a.Dependencies = append(a.Dependencies, App{
		Owner: owner,
		Repo:  repo,
		Tag:   tag,
	})
}
