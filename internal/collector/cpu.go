package collector

import "github.com/siderolabs/go-smbios/smbios"

// collectProcessorInfo extracts processor details from SMBIOS Type 4.
func collectProcessorInfo(s *smbios.SMBIOS) []ProcessorInfo {
	var result []ProcessorInfo
	for _, p := range s.ProcessorInformation {
		result = append(result, ProcessorInfo{
			SocketDesignation: p.SocketDesignation,
			Manufacturer:     p.ProcessorManufacturer,
			Version:          p.ProcessorVersion,
			MaxSpeedMHz:      p.MaxSpeed,
			CurrentSpeedMHz:  p.CurrentSpeed,
			SocketPopulated:  p.Status.SocketPopulated(),
			SerialNumber:     p.SerialNumber,
			AssetTag:         p.AssetTag,
			PartNumber:       p.PartNumber,
			CoreCount:        p.CoreCount,
			CoreEnabled:      p.CoreEnabled,
			ThreadCount:      p.ThreadCount,
		})
	}
	return result
}
