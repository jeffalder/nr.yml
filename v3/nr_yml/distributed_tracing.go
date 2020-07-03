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
type DistributedTracingYaml struct {
	Enabled *bool `yaml:"enabled"`
}

func (distributedTracingYaml DistributedTracingYaml) update(cfg *newrelic.Config) {
	if distributedTracingYaml.Enabled != nil {
		newrelic.ConfigDistributedTracerEnabled(*distributedTracingYaml.Enabled)(cfg)
	}
}
