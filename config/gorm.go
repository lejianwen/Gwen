package config

type Mysql struct {
	Dns          string
	MaxIdleConns int `mapstructure:"max-idle-conns"`
	MaxOpenConns int `mapstructure:"max-open-conns"`
}
