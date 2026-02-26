package collector

import (
	"strings"

	"github.com/siderolabs/go-smbios/smbios"
)

// collectMemoryInfo extracts memory array (Type 16) and
// per-DIMM details (Type 17) from SMBIOS.
func collectMemoryInfo(s *smbios.SMBIOS) MemoryInfo {
	pma := s.PhysicalMemoryArray

	info := MemoryInfo{
		Array: PhysicalMemoryArray{
			Location:              pma.Location.String(),
			Use:                   pma.Use.String(),
			ErrorCorrection:       pma.MemoryErrorCorrection.String(),
			MaximumCapacity:       pma.MaximumCapacity.String(),
			NumberOfMemoryDevices: pma.NumberOfMemoryDevices,
		},
	}

	var totalBytes uint64

	for _, d := range s.MemoryDevices {
		sizeMB := d.Size.Megabytes()
		if sizeMB == 0 {
			continue // empty slot
		}

		// Handle extended size for DIMMs > 32GB-1MB
		capBytes := uint64(sizeMB) * 1024 * 1024
		if uint16(d.Size) == 0x7FFF {
			capBytes = uint64(d.ExtendedSize) * 1024 * 1024
		}
		totalBytes += capBytes

		info.Modules = append(info.Modules, MemoryModule{
			DeviceLocator:      d.DeviceLocator,
			BankLocator:        d.BankLocator,
			CapacityBytes:      capBytes,
			FormFactor:         d.FormFactor.String(),
			MemoryType:         d.MemoryType.String(),
			TypeDetail:         d.TypeDetail.String(),
			SpeedMTs:           uint16(d.Speed),
			ConfiguredSpeedMTs: uint16(d.ConfiguredMemorySpeed),
			Manufacturer:       strings.TrimSpace(d.Manufacturer),
			SerialNumber:       strings.TrimSpace(d.SerialNumber),
			AssetTag:           strings.TrimSpace(d.AssetTag),
			PartNumber:         strings.TrimSpace(d.PartNumber),
			MinimumVoltage:     d.MinimumVoltage.String(),
			MaximumVoltage:     d.MaximumVoltage.String(),
			ConfiguredVoltage:  d.ConfiguredVoltage.String(),
			TotalWidthBits:     d.TotalWidth.String(),
			DataWidthBits:      d.DataWidth.String(),
		})
	}

	info.TotalPhysicalBytes = totalBytes
	info.TotalPhysicalGB = float64(totalBytes) / (1024 * 1024 * 1024)

	return info
}
