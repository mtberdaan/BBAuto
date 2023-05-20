package config

// app config interface
type Config interface {
	GetResourceCfgs() map[string]ResourceConfig
}

// config for resources
type ResourceConfig struct {
	ResourceName     string `yaml:"resource_name"`
	RequestUrl       string `yaml:"request_url"`
	RequestVar       string `yaml:"request_var"`
	RequestRateLimit int    `yaml:"request_rate_limit"`
}
