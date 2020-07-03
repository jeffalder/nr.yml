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

import "github.com/newrelic/go-agent/v3/newrelic"

// Don't use this; it's only exported for the yaml parser
type AttributesYaml struct {
	Enabled *bool     `yaml:"enabled"`
	Include *[]string `yaml:"include"`
	Exclude *[]string `yaml:"exclude"`
}

func (attributesYaml AttributesYaml) update(cfg *newrelic.Config) {
	if attributesYaml.Enabled != nil {
		cfg.Attributes.Enabled = *attributesYaml.Enabled
	}

	if attributesYaml.Include != nil {
		cfg.Attributes.Include = *attributesYaml.Include
	}

	if attributesYaml.Exclude != nil {
		cfg.Attributes.Exclude = *attributesYaml.Exclude
	}
}
