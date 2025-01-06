package bthome

import (
	"tinygo.org/x/bluetooth"
)

var ServiceUUID = bluetooth.New16BitUUID(0xFCD2)

// DeviceInformation is the BTHome device info byte, which has several bits indicating
// the capabilities of the device.
const DeviceInformation = 0x40

const (
	TypeInt8 = iota
	TypeInt16
	TypeInt32
	TypeUint8
	TypeUint16
	TypeUint32
	TypeFloat32
	TypeString
	TypeBool
)

type DataType interface {
	Name() string
	ID() byte
	Size() int
	Unit() string
	Factor() float32
	TypeID() int
}

// dataType represents a data type that can be added to the service data payload.
type dataType struct {
	// Name of the data type
	name string

	// ID of the data type
	id byte

	// Size of the data type in bytes
	size int

	// Unit of the data type
	unit string

	// Factor to convert the data type to bytes unit
	factor float32

	// TypeID of the data type
	typeID int
}

func (t dataType) Name() string {
	return t.name
}

func (t dataType) ID() byte {
	return t.id
}

func (t dataType) Size() int {
	return t.size
}

func (t dataType) Unit() string {
	return t.unit
}

func (t dataType) Factor() float32 {
	return t.factor
}

func (t dataType) TypeID() int {
	return t.typeID
}

// See full list of field types:
// https://bthome.io/format/

// Sensor data types
var (
	Acceleration        = dataType{"acceleration", 0x51, 2, "m/s²", 0.001, TypeFloat32}
	Battery             = dataType{"battery", 0x01, 1, "%", 1, TypeInt8}
	CO2                 = dataType{"co2", 0x12, 2, "ppm", 1, TypeInt16}
	Conductivity        = dataType{"conductivity", 0x56, 2, "µS/cm", 1, TypeInt16}
	Count8              = dataType{"count", 0x09, 1, "", 1, TypeUint8}
	Count16             = dataType{"count", 0x3D, 2, "", 1, TypeUint16}
	Count32             = dataType{"count", 0x3E, 4, "", 1, TypeUint32}
	CountSint8          = dataType{"count", 0x59, 1, "", 1, TypeInt8}
	CountSint16         = dataType{"count", 0x5A, 2, "", 1, TypeInt16}
	CountSint32         = dataType{"count", 0x5B, 4, "", 1, TypeUint32}
	Current             = dataType{"current", 0x43, 2, "A", 0.001, TypeFloat32}
	CurrentSint16       = dataType{"current", 0x5D, 2, "A", 0.001, TypeFloat32}
	Dewpoint            = dataType{"dewpoint", 0x08, 2, "°C", 0.01, TypeFloat32}
	DistanceMM          = dataType{"distance (mm)", 0x40, 2, "mm", 1, TypeUint16}
	DistanceM           = dataType{"distance (m)", 0x41, 2, "m", 0.01, TypeFloat32}
	Duration            = dataType{"duration", 0x42, 3, "s", 1, TypeFloat32}
	Energy              = dataType{"energy", 0x4D, 4, "Wh", 0.001, TypeFloat32}
	Energy24            = dataType{"energy", 0x0A, 3, "Wh", 0.001, TypeFloat32}
	Gas24               = dataType{"gas", 0x4B, 3, "m2", 1, TypeFloat32}
	Gas32               = dataType{"gas", 0x4C, 4, "m2", 1, TypeFloat32}
	Gyroscope           = dataType{"gyroscope", 0x52, 2, "°/s", 0.01, TypeFloat32}
	Humidity16          = dataType{"humidity", 0x03, 2, "%", 1, TypeFloat32}
	Humidity8           = dataType{"humidity", 0x2E, 1, "%", 1, TypeInt8}
	Illuminance         = dataType{"illuminance", 0x05, 3, "lux", 1, TypeFloat32}
	MassKG              = dataType{"mass (kg)", 0x06, 2, "kg", 0.001, TypeFloat32}
	MassLB              = dataType{"mass (lb)", 0x07, 2, "lb", 0.001, TypeFloat32}
	Moisture16          = dataType{"moisture", 0x14, 2, "%", 1, TypeFloat32}
	Moisture8           = dataType{"moisture", 0x2F, 1, "%", 1, TypeInt8}
	PM25                = dataType{"pm2.5", 0x0D, 2, "µg/m³", 1, TypeInt16}
	PM10                = dataType{"pm10", 0x0E, 2, "µg/m³", 1, TypeInt16}
	Power               = dataType{"power", 0x0B, 3, "W", 0.001, TypeFloat32}
	PowerSint32         = dataType{"power", 0x5C, 4, "W", 0.001, TypeFloat32}
	Pressure            = dataType{"pressure", 0x04, 3, "hPa", 0.01, TypeFloat32}
	Raw                 = dataType{"raw", 0x54, -1, "", 1, TypeString}
	Rotation            = dataType{"rotation", 0x3F, 2, "°", 0.01, TypeFloat32}
	Speed               = dataType{"speed", 0x44, 2, "m/s", 0.01, TypeFloat32}
	TemperatureSint8    = dataType{"temperature", 0x57, 1, "°C", 0.01, TypeInt8}
	TemperatureSint8_35 = dataType{"temperature", 0x58, 1, "°C", 0.01, TypeInt8}
	Temperature16       = dataType{"temperature", 0x45, 2, "°C", 0.01, TypeInt16}
	Temperature16_01    = dataType{"temperature", 0x02, 2, "°C", 0.01, TypeInt16}
	Text                = dataType{"text", 0x53, -1, "", 1, TypeString}
	Timestamp           = dataType{"timestamp", 0x50, 4, "s", 1, TypeFloat32}
	TVOC                = dataType{"tvoc", 0x13, 2, "ppb", 1, TypeInt16}
	Voltage             = dataType{"voltage", 0x0C, 2, "V", 0.001, TypeFloat32}
	Voltage10           = dataType{"voltage", 0x4A, 2, "V", 0.001, TypeFloat32}
	Volume              = dataType{"volume", 0x4E, 4, "m³", 0.001, TypeFloat32}
	Volume16            = dataType{"volume", 0x47, 2, "m³", 0.001, TypeFloat32}
	VolumeML            = dataType{"volume", 0x48, 2, "mL", 1, TypeUint16}
	VolumeStorage       = dataType{"volume storage", 0x55, 4, "m³", 0.001, TypeFloat32}
	VolumeFlowRate      = dataType{"volume flow rate", 0x49, 2, "m³/s", 0.001, TypeFloat32}
	UVIndex             = dataType{"UV index", 0x46, 1, "", 1, TypeFloat32}
	Water               = dataType{"water", 0x4F, 4, "L", 0.001, TypeFloat32}
)

