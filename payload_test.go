package bthome

import (
	"testing"
)

func TestAddData(t *testing.T) {
	buf := &Payload{}
	value := NewDataValue(Acceleration)
	value.Set(float32(1.0))
	err := buf.AddData(value)
	if err != nil {
		t.Error(err)
	}
}

func TestGetData(t *testing.T) {
	buf := &Payload{}
	value := NewDataValue(Acceleration)
	val := float32(1.23)
	value.Set(val)
	err := buf.AddData(value)
	if err != nil {
		t.Error(err)
	}
	data, err := buf.GetData(Acceleration)
	if err != nil {
		t.Error(err)
	}
	if data.Get() != val {
		t.Error("data mismatch", data.Get(), val)
	}
}

func TestReset(t *testing.T) {
	buf := &Payload{}
	value := NewDataValue(Acceleration)
	value.Set(float32(1.0))
	err := buf.AddData(value)
	if err != nil {
		t.Error(err)
	}
	buf.Reset()
	if len(buf.Bytes()) != 1 {
		t.Error("reset failed")
	}
}

func TestServiceData(t *testing.T) {
	buf := &Payload{}
	value := NewDataValue(Acceleration)
	value.Set(float32(1.0))
	err := buf.AddData(value)
	if err != nil {
		t.Error(err)
	}
	data := buf.ServiceData()
	if data.UUID != ServiceUUID {
		t.Error("UUID mismatch")
	}
	if len(data.Data) != 4 {
		t.Error("data length mismatch", len(data.Data))
	}
	if data.Data[0] != DeviceInformation || data.Data[1] != 81 || data.Data[2] != 231 || data.Data[3] != 3 {
		t.Error("data mismatch", data.Data[0], data.Data[1], data.Data[2], data.Data[3])
	}
}

func TestAddDataInvalidSize(t *testing.T) {
	buf := &Payload{}
	val := Float32Value{DataValue{DataType: Acceleration, Value: []byte{0x01}}}
	err := buf.AddData(val)
	if err != errInvalidSize {
		t.Error("expected invalid size error")
	}
}

func TestAddDataBufferFull(t *testing.T) {
	buf := &Payload{}
	for i := 0; i < 10; i++ {
		value := NewDataValue(Acceleration)
		value.Set(float32(1.0))
		err := buf.AddData(value)
		if err != nil {
			t.Error(err)
		}
	}

	value := NewDataValue(Acceleration)
	value.Set(float32(1.0))
	err := buf.AddData(value)

	if err != errBufferFull {
		t.Error("expected buffer full error")
	}
}

func TestGetDataNoData(t *testing.T) {
	buf := &Payload{}
	_, err := buf.GetData(Acceleration)
	if err != errDataNotFound {
		t.Error("expected data not found error")
	}
}

func TestGetDataNotFound(t *testing.T) {
	buf := &Payload{}
	value := NewDataValue(Acceleration)
	value.Set(float32(1.0))
	err := buf.AddData(value)
	if err != nil {
		t.Error(err)
	}

	_, err = buf.GetData(Battery)
	if err != errDataNotFound {
		t.Error("expected data not found error")
	}
}

func TestParse(t *testing.T) {
	buf := &Payload{}
	value := NewDataValue(Acceleration)
	av := float32(1.23)
	value.Set(av)
	err := buf.AddData(value)

	if err != nil {
		t.Error(err)
	}
	val2 := NewDataValue(Humidity8)
	val2.Set(27)
	err = buf.AddData(val2)
	if err != nil {
		t.Error(err)
	}

	buf2 := NewPayload(buf.Bytes())
	data, err := buf2.Parse()
	if err != nil {
		t.Error(err)
	}
	if len(data) != 2 {
		t.Error("data length mismatch", len(data))
	}
	if data[0].Type().ID() != Acceleration.ID() || data[0].Get() != av {
		t.Error("Acceleration data mismatch", data[0].Get(), av)
	}
	if data[1].Type().ID() != Humidity8.ID() || data[1].Get() != 27 {
		t.Error("Humidity8 data mismatch", data[1].Data()[0])
	}
}
