package main

import (
	"os"
	"fmt"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/victorsteven/fullstack/api"
)

func main() {
    // Create an Application:
    nrapp, err := newrelic.NewApplication(
        // Name your application
        newrelic.ConfigAppName("FullStackApiServer"),
        // Fill in your New Relic license key
        newrelic.ConfigLicense("55e069c79013db1caa153292cbf6229e1de0NRAL"),
        // Add logging:
        newrelic.ConfigDebugLogger(os.Stdout),
        // Optional: add additional changes to your configuration via a config function:
        func(cfg *newrelic.Config) {
            cfg.CustomInsightsEvents.Enabled = false
        },
    )
    // If an application could not be created then err will reveal why.
    if err != nil {
        fmt.Println("unable to create New Relic Application", err)
    }

	api.Run(nrapp)

}
