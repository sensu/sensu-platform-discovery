package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	corev2 "github.com/sensu/sensu-go/api/core/v2"
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

func checkArgs(event *corev2.Event) (int, error) {
	return sensu.CheckStateOK, nil
}

func platformSubs() ([]string, error) {
	subs := []string{}

	infoCtx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	info, err := system.Info()
	if err != nil {
		return subs, err
	}

	if plugin.GetCloudProvider {
		cloud := system.GetCloudProvider(infoCtx)
		if cloud != "" {
			subs = append(subs, cloud)
		}
	}

	subs = append(subs, []string{info.OS, info.Platform}...)

	if info.PlatformFamily != info.Platform {
		subs = append(subs, info.PlatformFamily)
	}

	return subs, nil
}

func executeCheck(event *corev2.Event) (int, error) {
	subs, err := platformSubs()

	fmt.Println(strings.Join(subs, "\n"))

	if err != nil {
		return sensu.CheckStateWarning, err
	}

	return sensu.CheckStateOK, err
}
