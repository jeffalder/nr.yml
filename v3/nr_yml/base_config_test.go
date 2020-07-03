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
	"fmt"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndividualItems(t *testing.T) {
	withContents(`
production:
  agent_enabled: false
  high_security: true
  security_policies_token: ffff-0000-ffff-0000
  host: staging-collector.example.com
  labels:
    label1: value1
    label2: value2
`, t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		ConfigFromYamlFile(filename)(cfg)
		assert.NoError(t, cfg.Error)
		assert.False(t, cfg.Enabled)
		assert.True(t, cfg.HighSecurity)
		assert.Equal(t, "ffff-0000-ffff-0000", cfg.SecurityPoliciesToken)
		assert.Equal(t, "staging-collector.example.com", cfg.Host)
		assert.Equal(t, "value1", cfg.Labels["label1"])
		assert.Equal(t, "value2", cfg.Labels["label2"])
	})
}

type MockLogger struct{}

func (MockLogger) Error(msg string, context map[string]interface{}) {}
func (MockLogger) Warn(msg string, context map[string]interface{})  {}
func (MockLogger) Info(msg string, context map[string]interface{})  {}
func (MockLogger) Debug(msg string, context map[string]interface{}) {}
func (MockLogger) DebugEnabled() bool {
	return false
}

func TestLogStream(t *testing.T) {
	for _, logStream := range [...]string{"STDOUT", "Stdout", "stdout", "STDERR", "Stderr", "stderr"} {
		withContents(fmt.Sprintf(`
production:
  log_stream_name: %s
`, logStream), t, func(filename string, t *testing.T) {
			cfg := new(newrelic.Config)
			mockLogger := new(MockLogger)
			cfg.Logger = mockLogger
			ConfigFromYamlFile(filename)(cfg)
			assert.NoError(t, cfg.Error)
			assert.NotEqual(t, mockLogger, cfg.Logger)
			assert.False(t, cfg.Logger.DebugEnabled())
		})
	}
}

func TestDebugLogging(t *testing.T) {
	for _, logLevel := range [...]string{"DEBUG", "Debug", "debug"} {
		withContents(fmt.Sprintf(`
production:
  log_stream_name: STDOUT
  log_level: %s
`, logLevel), t, func(filename string, t *testing.T) {
			cfg := new(newrelic.Config)
			ConfigFromYamlFile(filename)(cfg)
			assert.NoError(t, cfg.Error)
			assert.True(t, cfg.Logger.DebugEnabled())
		})
	}
}

func TestOtherLogLevelSettingIsNonDebug(t *testing.T) {
	withContents(`
production:
  log_stream_name: STDOUT
  log_level: info
		`, t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		ConfigFromYamlFile(filename)(cfg)
		assert.NoError(t, cfg.Error)
		assert.False(t, cfg.Logger.DebugEnabled())
	})
}
