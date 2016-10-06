package main

import (
	"fmt"
	"os"
	"time"

	"github.com/newrelic/go-agent"
)

const (
	licenseVar = "NEW_RELIC_LICENSE_KEY"
	appname    = "Short Lived Process Application"
)

func main() {
	lic := os.Getenv(licenseVar)
	if "" == lic {
		fmt.Printf("environment variable %s unset\n", licenseVar)
		os.Exit(1)
	}

	cfg := newrelic.NewConfig(appname, lic)
	cfg.Logger = newrelic.NewDebugLogger(os.Stdout)
	app, err := newrelic.NewApplication(cfg)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	// Wait for the application to connect.
	if err := app.WaitForConnection(5 * time.Second); nil != err {
		fmt.Println(err)
	}

	// Do the tasks at hand.  Perhaps record them using transactions and/or
	// custom events.
	tasks := []string{"white", "black", "red", "blue", "green", "yellow"}
	for _, task := range tasks {
		txn := app.StartTransaction("task", nil, nil)
		time.Sleep(10 * time.Millisecond)
		txn.End()
		app.RecordCustomEvent("task", map[string]interface{}{
			"color": task,
		})
	}

	// Shut down the application to flush data to New Relic.
	app.Shutdown()
}