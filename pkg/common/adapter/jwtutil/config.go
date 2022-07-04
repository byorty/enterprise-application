package jwtutil

type Config struct {
	PrivateKeyFile string `yaml:"private_key_file"`
	PublicKeyFile  string `yaml:"public_key_file"`
}
