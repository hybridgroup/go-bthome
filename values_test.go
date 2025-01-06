package bthome

import "testing"

func TestFloatDataValue(t *testing.T) {
	buf := &Payload{}

	value := NewDataValue(Acceleration)
	testvalue := float32(1.23)
	value.Set(testvalue)

	err := buf.AddData(value)
	if err != nil {
		t.Error(err)
	}

	data, err := buf.GetData(Acceleration)
	if err != nil {
		t.Error(err)
	}

	if data.Get() != testvalue {
		t.Error("data mismatch", data.Data()[0], data.Data()[1], data.Get())
	}
}

func TestInt8DataValue(t *testing.T) {
	buf := &Payload{}

	value := NewDataValue(Battery)
	value.Set(100)

	err := buf.AddData(value)
	if err != nil {
		t.Error(err)
	}

	data, err := buf.GetData(Battery)
	if err != nil {
		t.Error(err)
	}

	if data.Get() != 100 {
		t.Error("data mismatch", data.Data()[0])
	}
}

func TestInt16DataValue(t *testing.T) {
	buf := &Payload{}

	value := NewDataValue(Conductivity)
	value.Set(10000)

	err := buf.AddData(value)
	if err != nil {
		t.Error(err)
	}

	data, err := buf.GetData(Conductivity)
	if err != nil {
		t.Error(err)
	}

	if data.Get() != 10000 {
		t.Error("data mismatch", data.Data()[0])
	}
}

func TestBoolDataValue(t *testing.T) {
	buf := &Payload{}

	value := NewDataValue(Presence)
	value.Set(true)

	err := buf.AddData(value)
	if err != nil {
		t.Error(err)
	}

	data, err := buf.GetData(Presence)
	if err != nil {
		t.Error(err)
	}

	if data.Get() != true {
		t.Error("data mismatch", data.Data()[0])
	}
}
