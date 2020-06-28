package nr_yml


import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessHostSet(t *testing.T) {
	withContents(`
production:
  process_host:
    display_name: phdn-yes
`, t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		cfg.DistributedTracer.Enabled = false

		ConfigFromYamlFile(filename)(cfg)
		assert.NoError(t, cfg.Error, "did not expect error")
		assert.Equal(t, "phdn-yes", cfg.HostDisplayName)
	})
}

func TestProcessHostNoOverwrite(t *testing.T) {
	withContents(`
production:
  process_host:
	`, t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		cfg.HostDisplayName = "~~ hello ~~"
		ConfigFromYamlFile(filename)(cfg)
		assert.NoError(t, cfg.Error, "Did not expect error")
		assert.Equal(t, "~~ hello ~~", cfg.HostDisplayName)
	})
}

