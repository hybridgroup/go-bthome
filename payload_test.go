package bthome

import (
	"testing"
)

func TestAddData(t *testing.T) {
	buf := &Payload{}
	err := buf.AddData(Acceleration, []byte{0x01, 0x02})
	if err != nil {
		t.Error(err)
	}
}

func TestGetData(t *testing.T) {
	buf := &Payload{}
	err := buf.AddData(Acceleration, []byte{0x01, 0x02})
	if err != nil {
		t.Error(err)
	}
	data, err := buf.GetData(Acceleration)
	if err != nil {
		t.Error(err)
	}
	if data[0] != 0x01 || data[1] != 0x02 {
		t.Error("data mismatch")
	}
}

func TestReset(t *testing.T) {
	buf := &Payload{}
	err := buf.AddData(Acceleration, []byte{0x01, 0x02})
	if err != nil {
		t.Error(err)
	}
	buf.Reset()
	if len(buf.Bytes()) != 0 {
		t.Error("reset failed")
	}
}

func TestServiceData(t *testing.T) {
	buf := &Payload{}
	err := buf.AddData(Acceleration, []byte{0x01, 0x02})
	if err != nil {
		t.Error(err)
	}
	data := buf.ServiceData()
	if data.UUID != ServiceUUID {
		t.Error("UUID mismatch")
	}
	if len(data.Data) != 3 {
		t.Error("data length mismatch")
	}
	if data.Data[0] != 0x51 || data.Data[1] != 0x01 || data.Data[2] != 0x02 {
		t.Error("data mismatch")
	}
}
