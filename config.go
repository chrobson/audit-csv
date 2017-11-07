package main

// Config for server.
type Config struct {
	Server struct {
		Address   string `yaml:"address"`    // address of server to bind
		AuthKey   string `yaml:"auth_key"`   // auth key for clients
		TLS       bool   `yaml:"tls"`        // should server run with tls
		StaticDir string `yaml:"static_dir"` // directories with static files
	} `yaml:"server"`
}

// newDefaultConfig returns config with set defaults.
func newDefaultConfig() *Config {
	cfg := &Config{}
	cfg.Server.Address = ":8080"
	return cfg
}
