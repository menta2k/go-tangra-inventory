package collector

import "time"

// Inventory holds the complete hardware inventory of a host.
type Inventory struct {
	CollectedAt   time.Time        `json:"collected_at"`
	Hostname      string           `json:"hostname"`
	Username      string           `json:"username"`
	SMBIOSVersion VersionInfo      `json:"smbios_version"`
	BIOS          BIOSInfo         `json:"bios"`
	System        SystemInfo       `json:"system"`
	Baseboard     BaseboardInfo    `json:"baseboard"`
	Chassis       ChassisInfo      `json:"chassis"`
	Processors    []ProcessorInfo  `json:"processors"`
	Cache         []CacheInfo      `json:"cache,omitempty"`
	Memory        MemoryInfo       `json:"memory"`
	Ports         []PortInfo       `json:"ports,omitempty"`
	Slots         []SlotInfo       `json:"slots,omitempty"`
	OEMStrings    []string         `json:"oem_strings,omitempty"`
	BIOSLanguage  BIOSLanguageInfo `json:"bios_language,omitempty"`
	Monitor       []MonitorInfo    `json:"monitor,omitempty"`
}

// VersionInfo holds the SMBIOS specification version.
type VersionInfo struct {
	Major    int `json:"major"`
	Minor    int `json:"minor"`
	Revision int `json:"revision"`
}

// BIOSInfo holds BIOS vendor, version, and release date (Type 0).
type BIOSInfo struct {
	Vendor      string `json:"vendor"`
	Version     string `json:"version"`
	ReleaseDate string `json:"release_date"`
}

// SystemInfo holds system manufacturer, product, serial, and UUID (Type 1).
type SystemInfo struct {
	Manufacturer string `json:"manufacturer"`
	ProductName  string `json:"product_name"`
	Version      string `json:"version"`
	SerialNumber string `json:"serial_number"`
	UUID         string `json:"uuid"`
	WakeUpType   string `json:"wake_up_type"`
	SKUNumber    string `json:"sku_number"`
	Family       string `json:"family"`
}

// BaseboardInfo holds baseboard/motherboard details (Type 2).
type BaseboardInfo struct {
	Manufacturer      string `json:"manufacturer"`
	Product           string `json:"product"`
	Version           string `json:"version"`
	SerialNumber      string `json:"serial_number"`
	AssetTag          string `json:"asset_tag"`
	LocationInChassis string `json:"location_in_chassis,omitempty"`
	BoardType         string `json:"board_type"`
}

// ChassisInfo holds system enclosure/chassis details (Type 3).
type ChassisInfo struct {
	Manufacturer   string `json:"manufacturer"`
	Version        string `json:"version"`
	SerialNumber   string `json:"serial_number"`
	AssetTagNumber string `json:"asset_tag_number"`
	SKUNumber      string `json:"sku_number"`
}

// ProcessorInfo holds processor details (Type 4).
type ProcessorInfo struct {
	SocketDesignation string `json:"socket_designation"`
	Manufacturer      string `json:"manufacturer"`
	Version           string `json:"version"`
	MaxSpeedMHz       uint16 `json:"max_speed_mhz"`
	CurrentSpeedMHz   uint16 `json:"current_speed_mhz"`
	SocketPopulated   bool   `json:"socket_populated"`
	SerialNumber      string `json:"serial_number"`
	AssetTag          string `json:"asset_tag"`
	PartNumber        string `json:"part_number"`
	CoreCount         uint8  `json:"core_count"`
	CoreEnabled       uint8  `json:"core_enabled"`
	ThreadCount       uint8  `json:"thread_count"`
}

// CacheInfo holds cache designation (Type 7).
type CacheInfo struct {
	SocketDesignation string `json:"socket_designation"`
}

// MemoryInfo holds total physical memory and per-module details.
type MemoryInfo struct {
	TotalPhysicalBytes uint64              `json:"total_physical_bytes"`
	TotalPhysicalGB    float64             `json:"total_physical_gb"`
	Array              PhysicalMemoryArray `json:"array"`
	Modules            []MemoryModule      `json:"modules,omitempty"`
}

// PhysicalMemoryArray holds the memory array metadata (Type 16).
type PhysicalMemoryArray struct {
	Location              string `json:"location"`
	Use                   string `json:"use"`
	ErrorCorrection       string `json:"error_correction"`
	MaximumCapacity       string `json:"maximum_capacity"`
	NumberOfMemoryDevices uint16 `json:"number_of_memory_devices"`
}

// MemoryModule holds details for a single physical memory DIMM (Type 17).
type MemoryModule struct {
	DeviceLocator      string `json:"device_locator"`
	BankLocator        string `json:"bank_locator"`
	CapacityBytes      uint64 `json:"capacity_bytes"`
	FormFactor         string `json:"form_factor"`
	MemoryType         string `json:"memory_type"`
	TypeDetail         string `json:"type_detail,omitempty"`
	SpeedMTs           uint16 `json:"speed_mt_s"`
	ConfiguredSpeedMTs uint16 `json:"configured_speed_mt_s"`
	Manufacturer       string `json:"manufacturer"`
	SerialNumber       string `json:"serial_number"`
	AssetTag           string `json:"asset_tag"`
	PartNumber         string `json:"part_number"`
	MinimumVoltage     string `json:"minimum_voltage,omitempty"`
	MaximumVoltage     string `json:"maximum_voltage,omitempty"`
	ConfiguredVoltage  string `json:"configured_voltage,omitempty"`
	TotalWidthBits     string `json:"total_width"`
	DataWidthBits      string `json:"data_width"`
}

// PortInfo holds port connector details (Type 8).
type PortInfo struct {
	InternalDesignator string `json:"internal_designator"`
	ExternalDesignator string `json:"external_designator"`
}

// SlotInfo holds system slot details (Type 9).
type SlotInfo struct {
	Designation string `json:"designation"`
}

// BIOSLanguageInfo holds BIOS language settings (Type 13).
type BIOSLanguageInfo struct {
	CurrentLanguage      string   `json:"current_language,omitempty"`
	InstallableLanguages []string `json:"installable_languages,omitempty"`
}

// MonitorInfo holds connected display details.
type MonitorInfo struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	SerialNumber string `json:"serial_number"`
}
