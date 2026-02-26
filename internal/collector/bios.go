package collector

import "github.com/siderolabs/go-smbios/smbios"

// collectBIOSInfo extracts BIOS vendor information from SMBIOS Type 0.
func collectBIOSInfo(s *smbios.SMBIOS) BIOSInfo {
	bi := s.BIOSInformation
	return BIOSInfo{
		Vendor:      bi.Vendor,
		Version:     bi.Version,
		ReleaseDate: bi.ReleaseDate,
	}
}
