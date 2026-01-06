package unleash

import (
	"feature-flag-poc/internal/application/port"

	"github.com/Unleash/unleash-go-sdk/v5"
)

type unleashFeatureFlag struct{}

func NewUnleashFeatureFlag() port.FeatureFlag {
	return &unleashFeatureFlag{}
}

func (f *unleashFeatureFlag) IsEnabled(key string) bool {
	return unleash.IsEnabled(key)
}
