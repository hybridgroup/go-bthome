# go-bthome

Minimal package for creating and parsing BTHome service data using Go/TinyGo.

## How to use

```go
// create the data payload
buf := &bthome.Payload{}

// add some data
err := buf.AddData(bthome.Acceleration, []byte{0x01, 0x02})
if err != nil {
    t.Error(err)
}

// now publish the bluetooth.ServiceDataElement
data := buf.ServiceData()
```
