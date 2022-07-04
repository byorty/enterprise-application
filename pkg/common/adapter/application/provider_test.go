package application_test

import (
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/config"
)

type Config struct {
	A         A
	Partition int
}

type A struct {
	B string
	C struct {
		D bool
		F int
	}
}

func TestConfigProviderSuite(t *testing.T) {
	suite.Run(t, new(ConfigProviderSuite))
}

type ConfigProviderSuite struct {
	suite.Suite
}

func (s *ConfigProviderSuite) TestPopulate() {
	reader := strings.NewReader("a: {b: bar, c: {d: true, f: 12}}")

	provider, err := application.NewProviderByOptions(config.Source(reader))
	s.Nil(err)

	var a A
	s.Nil(provider.PopulateByKey("a", &a))
	s.Equal("bar", a.B)
	s.Equal(true, a.C.D)
	s.Equal(12, a.C.F)

	var f int
	s.Nil(provider.PopulateByKey("a.c.f", &f))
	s.Equal(12, f)

	var cfg Config
	s.Nil(provider.Populate(&cfg))
	s.Equal(cfg.A, a)
	s.Equal(cfg.A.B, a.B)
}

func (s *ConfigProviderSuite) TestExpand() {
	var a int
	var b string
	var c string
	varB := "hello world"
	err := os.Setenv("VAR_B", varB)
	if err != nil {
		s.Error(err)
	}

	reader := strings.NewReader(`
a: 1
b: "$VAR_B"
c: "$VAR_C"
`)

	provider, err := application.NewProviderByOptions(config.Source(reader))
	s.Nil(err)

	s.Nil(provider.PopulateByKey("a", &a))
	s.Equal(1, a)

	s.Nil(provider.PopulateByKey("b", &b))
	s.Equal(varB, b)

	s.Nil(provider.PopulateByKey("b", &b))
	s.Equal("", c)
}
