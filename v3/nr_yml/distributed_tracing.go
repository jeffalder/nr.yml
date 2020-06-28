package nr_yml

import "github.com/newrelic/go-agent/v3/newrelic"

// Don't use this; it's only exported for the yaml parser
type DistributedTracingYaml struct {
	Enabled *bool `yaml:"enabled"`
}

func (distributedTracingYaml DistributedTracingYaml) update(cfg *newrelic.Config) {
	if distributedTracingYaml.Enabled != nil {
		newrelic.ConfigDistributedTracerEnabled(*distributedTracingYaml.Enabled)(cfg)
	}
}
