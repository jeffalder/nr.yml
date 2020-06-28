package nr_yml

import "github.com/newrelic/go-agent/v3/newrelic"

// Don't use this; it's only exported for the yaml parser
type AttributesYaml struct {
	Enabled *bool     `yaml:"enabled"`
	Include *[]string `yaml:"include"`
	Exclude *[]string `yaml:"exclude"`
}

func (attributesYaml AttributesYaml) update(cfg *newrelic.Config) {
	if attributesYaml.Enabled != nil {
		cfg.Attributes.Enabled = *attributesYaml.Enabled
	}

	if attributesYaml.Include != nil {
		cfg.Attributes.Include = *attributesYaml.Include
	}

	if attributesYaml.Exclude != nil {
		cfg.Attributes.Exclude = *attributesYaml.Exclude
	}
}
