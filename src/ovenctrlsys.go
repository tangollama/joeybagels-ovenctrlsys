package main

import (
	"os"
	"encoding/json"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/data/event"
	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
}

const (
	integrationName    = "com.joeybagels.ovenctrlsys"
	integrationVersion = "0.1.0"
)

var (
	args argumentList
)

func main() {
	var data Status
	byt := []byte(Simulate())
	if err := json.Unmarshal(byt, &data); err != nil {
		//fmt.Println("There was an error:", err)
		log.Error(err.Error())
		os.Exit(1)
    }
    //fmt.Println(dat)

	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	// Create Entity, entities name must be unique
	e1, err := i.Entity("OvenCtrlSys", "custom")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	// Add Event
	if data.Restart == true {
		err = e1.AddEvent(event.New("restart", "status"))
		if err != nil {
			log.Error(err.Error())
		}
	}

	// Add Inventory item
	if data.Version != "" {
		err = e1.SetInventoryItem("instance", "version", data.Version)
		if err != nil {
			log.Error(err.Error())
		}
	}

	// Add Metric
	if data.Temp > -1 {
		m1 := e1.NewMetricSet("OvenCtrlSysMetrics")
		err = m1.SetMetric("temp", data.Temp, metric.GAUGE)
		if err != nil {
			log.Error(err.Error())
		}
	}

	if err = i.Publish(); err != nil {
		log.Error(err.Error())
	}
}
