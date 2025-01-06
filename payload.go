package bthome

import "tinygo.org/x/bluetooth"

// Payload encapsulates a raw BTHome data packet.
type Payload struct {
	data [31]byte
	len  uint8
}

// NewPayload creates a new service data payload with the given data.
func NewPayload(data []byte) *Payload {
	buf := &Payload{}
	buf.Reset()
	copy(buf.data[:], data)
	buf.len = uint8(len(data)) - 1

	return buf
}

// Bytes returns the raw data packet as a byte slice.
func (buf *Payload) Bytes() []byte {
	return buf.data[:buf.len+1]
}

// AddData adds data ([]byte) entries to the service data payload.
func (buf *Payload) AddData(value DataValuer) error {
	// Check if length of value is correct
	if len(value.Data()) != value.Type().Size() {
		return errInvalidSize
	}

	// make sure we have device information
	if buf.data[0] != DeviceInformation {
		buf.Reset()
	}

	// Check whether the field can fit this data.
	fieldLength := len(value.Data()) + 1
	if int(buf.len)+fieldLength > len(buf.data) {
		return errBufferFull
	}

	// Add the data.
	buf.data[buf.len+1] = value.Type().ID()
	copy(buf.data[buf.len+2:], value.Data())
	buf.len += uint8(fieldLength)

	return nil
}

// GetData retrieves data from the service data payload.
func (buf *Payload) GetData(typ DataType) (DataValuer, error) {
	values, err := buf.Parse()
	if err != nil {
		return nil, err
	}

	for _, value := range values {
		if value.Type().ID() == typ.ID() {
			return value, nil
		}
	}

	return nil, errDataNotFound
}

// Reset clears the service data payload.
func (buf *Payload) Reset() {
	buf.data[0] = DeviceInformation
	for i := 1; i < len(buf.data); i++ {
		buf.data[i] = 0
	}
	buf.len = 0
}

// ServiceData returns the service data payload as a bluetooth.ServiceDataElement.
func (buf *Payload) ServiceData() bluetooth.ServiceDataElement {
	return bluetooth.ServiceDataElement{
		UUID: ServiceUUID,
		Data: buf.Bytes(),
	}
}

// Parse extracts the data values from the service data payload.
func (buf *Payload) Parse() ([]DataValuer, error) {
	values := []DataValuer{}
	data := buf.Bytes()
	for i := 1; i < len(data); i++ {
		dt := FindDataType(data[i])
		if dt.ID() == 0 {
			// unknown data type
			return nil, errDataNotFound
		}

		values = append(values, GetDataValue(dt, data[i+1:i+1+dt.Size()]))
		i += dt.Size()
	}

	return values, nil
}
