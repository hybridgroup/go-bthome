package bthome

import (
	"testing"
)

func TestAddData(t *testing.T) {
	buf := &Payload{}
	err := buf.AddData(DataValue{Acceleration, []byte{0x01, 0x02}})
	if err != nil {
		t.Error(err)
	}
}

func TestGetData(t *testing.T) {
	buf := &Payload{}
	err := buf.AddData(DataValue{Acceleration, []byte{0x01, 0x02}})
	if err != nil {
		t.Error(err)
	}
	data, err := buf.GetData(Acceleration)
	if err != nil {
		t.Error(err)
	}
	if data.Value[0] != 0x01 || data.Value[1] != 0x02 {
		t.Error("data mismatch")
	}
}

func TestReset(t *testing.T) {
	buf := &Payload{}
	err := buf.AddData(DataValue{Acceleration, []byte{0x01, 0x02}})
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
	err := buf.AddData(DataValue{Acceleration, []byte{0x01, 0x02}})
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
	if data.Data[0] != DeviceInformation || data.Data[1] != 0x51 || data.Data[2] != 0x01 || data.Data[3] != 0x02 {
		t.Error("data mismatch")
	}
}

func TestAddDataInvalidSize(t *testing.T) {
	buf := &Payload{}
	err := buf.AddData(DataValue{Acceleration, []byte{0x01}})
	if err != errInvalidSize {
		t.Error("expected invalid size error")
	}
}

func TestAddDataBufferFull(t *testing.T) {
	buf := &Payload{}
	for i := 0; i < 10; i++ {
		err := buf.AddData(DataValue{Acceleration, []byte{0x01, 0x02}})
		if err != nil {
			t.Error(err)
		}
	}

	err := buf.AddData(DataValue{Acceleration, []byte{0x01, 0x02}})
	if err != errBufferFull {
		t.Error("expected buffer full error")
	}
}

func TestGetDataNotFound(t *testing.T) {
	buf := &Payload{}
	_, err := buf.GetData(Acceleration)
	if err != errDataNotFound {
		t.Error("expected data not found error")
	}
}

func TestGetDataOutOfBounds(t *testing.T) {
	buf := &Payload{}
	err := buf.AddData(DataValue{Acceleration, []byte{0x01, 0x02}})
	if err != nil {
		t.Error(err)
	}
	_, err = buf.GetData(DataType{"invalid", 0x00, 2})
	if err != errDataNotFound {
		t.Error("expected data not found error")
	}
}

func TestParse(t *testing.T) {
	buf := &Payload{}
	err := buf.AddData(DataValue{Acceleration, []byte{0x01, 0x02}})
	if err != nil {
		t.Error(err)
	}
	err = buf.AddData(DataValue{Humidity8, []byte{0x03}})
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
	if data[0].Type.ID != Acceleration.ID || data[0].Value[0] != 0x01 || data[0].Value[1] != 0x02 {
		t.Error("data mismatch")
	}
	if data[1].Type.ID != Humidity8.ID || data[1].Value[0] != 0x03 {
		t.Error("data mismatch")
	}
}
