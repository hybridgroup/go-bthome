package main

import (
	"github.com/hybridgroup/go-bthome"
	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	// Enable BLE interface.
	must("enable BLE stack", adapter.Enable())

	// Start scanning.
	println("scanning...")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		println("found device:", device.Address.String(), device.RSSI, device.LocalName())
		if device.ServiceData() != nil {
			if device.ServiceData()[0].UUID == bthome.ServiceUUID {
				println("BTHome device found")
				payload := bthome.NewPayload(device.ServiceData()[0].Data)
				values, err := payload.Parse()
				if err != nil {
					println("failed to parse payload:", err.Error())
					return
				}
				for _, value := range values {
					println(value.String())
				}
			}
		}
	})
	must("start scan", err)
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
