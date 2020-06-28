package main

import (
	"fmt"
	"github.com/jeffalder/nr.yml/v3/nr_yml"
	"github.com/newrelic/go-agent/v3/newrelic"
	"path"
	"runtime"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	cfgFile := path.Join(path.Dir(filename), "config.yml")
	fmt.Println(cfgFile)

	cfg := new(newrelic.Config)

	nr_yml.ConfigFromYaml(cfgFile)(cfg)

	fmt.Println("app name:", cfg.AppName)
	fmt.Println("license key:", cfg.License)
	fmt.Println("dt enabled:", cfg.DistributedTracer.Enabled)
	fmt.Println("policy token:", cfg.SecurityPoliciesToken)
	fmt.Println("trace observer host:", cfg.InfiniteTracing.TraceObserver.Host)
	fmt.Println("total ram mib:", cfg.Utilization.TotalRAMMIB)

	nr_yml.ConfigFromYamlEnvironment(cfgFile, "development")(cfg)
	fmt.Println("app name (env = development):", cfg.AppName)
}
