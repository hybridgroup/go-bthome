package main

import (
	"math/rand"
	"time"

	"github.com/hybridgroup/go-bthome"
	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

var (
	opts = bluetooth.AdvertisementOptions{
		LocalName:         "Go BTHome",
		AdvertisementType: bluetooth.AdvertisingTypeScanInd,
	}

	payload         = bthome.Payload{}
	temperatureData = bthome.NewDataValue(bthome.Temperature)
)

func main() {
	must("enable BLE stack", adapter.Enable())
	adv := adapter.DefaultAdvertisement()

	println("advertising...")
	address, _ := adapter.Address()
	for {
		payload.Reset()

		temperatureData.Set(float32(randomInt(150, 250)) / float32(10))

		payload.AddData(temperatureData)

		opts.ServiceData = []bluetooth.ServiceDataElement{payload.ServiceData()}

		must("config adv", adv.Configure(opts))
		must("start adv", adv.Start())

		println("Go BTHome /", address.MAC.String())
		time.Sleep(time.Second)
		adv.Stop()
	}
}

// Returns an int >= min, < max
func randomInt(min, max int) uint16 {
	return uint16(min + rand.Intn(max-min))
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
