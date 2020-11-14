package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/system"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	GetCloudProvider bool
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "sensu-platform-discovery",
			Short:    "Discover system platform information and output a list of agent subscriptions.",
			Keyspace: "sensu.io/plugins/sensu-platform-discovery/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		&sensu.PluginConfigOption{
			Path:      "get-cloud-provider",
			Env:       "GET_CLOUD_PROVIDER",
			Argument:  "get-cloud-provider",
			Shorthand: "c",
			Usage:     "Determine the cloud provider and include it in subscriptions.",
			Value:     &plugin.GetCloudProvider,
		},
	}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
	if len(plugin.Example) == 0 {
		return sensu.CheckStateWarning, fmt.Errorf("--example or CHECK_EXAMPLE environment variable is required")
	}
	return sensu.CheckStateOK, nil
}

func platformSubs() (string, error) {
	subs := ""

	infoCtx, cancel := context.WithTimeout(ctx, time.Duration(10)*time.Second)

	info, err := system.Info()
	if err != nil {
		return nil, err
	}

	if plugin.GetCloudProvider {
		subs = subs + system.GetCloudProvider(infoCtx) + "\n"
	}

	subs = subs + fmt.Sprintf("%s\n%s\n%s\n", info.OS, info.Platform, info.PlatformFamily)

	return subs, nil
}

func executeCheck(event *types.Event) (int, error) {
	log.Println("executing check with --example", plugin.Example)
	return sensu.CheckStateOK, nil
}
