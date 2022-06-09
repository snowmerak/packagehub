package config

import (
	"fmt"
)

type Config struct {
	Apps []App `yaml:"apps"`
}

func NewConfig() Config {
	return Config{}
}

func (c *Config) AppendApp(owner, repo, tag string) error {
	for _, app := range c.Apps {
		if app.Owner == owner && app.Repo == repo {
			return fmt.Errorf("Config.AppendApp: app already exists")
		}
	}
	c.Apps = append(c.Apps, App{
		Owner: owner,
		Repo:  repo,
		Tag:   tag,
	})
	return nil
}

func (c *Config) SetTag(owner, repo, tag string) error {
	for _, app := range c.Apps {
		if app.Owner == owner && app.Repo == repo {
			app.Tag = tag
			return nil
		}
	}
	return fmt.Errorf("Config.SetTag: app not found")
}

func (c *Config) GetTag(owner, repo string) (string, error) {
	for _, app := range c.Apps {
		if app.Owner == owner && app.Repo == repo {
			return app.Tag, nil
		}
	}
	return "", fmt.Errorf("Config.GetTag: app not found")
}

func (c *Config) ContainsApp(owner, repo string) bool {
	for _, app := range c.Apps {
		if app.Owner == owner && app.Repo == repo {
			return true
		}
	}
	return false
}

func (c *Config) RemoveApp(owner, repo string) error {
	for i, app := range c.Apps {
		if app.Owner == owner && app.Repo == repo {
			c.Apps = append(c.Apps[:i], c.Apps[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Config.RemoveApp: app not found")
}
