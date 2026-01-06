package config

import (
	"log"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

const configPath = "./configs/config.yml"

func LoadEnv() (*EnvConfig, error) {
	var env EnvConfig

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("[WARN] failed to read config from file use current shell env", err)
		autoBindEnv(viper.GetViper(), EnvConfig{})
	}

	if err := viper.Unmarshal(&env); err != nil {
		return nil, err
	}

	return &env, nil
}

func autoBindEnv(v *viper.Viper, s any) {
	bindStruct(v, reflect.TypeOf(s), "")
}

func bindStruct(v *viper.Viper, t reflect.Type, prefix string) {
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if !f.IsExported() {
			continue
		}

		tag := f.Tag.Get("mapstructure")
		if tag == "" || tag == "-" {
			continue
		}

		key := tag
		if prefix != "" {
			key = prefix + "." + tag
		}

		if f.Type.Kind() == reflect.Struct {
			bindStruct(v, f.Type, key)
			continue
		}

		_ = v.BindEnv(key)
	}
}
