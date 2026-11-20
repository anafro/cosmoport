package config

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/anafro/cosmoport/internal/config/errors"
	"github.com/goccy/go-yaml"
)

type RuleDestination struct {
	Type  string
	Value string
}

type Rule struct {
	Path    string
	Https   string `yaml:"https"`
	Domain  string `yaml:"domain"`
	Backend string `yaml:"backend"`
	Static  string `yaml:"static"`
}

type Config struct {
	Rules map[string]Rule `yaml:"rules"`
}

func Load(path string) (*Config, error) {
	config := &Config{}
	contentsBytes, err := os.ReadFile(path)

	if err != nil {
		return nil, errors.NewErrCannotReadConfigFile(path, err)
	}

	if !utf8.Valid(contentsBytes) {
		return nil, errors.NewErrNonUtf8ConfigFile(path)
	}

	err = yaml.Unmarshal(contentsBytes, config)

	if err != nil {
		return nil, errors.NewMalformedConfigFile(path, err)
	}

	return config, nil
}

func (r *Rule) Destination() (*RuleDestination, error) {
	if r.Backend != "" {
		return &RuleDestination{
			Type:  "backend",
			Value: r.Backend,
		}, nil
	}

	if r.Static != "" {
		return &RuleDestination{
			Type:  "static",
			Value: r.Static,
		}, nil
	}

	return nil, errors.NewErrMalformedDestination()
}

func (c *Config) Sprint() (string, error) {
	var formatted strings.Builder

	for name, rule := range c.Rules {
		destination, err := rule.Destination()

		if err != nil {
			return "", err
		}

		formatted.WriteString(fmt.Sprintf(
			"%s: %s -> (%s) %s\n",
			name,
			rule.Domain,
			destination.Type,
			destination.Value,
		))
	}

	return formatted.String(), nil
}
