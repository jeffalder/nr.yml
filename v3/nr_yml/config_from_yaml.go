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
	return ConfigFromYamlFile("newrelic.yml")
}

// Uses the supplied filename, plus the environment defined by `NEW_RELIC_ENVIRONMENT`.
// if `NEW_RELIC_ENVIRONMENT` is not defined or is empty, the literal string `production` is used.
func ConfigFromYamlFile(filename string) newrelic.ConfigOption {
	return ConfigFromYamlFileEnvironment(filename, os.Getenv("NEW_RELIC_ENVIRONMENT"))
}

// Uses the supplied filename and environment.
func ConfigFromYamlFileEnvironment(filename string, environment string) newrelic.ConfigOption {
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
