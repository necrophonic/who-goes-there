package rules

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type (
	// Rules contains a set of rules to be executed
	Rules struct {
		User []Rule `yaml:"user"`
		Team []Rule `yaml:"team"`
	}

	// Rule represents a single rule type to execute
	Rule struct {
		// Name and description of the rule. These show up in the reports
		Name        string `yaml:"name"`
		Description string `yaml:"description"`

		// Action to perform if the rule is failed
		Action string `yaml:"action"`
	}
)

// New returns an empty rules instance
func New() *Rules {
	return &Rules{}
}

// NewFromYAML instantiates a new ruleset from a yaml configuration
func NewFromYAML(yml []byte) (*Rules, error) {
	var rules Rules
	err := yaml.Unmarshal(yml, &rules)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse ruleset")
	}
	return &rules, nil
}
