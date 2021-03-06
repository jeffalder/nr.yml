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
package main

import (
	"fmt"
	"github.com/jeffalder/nr.yml/v3/nr_yml"
	"github.com/newrelic/go-agent/v3/newrelic"
	"path"
	"runtime"
)

// This example will prioritize the environment over
// the file named "config.yml" in the current directory.
// This is because newrelic.NewApplication sets values via last-wins.
//
//noinspection GoUnusedFunction,GoUnusedVariable
//func unusedUntestableExample() {
//	app, err := newrelic.NewApplication(
//		nr_yml.ConfigFromYamlFile("config.yml"),
//		newrelic.ConfigFromEnvironment())
//}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	cfgFile := path.Join(path.Dir(filename), "config.yml")
	fmt.Println(cfgFile)

	cfg := new(newrelic.Config)

	nr_yml.ConfigFromYamlFile(cfgFile)(cfg)
	fmt.Println("--- start raw print ---")
	fmt.Println(cfg)
	fmt.Println("--- end raw print ---")
	fmt.Println("app name:", cfg.AppName)
	fmt.Println("dt enabled:", cfg.DistributedTracer.Enabled)
	fmt.Println("policy token:", cfg.SecurityPoliciesToken)
	fmt.Println("trace observer host:", cfg.InfiniteTracing.TraceObserver.Host)
	fmt.Println("total ram mib:", cfg.Utilization.TotalRAMMIB)
	fmt.Println("labels:")
	for key, value := range cfg.Labels {
		fmt.Printf("  label \"%s\" has value \"%s\"\n", key, value)
	}

	nr_yml.ConfigFromYamlFileEnvironment(cfgFile, "development")(cfg)
	fmt.Println("app name (env = development):", cfg.AppName)
}
