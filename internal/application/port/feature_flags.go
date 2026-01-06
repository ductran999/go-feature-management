package port

type FeatureFlag interface {
	IsEnabled(key string) bool
}
