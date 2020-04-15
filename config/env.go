package config

func WithEnvConfig(getEnv func(key string) string) ReadConfigKeyFunc {
	return func(key string) string {
		return getEnv(key)
	}
}
