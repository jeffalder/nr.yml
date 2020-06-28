package nr_yml

import "github.com/newrelic/go-agent/v3/newrelic"

// Don't use this; it's only exported for the yaml parser
type UtilizationYaml struct {
	BillingHostname   *string `yaml:"billing_hostname"`
	LogicalProcessors *uint16 `yaml:"logical_processors"`
	TotalRAMMIB       *uint32 `yaml:"total_ram_mib"`
	DetectAWS         *bool   `yaml:"detect_aws"`
	DetectDocker      *bool   `yaml:"detect_docker"`
	DetectPCF         *bool   `yaml:"detect_pcf"`
	DetectGCP         *bool   `yaml:"detect_gcp"`
	DetectAzure       *bool   `yaml:"detect_azure"`
	DetectKubernetes  *bool   `yaml:"detect_kubernetes"`
}

func (utilizationYaml UtilizationYaml) update(cfg *newrelic.Config) {
	if utilizationYaml.BillingHostname != nil {
		cfg.Utilization.BillingHostname = *utilizationYaml.BillingHostname
	}
	if utilizationYaml.LogicalProcessors != nil {
		cfg.Utilization.LogicalProcessors = int(*utilizationYaml.LogicalProcessors)
	}
	if utilizationYaml.TotalRAMMIB != nil {
		cfg.Utilization.TotalRAMMIB = int(*utilizationYaml.TotalRAMMIB)
	}
	if utilizationYaml.DetectAWS != nil {
		cfg.Utilization.DetectAWS = *utilizationYaml.DetectAWS
	}
	if utilizationYaml.DetectAzure != nil {
		cfg.Utilization.DetectAzure = *utilizationYaml.DetectAzure
	}
	if utilizationYaml.DetectDocker != nil {
		cfg.Utilization.DetectDocker = *utilizationYaml.DetectDocker
	}
	if utilizationYaml.DetectGCP != nil {
		cfg.Utilization.DetectGCP = *utilizationYaml.DetectGCP
	}
	if utilizationYaml.DetectPCF != nil {
		cfg.Utilization.DetectPCF = *utilizationYaml.DetectPCF
	}
	if utilizationYaml.DetectKubernetes != nil {
		cfg.Utilization.DetectKubernetes = *utilizationYaml.DetectKubernetes
	}
}
