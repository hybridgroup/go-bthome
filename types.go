package bthome

import "tinygo.org/x/bluetooth"

var ServiceUUID = bluetooth.New16BitUUID(0xFCD2)

// DataType represents a data type that can be added to the service data payload.
type DataType struct {
	// Name of the data type
	Name string

	// ID of the data type
	ID byte

	// Size of the data type in bytes
	Size int
}

// See this list of field types:
// https://bthome.io/format/
var (
	Acceleration        = DataType{"acceleration", 0x51, 2}
	Battery             = DataType{"battery", 0x01, 1}
	CO2                 = DataType{"co2", 0x12, 2}
	Conductivity        = DataType{"conductivity", 0x56, 2}
	Count8              = DataType{"count", 0x09, 1}
	Count16             = DataType{"count", 0x3D, 2}
	Count32             = DataType{"count", 0x3E, 4}
	CountSint8          = DataType{"count", 0x59, 1}
	CountSint16         = DataType{"count", 0x5A, 2}
	CountSint32         = DataType{"count", 0x5B, 4}
	Current             = DataType{"current", 0x43, 2}
	CurrentSint16       = DataType{"current", 0x5D, 2}
	Dewpoint            = DataType{"dewpoint", 0x08, 2}
	DistanceMM          = DataType{"distance (mm)", 0x40, 2}
	DistanceM           = DataType{"distance (m)", 0x41, 2}
	Duration            = DataType{"duration", 0x42, 3}
	Energy              = DataType{"energy", 0x4D, 4}
	Energy24            = DataType{"energy", 0x0A, 3}
	Gas24               = DataType{"gas", 0x4B, 3}
	Gas32               = DataType{"gas", 0x4C, 4}
	Gyroscope           = DataType{"gyroscope", 0x52, 2}
	Humidity16          = DataType{"humidity", 0x03, 2}
	Humidity8           = DataType{"humidity", 0x2E, 1}
	Illuminance         = DataType{"illuminance", 0x05, 3}
	MassKG              = DataType{"mass (kg)", 0x06, 2}
	MassLB              = DataType{"mass (lb)", 0x07, 2}
	Moisture16          = DataType{"moisture", 0x14, 2}
	Moisture8           = DataType{"moisture", 0x2F, 1}
	PM25                = DataType{"pm2.5", 0x0D, 2}
	PM10                = DataType{"pm10", 0x0E, 2}
	Power               = DataType{"power", 0x0B, 3}
	PowerSint32         = DataType{"power", 0x5C, 4}
	Pressure            = DataType{"pressure", 0x04, 3}
	Raw                 = DataType{"raw", 0x54, -1}
	Rotation            = DataType{"rotation", 0x3F, 2}
	Speed               = DataType{"speed", 0x44, 2}
	TemperatureSint8    = DataType{"temperature", 0x57, 1}
	TemperatureSint8_35 = DataType{"temperature", 0x58, 1}
	Temperature16       = DataType{"temperature", 0x45, 2}
	Temperature16_01    = DataType{"temperature", 0x02, 2}
	Text                = DataType{"text", 0x53, -1}
	Timestamp           = DataType{"timestamp", 0x50, 4}
	TVOC                = DataType{"tvoc", 0x13, 2}
	Voltage             = DataType{"voltage", 0x0C, 2}
	Voltage10           = DataType{"voltage", 0x4A, 2}
	Volume              = DataType{"volume", 0x4E, 4}
	Volume16            = DataType{"volume", 0x47, 2}
	VolumeML            = DataType{"volume", 0x48, 2}
	VolumeStorage       = DataType{"volume storage", 0x55, 4}
	VolumeFlowRate      = DataType{"volume flow rate", 0x49, 2}
	UVIndex             = DataType{"UV index", 0x46, 1}
	Water               = DataType{"water", 0x4F, 4}
)
