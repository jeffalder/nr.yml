package nr_yml

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistributedTracingEnabled(t *testing.T) {
	withContents(`
production:
  distributed_tracing:
    enabled: true
`, t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		cfg.DistributedTracer.Enabled = false

		ConfigFromYamlFile(filename)(cfg)
		assert.NoError(t, cfg.Error)
		assert.True(t, cfg.DistributedTracer.Enabled)
	})
}

func TestDistributedTracingNoOverwrite(t *testing.T) {
	withContents("production:", t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		cfg.DistributedTracer.Enabled = true

		ConfigFromYamlFile(filename)(cfg)
		assert.NoError(t, cfg.Error)
		assert.True(t, cfg.DistributedTracer.Enabled)
	})
}
