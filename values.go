package bthome

import (
	"encoding/binary"
	"strconv"
)

// GetDataValue returns a new DataValuer for the given data type and data.
func GetDataValue(typ DataType, data []byte) DataValuer {
	switch typ.TypeID() {
	case TypeFloat32:
		return Float32Value{DataValue{DataType: typ, Value: data}}
	case TypeInt8:
		return Int8Value{DataValue{DataType: typ, Value: data}}
	case TypeInt16:
		return Int16Value{DataValue{DataType: typ, Value: data}}
	case TypeBool:
		return BoolValue{DataValue{DataType: typ, Value: data}}
	}

	return nil
}

// NewDataValue returns a new DataValuer for the given data type.
func NewDataValue(typ DataType) DataValuer {
	switch typ.TypeID() {
	case TypeFloat32:
		return Float32Value{DataValue{DataType: typ, Value: make([]byte, typ.Size())}}
	case TypeInt8:
		return Int8Value{DataValue{DataType: typ, Value: make([]byte, typ.Size())}}
	case TypeInt16:
		return Int16Value{DataValue{DataType: typ, Value: make([]byte, typ.Size())}}
	case TypeBool:
		return BoolValue{DataValue{DataType: typ, Value: make([]byte, typ.Size())}}
	}

	return nil
}

// DataValuer is an interface for data values in the service data payload.
type DataValuer interface {
	Type() DataType
	Data() []byte
	Get() any
	Set(any)
	String() string
}

// DataValue represents an individual data value in the service data payload.
type DataValue struct {
	DataType DataType
	Value    []byte
}

// Type returns the data type of the data value.
func (d DataValue) Type() DataType {
	return d.DataType
}

// Data returns the raw data of the data value.
func (d DataValue) Data() []byte {
	return d.Value
}

// Float32Value represents a float32 data value in the service data payload.
type Float32Value struct {
	DataValue
}

// Get returns the float32 value of the data value.
func (d Float32Value) Get() any {
	return float32(float32(binary.LittleEndian.Uint16(d.Value)) * d.DataType.Factor())
}

// Set sets the float32 value of the data value.
func (d Float32Value) Set(v any) {
	var val float32
	switch v := v.(type) {
	case float32:
		val = v / float32(d.DataType.Factor())
	case float64:
		val = float32(v) / float32(d.DataType.Factor())
	}

	binary.LittleEndian.PutUint16(d.Value, uint16(val))
}

func (d Float32Value) String() string {
	return d.Type().Name() + ": " + strconv.FormatFloat(float64(d.Get().(float32)), 'f', -1, 32)
}

type Int8Value struct {
	DataValue
}

func (d Int8Value) Get() any {
	return int(d.Value[0]) / int(d.DataType.Factor())
}

func (d Int8Value) Set(v any) {
	val := v.(int) * int(d.DataType.Factor())
	d.Value[0] = byte(val)
}

func (d Int8Value) String() string {
	return d.Type().Name() + ": " + strconv.Itoa(d.Get().(int))
}

type Int16Value struct {
	DataValue
}

func (d Int16Value) Get() any {
	val := binary.LittleEndian.Uint16(d.Value)
	return int(float32(val) / d.DataType.Factor())
}

func (d Int16Value) Set(v any) {
	val := v.(int) * int(d.DataType.Factor())
	binary.LittleEndian.PutUint16(d.Value, uint16(val))
}

func (d Int16Value) String() string {
	return d.Type().Name() + ": " + strconv.Itoa(d.Get().(int))
}

type BoolValue struct {
	DataValue
}

func (d BoolValue) Get() any {
	return d.Value[0] != 0
}

func (d BoolValue) Set(v any) {
	val := v.(bool)
	if val {
		d.Value[0] = 1
	} else {
		d.Value[0] = 0
	}
}

func (d BoolValue) String() string {
	return d.Type().Name() + ": " + strconv.FormatBool(d.Get().(bool))
}
