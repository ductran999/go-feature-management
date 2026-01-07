package config

type EnvConfig struct {
	App     App              `mapstructure:"app"`
	DB      PostgreSQLConfig `mapstructure:"db"`
	Unleash Unleash          `mapstructure:"unleash"`
}

type App struct {
	Environment string `mapstructure:"environment"`
	Name        string `mapstructure:"name"`
}

type PostgreSQLConfig struct {
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	User    string `mapstructure:"user"`
	Pass    string `mapstructure:"pass"`
	DatName string `mapstructure:"dat_name"`
}

type Unleash struct {
	Token      string `mapstructure:"token"`
	BackendUrl string `mapstructure:"backend_url"`
}
