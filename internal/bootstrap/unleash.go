package bootstrap

import (
	"net/http"

	"github.com/Unleash/unleash-go-sdk/v5"
)

type unleashConfig struct {
	AppName    string
	BackendUrl string
	Token      string
	Env        string
}

func initUnleash(cfg unleashConfig) error {
	opts := []unleash.ConfigOption{
		unleash.WithEnvironment(cfg.Env),
		unleash.WithAppName(cfg.AppName),
		unleash.WithUrl(cfg.BackendUrl),
		unleash.WithCustomHeaders(http.Header{
			"Authorization": []string{cfg.Token},
		}),
	}

	if cfg.Env != "production" {
		opts = append(opts, unleash.WithListener(&unleash.DebugListener{}))
	}

	if err := unleash.Initialize(opts...); err != nil {
		return err
	}

	// Block until the default client is ready
	unleash.WaitForReady()

	return nil
}
