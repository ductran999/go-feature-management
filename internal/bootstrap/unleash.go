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

func initUnleash(config unleashConfig) error {
	err := unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithEnvironment(config.Env),
		unleash.WithAppName(config.AppName),
		unleash.WithUrl(config.BackendUrl),
		unleash.WithCustomHeaders(http.Header{"Authorization": {config.Token}}),
	)
	if err != nil {
		return err
	}

	// Note this will block until the default client is ready
	unleash.WaitForReady()

	return nil
}
