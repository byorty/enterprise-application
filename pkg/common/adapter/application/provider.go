package application

import (
	"github.com/joho/godotenv"
	"go.uber.org/config"
	"os"
)

type Provider interface {
	Populate(interface{}) error
	PopulateByKey(string, interface{}) error
}

func NewFxProvider(args Arguments) (Provider, error) {
	return NewProviderByOptions(config.File(args.ConfigFilename))
}

func NewProviderByOptions(options ...config.YAMLOption) (Provider, error) {
	dotenvFile, _ := godotenv.Read(".env")
	configProvider, err := config.NewYAML(append(
		[]config.YAMLOption{
			config.Expand(expandFunc(dotenvFile)),
		},
		options...,
	)...)
	if err != nil {
		return nil, err
	}

	return &provider{configProvider}, nil
}

func expandFunc(extraEnv map[string]string) func(key string) (val string, ok bool) {
	return func(key string) (val string, ok bool) {
		val, ok = os.LookupEnv(key)
		if !ok {
			val, ok = extraEnv[key]
			if !ok {
				val = ""
				ok = true
			}
		}
		return val, ok
	}
}

type provider struct {
	provider config.Provider
}

func (p *provider) Populate(target interface{}) error {
	return p.PopulateByKey("", target)
}

func (p *provider) PopulateByKey(key string, target interface{}) error {
	return p.provider.Get(key).Populate(target)
}
