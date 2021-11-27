package config

type Gin struct {
	Addr          string
	Mode          string
	ResourcesPath string `mapstructure:"resources-path"`
}
