package config

type Config struct {
	Logger  `yaml:"logger"`
	Router  `yaml:"transport"`
	Service `yaml:"service"`
}
