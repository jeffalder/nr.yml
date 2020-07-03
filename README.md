# nr.yml

Yaml configuration for New Relic's Go SDK

**Important Support Note**: This is _not_ a New Relic product. _Do not_ contact New Relic support for issues with this code. Instead, _do_ file GitHub issues in this project.

## Usage

Import this project alongside the New Relic Go agent:

```go
import (
	"github.com/jeffalder/nr.yml/v3/nr_yml"
	"github.com/newrelic/go-agent/v3/newrelic"
)
```

Add this as a `ConfigOption`:

```go
	app, err := newrelic.NewApplication(
		nr_yml.ConfigFromYamlFile("config.yml"),
		newrelic.ConfigFromEnvironment())
```

The New Relic Go SDK will apply configuration options in order. So in the example above, The YAML settings will apply first, then environment settings will apply over the YAML settings.

For more choices, see [config_from_yaml.go](https://github.com/jeffalder/nr.yml/blob/main/v3/nr_yml/config_from_yaml.go)

*Note*: Some types are Exposed for Yaml parsing and databinding. They are not intended to be used.

## Environment

The Yaml file is comprised of a single mapping. The keys of this mapping are "environments", such as `development`, `qa`, or `production`. The environment must be a string. It is chosen in the code from one of three ways:
1. As the second parameter of `ConfigFromYamlFileEnvironment`
2. As the environment variable `NEW_RELIC_ENVIRONMENT` (which is used by `ConfigFromDefaultYaml` or `ConfigFromYamlFile`)
3. If neither version above is specified, the code uses `production`.

The values of the mapping are configuration hierarchies.

Most commonly, you will give a `common` mapping an anchor, and use that anchor in the actual environments, like this:

```yaml
common: &default_settings
  app_name: my application
  agent_enabled: true
  license_key: abcabc123123abcabc
  
production:
  <<: *default_settings
  app_name: my application (prod)
```

This is how the template Yaml files provided for the New Relic Java agent and New Relic Ruby agent are configured.

## Supported Yaml options

The Yaml format is based on the [Java agent configuration file](https://docs.newrelic.com/docs/agents/java-agent/configuration/java-agent-configuration-config-file). The supported options, and their defaults, are given below.

```yaml
common: &default_settings
  # The application name, from `newrelic.ConfigAppName`
  app_name: <No Default>
  
  # True to enable the agent, false to prevent instrumentation.
  agent_enabled: true
  
  # Your New Relic APM License Key, as in `newrelic.ConfigLicense`
  license_key: <No Default>

  # High Security if your New Relic APM Account is configured for high security.
  # Leave default if you're not sure.
  high_security: false
  
  # Security Policies Token if your New Relic APM Account is configured with a Language Agent Security Policy.
  # Leave unset if you're not sure.
  security_policies_token: ""
  
  # These settings are not independent. You cannot set one here and one in the environment.
  # They control the location and level of agent logging, if you don't configure agent logging via `ConfigLogger`.
  # Stream: Your choice of STDOUT or STDERR.
  log_stream_name: <No Default>
  # If Stream is set to a valid value, then this can be set to DEBUG for Debug logging. Any other value configures Info-level logging.
  log_level: info

  # New Relic APM Labels. There are no labels by default. Labels must be set as a mapping, like this:
  # labels: 
  #   - label1: value1
  #   - label2: value2

  # Configure New Relic Distributed Tracing, as in newrelic.ConfigDistributedTracerEnabled.
  # Note that yaml uses `tracing` for consistency with other Yaml files and the environment variable NEW_RELIC_DISTRIBUTED_TRACING_ENABLED.  
  distributed_tracing:
    enabled: false

  # Configure Infinite Tracing on New Relic Edge. Here is an example:
  # infinite_tracing:
  #   trace_observer:
  #     host: my-trace-observer.example.com

  # Sets the values for attributes to be included or excluded from being sent to New Relic.
  # These lists replace what might have been set in the configuration already.
  attributes:
    enabled: true
    # Include and exclude must be passed as Yaml sequences.
    # include: 
    #   - attrib1
    # exclude:
    #   - attrib2

  # Sets the display name in New Relic APM, as in `HostDisplayName`.
  process_host:
    display_name: <No Default>

production:
  <<: *default_settings
```

## Contributions

Contributions are welcome. Submit an issue or (preferably) a PR!

## Code of Conduct

This repo follows the [Contributor Covenant Code of Conduct 2.0](https://www.contributor-covenant.org/version/2/0/code_of_conduct/). To report violations, reach out to me on Gitter.
