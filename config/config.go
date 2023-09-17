package config

type Config struct {
	AllowedDomains []string
	StartingUrl    string
	CacheDir       string
}

func New(allowedDomains []string, startingUrl string, cacheDir string) *Config {
	return &Config{
		allowedDomains, 
		startingUrl, 
		cacheDir,
	}
}
