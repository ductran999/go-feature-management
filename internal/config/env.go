package config

type EnvConfig struct {
	Environment string           `mapstructure:"environment"`
	DB          PostgreSQLConfig `mapstructure:"db"`
}

type PostgreSQLConfig struct {
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	User    string `mapstructure:"user"`
	Pass    string `mapstructure:"pass"`
	DatName string `mapstructure:"dat_name"`
}
