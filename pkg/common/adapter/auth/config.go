package auth

import "time"

type RoleEnforcerConfig struct {
	ModelFile  string `yaml:"model_file"`
	PolicyFile string `yaml:"policy_file"`
}

type SslConfig struct {
	CertFile       string `yaml:"cert_file"`
	PrivateKeyFile string `yaml:"private_key_file"`
	PublicKeyFile  string `yaml:"public_key_file"`
}

type SessionConfig struct {
	Audience string        `yaml:"audience"`
	Issuer   string        `yaml:"issuer"`
	Duration time.Duration `yaml:"duration"`
}
