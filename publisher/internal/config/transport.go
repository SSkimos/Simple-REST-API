package config

type Router struct {
	IP    string `yaml:"address"`
	Port  string `yaml:"port"`
	Debug bool   `yaml:"debug"`
}
