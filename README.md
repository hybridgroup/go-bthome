# go-bthome

Minimal package for creating and parsing [BTHome](https://bthome.io/) service data using Go/TinyGo.

## How to use

### Create a BTHome Payload

```go
// create the data payload
buf := &bthome.Payload{}

// add some data
value := bthome.DataValue{bthome.Acceleration, []byte{0x01, 0x02}}
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
	println(v.Type.Name, v.Value)
}
```

## Missing features

- Encryption
- Events
