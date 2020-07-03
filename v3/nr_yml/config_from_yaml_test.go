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
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestFileDoesNotExist(t *testing.T) {
	cfg := new(newrelic.Config)
	ConfigFromYamlFile("/tmp/file/does/not/exist/please/do/not/create/me")(cfg)
	assert.Error(t, cfg.Error)
}

func TestEmptyConfigFileProducesError(t *testing.T) {
	withContents("", t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		ConfigFromYamlFile(filename)(cfg)
		assert.Error(t, cfg.Error)
	})
}

func TestGarbageConfigFileProducesError(t *testing.T) {
	withContents("asdf", t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		ConfigFromYamlFile(filename)(cfg)
		assert.Error(t, cfg.Error)
	})
}

func TestFileWithOnlyEnvNoError(t *testing.T) {
	withContents("production:", t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		ConfigFromYamlFile(filename)(cfg)
		assert.NoError(t, cfg.Error)
	})
}

func TestFileWithWrongEnvError(t *testing.T) {
	withContents("other_env:", t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		ConfigFromYamlFile(filename)(cfg)
		assert.Error(t, cfg.Error)
	})
}

func TestFileWithNonDefaultEnvNoError(t *testing.T) {
	withContents("other_env:", t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		ConfigFromYamlFileEnvironment(filename, "other_env")(cfg)
		assert.NoError(t, cfg.Error)
	})
}

func TestFileWithAnchorAliasIncludes(t *testing.T) {
	withContents(`
common: &default
  app_name: app-default
  license_key: license-default
production:
  <<: *default
  license_key: license-production
`, t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		ConfigFromYamlFile(filename)(cfg)
		assert.NoError(t, cfg.Error)
		assert.Equal(t, "app-default", cfg.AppName)
		assert.Equal(t, "license-production", cfg.License)
	})
}

func TestDefaultFile(t *testing.T) {
	// create temp directory, queue for removal
	tempDir, err := ioutil.TempDir(os.TempDir(), "somedir*")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	tempFile := path.Join(tempDir, "newrelic.yml")
	err = ioutil.WriteFile(tempFile, []byte(`
common: &default
  app_name: app-default
  license_key: license-default
production:
  <<: *default
  license_key: license-production
`), 0644)
	assert.NoError(t, err)

	wd, _ := os.Getwd()
	assert.NoError(t, os.Chdir(tempDir))
	defer os.Chdir(wd)
	
	cfg := new(newrelic.Config)
	ConfigFromDefaultYaml()(cfg)
	assert.NoError(t, cfg.Error)
	assert.Equal(t, "app-default", cfg.AppName)
	assert.Equal(t, "license-production", cfg.License)

}