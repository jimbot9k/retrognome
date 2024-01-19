package types

type ApplicationConfiguration struct {
	AppName string `yaml:"app_name"`
	Port    int    `yaml:"port"`
}
