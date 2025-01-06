# go-bthome

Minimal package for creating and parsing [BTHome](https://bthome.io/) service data using Go/TinyGo.

## How to use

### Create a BTHome Payload

```go
// create the data payload
buf := &bthome.Payload{}

// add some data
value := bthome.NewDataValue(bthome.Acceleration)
value.Set(float32(1.23))
err := buf.AddData(value)
if err != nil {
    t.Error(err)
}

// now publish the bluetooth.ServiceDataElement
data := buf.ServiceData()
```

### Parse a BTHome Payload

```go
data := []byte{...}
buf := bthome.NewPayload(data)
values, _ := buf.Parse()
for _, v := range values {
	println(v.Type().Name(), v.Get())
}
```

## Examples

### Thermometer

[examples/thermometer](./examples/thermometer)

Thermometer beacon device written using TinyGo

```shell
$ tinygo flash -target nano-rp2040 -monitor ./examples/thermometer/
Connected to /dev/ttyACM0. Press Ctrl-C to exit.
advertising...
Go BTHome / D2:00:10:09:05:0A
Go BTHome / D2:00:10:09:05:0A
...
```

## CLI Tools

### bthomescan

[cmd/bthomescan](./cmd/bthomescan)

Scans for any devices that are advertising BTHome packets, then displays the data.

```shell
go run ./cmd/bthomescan
```

```shell
$ go run ./cmd/bthomescan/
scanning...
...
found device: 58:BF:25:3A:EB:D2 -50 Go BTHome
BTHome device found
temperature: 18.5
...
found device: 58:BF:25:3A:EB:D2 -50 Go BTHome
BTHome device found
temperature: 15.400001
...
found device: 58:BF:25:3A:EB:D2 -51 Go BTHome
BTHome device found
temperature: 24.6
```


## Missing features

- Encryption
- Events
