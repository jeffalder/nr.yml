package nr_yml

import (
	"fmt"
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

	wd2, _ := os.Getwd()
	fmt.Println("new working dir:", wd2)

	cfg := new(newrelic.Config)
	ConfigFromDefaultYaml()(cfg)
	assert.NoError(t, cfg.Error)
	assert.Equal(t, "app-default", cfg.AppName)
	assert.Equal(t, "license-production", cfg.License)

}