// Binary Sensor data should always be an uint8 of a single byte. Its value should be 1 for on, and 0 for off.
var (
	BatteryLow      = dataType{"battery", 0x15, 1, "", 1, TypeBool}
	BatteryCharging = dataType{"battery charging", 0x16, 1, "", 1, TypeBool}
	CarbonMonoxide  = dataType{"carbon monoxide", 0x17, 1, "", 1, TypeBool}
	Cold            = dataType{"cold", 0x18, 1, "", 1, TypeBool}
	Connectivity    = dataType{"connectivity", 0x19, 1, "", 1, TypeBool}
	Door            = dataType{"door", 0x1A, 1, "", 1, TypeBool}
	GarageDoor      = dataType{"garage door", 0x1B, 1, "", 1, TypeBool}
	Gas             = dataType{"gas", 0x1C, 1, "", 1, TypeBool}
	GenericBoolean  = dataType{"generic boolean", 0x0F, 1, "", 1, TypeBool}
	Heat            = dataType{"heat", 0x1D, 1, "", 1, TypeBool}
	Light           = dataType{"light", 0x1E, 1, "", 1, TypeBool}
	Lock            = dataType{"lock", 0x1F, 1, "", 1, TypeBool}
	Moisture        = dataType{"moisture", 0x20, 1, "", 1, TypeBool}
	Motion          = dataType{"motion", 0x21, 1, "", 1, TypeBool}
	Moving          = dataType{"moving", 0x22, 1, "", 1, TypeBool}
	Occupancy       = dataType{"occupancy", 0x23, 1, "", 1, TypeBool}
	Opening         = dataType{"opening", 0x11, 1, "", 1, TypeBool}
	Plug            = dataType{"plug", 0x24, 1, "", 1, TypeBool}
	Powered         = dataType{"power", 0x10, 1, "", 1, TypeBool}
	Presence        = dataType{"presence", 0x25, 1, "", 1, TypeBool}
	Problem         = dataType{"problem", 0x26, 1, "", 1, TypeBool}
	Running         = dataType{"running", 0x27, 1, "", 1, TypeBool}
	Safety          = dataType{"safety", 0x28, 1, "", 1, TypeBool}
	Smoke           = dataType{"smoke", 0x29, 1, "", 1, TypeBool}
	Sound           = dataType{"sound", 0x2A, 1, "", 1, TypeBool}
	Tamper          = dataType{"tamper", 0x2B, 1, "", 1, TypeBool}
	Vibration       = dataType{"vibration", 0x2C, 1, "", 1, TypeBool}
	Window          = dataType{"window", 0x2D, 1, "", 1, TypeBool}
)

var (
	DataTypes = []dataType{
		Acceleration,
		Battery,
		CO2,
		Conductivity,
		Count8,
		Count16,
		Count32,
		CountSint8,
		CountSint16,
		CountSint32,
		Current,
		CurrentSint16,
		Dewpoint,
		DistanceMM,
		DistanceM,
		Duration,
		Energy,
		Energy24,
		Gas24,
		Gas32,
		Gyroscope,
		Humidity16,
		Humidity8,
		Illuminance,
		MassKG,
		MassLB,
		Moisture16,
		Moisture8,
		PM25,
		PM10,
		Power,
		PowerSint32,
		Pressure,
		Raw,
		Rotation,
		Speed,
		TemperatureSint8,
		TemperatureSint8_35,
		Temperature16,
		Temperature16_01,
		Text,
		Timestamp,
		TVOC,
		Voltage,
		Voltage10,
		Volume,
		Volume16,
		VolumeML,
		VolumeStorage,
		VolumeFlowRate,
		UVIndex,
		Water,
		BatteryLow,
		BatteryCharging,
		CarbonMonoxide,
		Cold,
		Connectivity,
		Door,
		GarageDoor,
		Gas,
		GenericBoolean,
		Heat,
		Light,
		Lock,
		Moisture,
		Motion,
		Moving,
		Occupancy,
		Opening,
		Plug,
		Powered,
		Presence,
		Problem,
		Running,
		Safety,
		Smoke,
		Sound,
		Tamper,
		Vibration,
		Window,
	}
)

func FindDataType(id byte) DataType {
	for _, t := range DataTypes {
		if t.ID() == id {
			return &t
		}
	}
	return nil
}
