package nr_yml

import (
	"errors"
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/newrelic/go-agent/v3/newrelic"
	"io/ioutil"
	"os"
)

// Uses `newrelic.yml` in the current directory, plus the environment defined by `NEW_RELIC_ENVIRONMENT`.
// if `NEW_RELIC_ENVIRONMENT` is not defined or is empty, the literal string `production` is used.
func ConfigFromDefaultYaml() newrelic.ConfigOption {
	return ConfigFromYaml("newrelic.yml")
}

// Uses the supplied filename, plus the environment defined by `NEW_RELIC_ENVIRONMENT`.
// if `NEW_RELIC_ENVIRONMENT` is not defined or is empty, the literal string `production` is used.
func ConfigFromYaml(filename string) newrelic.ConfigOption {
	return ConfigFromYamlEnvironment(filename, os.Getenv("NEW_RELIC_ENVIRONMENT"))
}

// Uses the supplied filename and environment.
func ConfigFromYamlEnvironment(filename string, environment string) newrelic.ConfigOption {
	return func(cfg *newrelic.Config) {
		dat, err := ioutil.ReadFile(filename)
		if nil != err {
			cfg.Error = err
			return
		}

		var yamlEnvs map[string]ConfigYaml

		if err := yaml.Unmarshal(dat, &yamlEnvs); nil != err {
			cfg.Error = err
			return
		}

		if "" == environment {
			environment = "production"
		}

		if yamlConfig, ok := yamlEnvs[environment]; ok {
			yamlConfig.update(cfg)
			return
		}

		cfg.Error = errors.New(fmt.Sprintf("Environment %s was not defined in %s.", environment, filename))
	}
}
