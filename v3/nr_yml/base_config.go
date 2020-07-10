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
	"os"
)

// Don't use this; it's only exported for the yaml parser
type ConfigYaml struct {
	AppName               *string                 `yaml:"app_name"`
	License               *string                 `yaml:"license_key"`
	Host                  *string                 `yaml:"host"`
	Enabled               *bool                   `yaml:"agent_enabled"`
	HighSecurity          *bool                   `yaml:"high_security"`
	SecurityPoliciesToken *string                 `yaml:"security_policies_token"`
	LogStreamName         *string                 `yaml:"log_stream_name"`
	LogLevel              *string                 `yaml:"log_level"`
	Labels                *map[string]string      `yaml:"labels"`
	DistributedTracing    *DistributedTracingYaml `yaml:"distributed_tracing"`
	InfiniteTracing       *InfiniteTracingYaml    `yaml:"infinite_tracing"`
	ProcessHost           *ProcessHostYaml        `yaml:"process_host"`
	Attributes            *AttributesYaml         `yaml:"attributes"`
	Utilization           *UtilizationYaml        `yaml:"utilization"`
}

func (yamlValues ConfigYaml) update(cfg *newrelic.Config) {
	if yamlValues.AppName != nil {
		newrelic.ConfigAppName(*yamlValues.AppName)(cfg)
	}

	if yamlValues.Enabled != nil {
		newrelic.ConfigEnabled(*yamlValues.Enabled)(cfg)
	}

	if yamlValues.License != nil {
		newrelic.ConfigLicense(*yamlValues.License)(cfg)
	}

	if yamlValues.HighSecurity != nil {
		cfg.HighSecurity = *yamlValues.HighSecurity
	}

	if yamlValues.SecurityPoliciesToken != nil {
		cfg.SecurityPoliciesToken = *yamlValues.SecurityPoliciesToken
	}

	if yamlValues.Host != nil {
		cfg.Host = *yamlValues.Host
	}

	if yamlValues.Labels != nil {
		cfg.Labels = *yamlValues.Labels
	}

	if yamlValues.Attributes != nil {
		yamlValues.Attributes.update(cfg)
	}

	if yamlValues.InfiniteTracing != nil {
		yamlValues.InfiniteTracing.update(cfg)
	}

	if yamlValues.DistributedTracing != nil {
		yamlValues.DistributedTracing.update(cfg)
	}

	if yamlValues.ProcessHost != nil {
		yamlValues.ProcessHost.update(cfg)
	}

	if yamlValues.Utilization != nil {
		yamlValues.Utilization.update(cfg)
	}

	yamlValues.updateLogging(cfg)
}

func (yamlValues ConfigYaml) updateLogging(cfg *newrelic.Config) {
	var logStream *os.File

	if yamlValues.LogStreamName == nil {
		return
	}

	switch *yamlValues.LogStreamName {
	case "STDOUT", "Stdout", "stdout":
		logStream = os.Stdout
	case "STDERR", "Stderr", "stderr":
		logStream = os.Stderr
	}

	if logStream != nil {
		if yamlValues.LogLevel == nil {
			newrelic.ConfigInfoLogger(logStream)(cfg)
		} else {
			switch *yamlValues.LogLevel {
			case "DEBUG", "Debug", "debug":
				newrelic.ConfigDebugLogger(logStream)(cfg)
			default:
				newrelic.ConfigInfoLogger(logStream)(cfg)
			}
		}
	}
}
