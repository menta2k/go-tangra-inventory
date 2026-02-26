package collector

import "github.com/siderolabs/go-smbios/smbios"

// collectSystemInfo extracts system identification from SMBIOS Type 1.
func collectSystemInfo(s *smbios.SMBIOS) SystemInfo {
	si := s.SystemInformation
	return SystemInfo{
		Manufacturer: si.Manufacturer,
		ProductName:  si.ProductName,
		Version:      si.Version,
		SerialNumber: si.SerialNumber,
		UUID:         si.UUID,
		WakeUpType:   si.WakeUpType.String(),
		SKUNumber:    si.SKUNumber,
		Family:       si.Family,
	}
}
