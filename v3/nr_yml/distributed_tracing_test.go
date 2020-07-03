//   Copyright 2020, Jeff Alder
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
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
