package config

var Env EnvConfig

type EnvConfig struct {
	Environment string           `mapstructure:"environment"`
	DBConfig    PostgreSQLConfig `mapstructure:"db"`
}

type PostgreSQLConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
}
