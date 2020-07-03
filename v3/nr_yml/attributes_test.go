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
