package collector

import "github.com/siderolabs/go-smbios/smbios"

// collectBaseboardInfo extracts baseboard details from SMBIOS Type 2.
func collectBaseboardInfo(s *smbios.SMBIOS) BaseboardInfo {
	bb := s.BaseboardInformation
	return BaseboardInfo{
		Manufacturer:    bb.Manufacturer,
		Product:         bb.Product,
		Version:         bb.Version,
		SerialNumber:    bb.SerialNumber,
		AssetTag:        bb.AssetTag,
		LocationInChassis: bb.LocationInChassis,
		BoardType:       bb.BoardType.String(),
	}
}
