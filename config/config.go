package config

import (
	"os"
	"time"
)

// Environment-based configuration
var (
	DbHost     = getEnv("DB_HOST", "localhost")
	DbPort     = getEnv("DB_PORT", "5432")
	DbUser     = getEnv("DB_USER", "admin")
	DbPassword = getEnv("DB_PASSWORD", "admin")
	DbName     = getEnv("DB_NAME", "monitoring_db_go")

	MqttBroker = getEnv("MQTT_BROKER", "192.168.1.153")
	MqttPort   = getEnv("MQTT_PORT", "1883")

	MqttTopics = []string{
		"balluff/cmtk/master1/iolink/devices/port1/data/fromdevice",
		"balluff/cmtk/master1/iolink/devices/port2/data/fromdevice",
		"balluff/cmtk/master1/iolink/devices/port3/data/fromdevice",
		"balluff/cmtk/master1/iolink/devices/port4/data/fromdevice",
		"balluff/cmtk/master2/iolink/devices/port0/data/fromdevice",
		"balluff/cmtk/master2/iolink/devices/port1/data/fromdevice",
		"balluff/cmtk/master2/iolink/devices/port2/data/fromdevice",
	}

	FlowPorts = []string{
	"master1/port3",
	"master1/port4",
	"master2/port0",
	"master2/port1",
	"master2/port2",
	}

	AnalyzerIPs = []string{
	getEnv("ANALYZER_IP01", "192.168.1.130"),
	getEnv("ANALYZER_IP02", "192.168.1.131"),
	getEnv("ANALYZER_IP03", "192.168.1.132"),
	}
)

// File paths and intervals
const (
	IdleTimeoutSeconds     = 5
	ProductionCycleDefault = 14.0
	FlowDeviceCount        = 5
	MeasurementDeviceCount = 5

	IntervalMQTTData          = 50 * time.Millisecond
	IntervalRestData 		  = 100 * time.Millisecond
	FlowUpdateInterval        = 10 * time.Second
	MeasurementUpdateInterval = 10 * time.Second
	MetersUpdateInterval      = 5 * time.Second
	OEEUpdateInterval		  = 5 * time.Second
	
	AirFactor = 1.0
	ImpulsyNaObrot        = 8
	ElementWindow         = 10
	MaxChangeoverDuration = 10 * 60 // 10 minut w sekundach

	SummaryFilePath         = "logs/summary.json"
	OeeFilePath             = "logs/oee.json"
	MqttOeeFilePath         = "logs/mqttOEE.json"
	MqttFlowFilePath        = "logs/mqttFlow.json"
	MeasurementFilePath     = "logs/measurements.json"
	MetersFilePath          = "logs/meters.json"
	FakeMeasurementFilePath = "logs/template/measurements_multi.json"
	FakeMetersFilePath      = "logs/template/meters_multi.json"
	SystemLogPath           = "logs/system.log"
	DefaultJsonFile         = "logs/system_report.json"


)

var JsonWithBackup = map[string]bool{
    OeeFilePath:      true,
    SummaryFilePath:  true,
}

// CycleRule defines rules for dynamic cycle assignment
type CycleRule struct {
	MaxLength int
	MaxWidth  int
	CycleLPM  float64
}

// From excel file cycles
var CycleTable = []CycleRule{
	{MaxLength: 600, MaxWidth: 9999, CycleLPM: 15.0},    // <600 mm → 4s
	{MaxLength: 800, MaxWidth: 9999, CycleLPM: 12.875},  // <800 mm → ~4.66s
	{MaxLength: 1200, MaxWidth: 9999, CycleLPM: 12.0},   // <1200 mm → 5s
	{MaxLength: 99999, MaxWidth: 9999, CycleLPM: 7.06},  // >=1200 mm → 8.5s
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
