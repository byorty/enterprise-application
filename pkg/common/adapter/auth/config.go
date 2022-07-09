package auth

type RoleEnforcerConfig struct {
	ModelFile  string `yaml:"model_file"`
	PolicyFile string `yaml:"policy_file"`
}

type SslConfig struct {
	PrivateKeyFile string `yaml:"private_key_file"`
	PublicKeyFile  string `yaml:"public_key_file"`
}
