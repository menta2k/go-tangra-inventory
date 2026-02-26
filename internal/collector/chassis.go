package collector

import "github.com/siderolabs/go-smbios/smbios"

// collectChassisInfo extracts system enclosure details from SMBIOS Type 3.
func collectChassisInfo(s *smbios.SMBIOS) ChassisInfo {
	se := s.SystemEnclosure
	return ChassisInfo{
		Manufacturer:   se.Manufacturer,
		Version:        se.Version,
		SerialNumber:   se.SerialNumber,
		AssetTagNumber: se.AssetTagNumber,
		SKUNumber:      se.SKUNumber,
	}
}
