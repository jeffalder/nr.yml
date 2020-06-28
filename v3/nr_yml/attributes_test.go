package nr_yml

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndividualLists(t *testing.T) {
	withContents(`
production:
  attributes:
    enabled: true
    include:
      - attrib-include-1
      - attrib-include-2
    exclude:
      - attrib-exclude-1
      - attrib-exclude-2
`, t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		ConfigFromYamlFile(filename)(cfg)
		assert.NoError(t, cfg.Error)
		assert.True(t, cfg.Attributes.Enabled)
		assert.Contains(t, cfg.Attributes.Include, "attrib-include-1")
		assert.Contains(t, cfg.Attributes.Include, "attrib-include-2")
		assert.Equal(t, 2, len(cfg.Attributes.Include))
		assert.Contains(t, cfg.Attributes.Exclude, "attrib-exclude-1")
		assert.Contains(t, cfg.Attributes.Exclude, "attrib-exclude-2")
		assert.Equal(t, 2, len(cfg.Attributes.Exclude))
	})
}
