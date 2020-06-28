package nr_yml

import (
	"fmt"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtilizationAll(t *testing.T) {
	withContents(`
production:
  utilization:
    billing_hostname: host123
    total_ram_mib: 2344
    logical_processors: 15
`, t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		ConfigFromYamlFile(filename)(cfg)
		if cfg.Error != nil {
			t.Fatal("specified env should not generate error", cfg.Error)
		}
		assert.Equal(t, "host123", cfg.Utilization.BillingHostname)
		assert.Equal(t, 2344, cfg.Utilization.TotalRAMMIB)
		assert.Equal(t, 15, cfg.Utilization.LogicalProcessors)
	})
}

func TestDetectClouds(t *testing.T) {
	for _, cloud := range [...]string{"aws", "azure", "docker", "gcp", "kubernetes", "pcf"} {
		withContents(fmt.Sprintf(`
production:
  utilization:
    detect_%s: true
`, cloud), t, func(filename string, t *testing.T) {
			cfg := new(newrelic.Config)
			ConfigFromYamlFile(filename)(cfg)
			assert.Equal(t, cloud == "aws", cfg.Utilization.DetectAWS)
			assert.Equal(t, cloud == "azure", cfg.Utilization.DetectAzure)
			assert.Equal(t, cloud == "docker", cfg.Utilization.DetectDocker)
			assert.Equal(t, cloud == "gcp", cfg.Utilization.DetectGCP)
			assert.Equal(t, cloud == "kubernetes", cfg.Utilization.DetectKubernetes)
			assert.Equal(t, cloud == "pcf", cfg.Utilization.DetectPCF)
		})
	}
}

