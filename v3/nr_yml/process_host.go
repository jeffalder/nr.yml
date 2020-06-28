package nr_yml

import "github.com/newrelic/go-agent/v3/newrelic"

// Don't use this; it's only exported for the yaml parser
type ProcessHostYaml struct {
	DisplayName *string `yaml:"display_name"`
}

func (processHostYaml ProcessHostYaml) update(cfg *newrelic.Config) {
	if processHostYaml.DisplayName != nil {
		cfg.HostDisplayName = *processHostYaml.DisplayName
	}
}
