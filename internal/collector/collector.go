package collector

import (
	"fmt"
	"os"
	"time"

	"github.com/siderolabs/go-smbios/smbios"
)

// Collect gathers a full hardware inventory from the local host
// using SMBIOS data.
func Collect() (*Inventory, error) {
	hostname, _ := os.Hostname()

	inv := &Inventory{
		CollectedAt: time.Now().UTC(),
		Hostname:    hostname,
	}
	monitorInfo, err := CollectMonitorInfo()
	if err != nil {
		fmt.Printf("warning: cannot collect monitor info: %v\n", err)
	} else {
		inv.Monitor = monitorInfo
	}
	userName, err := GetUserInfo()
	if err != nil {
		fmt.Printf("warning: cannot collect user info: %v\n", err)
	} else {
		inv.Username = userName
	}
	s, err := smbios.New()
	if err != nil {
		return inv, fmt.Errorf("opening SMBIOS: %w", err)
	}

	inv.SMBIOSVersion = VersionInfo{
		Major:    s.Version.Major,
		Minor:    s.Version.Minor,
		Revision: s.Version.Revision,
	}
	inv.BIOS = collectBIOSInfo(s)
	inv.System = collectSystemInfo(s)
	inv.Baseboard = collectBaseboardInfo(s)
	inv.Chassis = collectChassisInfo(s)
	inv.Processors = collectProcessorInfo(s)
	inv.Memory = collectMemoryInfo(s)

	// Cache (Type 7)
	for _, c := range s.CacheInformation {
		inv.Cache = append(inv.Cache, CacheInfo{
			SocketDesignation: c.SocketDesignation,
		})
	}

	// Port connectors (Type 8)
	for _, p := range s.PortConnectorInformation {
		inv.Ports = append(inv.Ports, PortInfo{
			InternalDesignator: p.InternalReferenceDesignator,
			ExternalDesignator: p.ExternalReferenceDesignator,
		})
	}

	// System slots (Type 9)
	for _, sl := range s.SystemSlots {
		inv.Slots = append(inv.Slots, SlotInfo{
			Designation: sl.SlotDesignation,
		})
	}

	// OEM strings (Type 11)
	inv.OEMStrings = s.OEMStrings.Strings

	// BIOS language (Type 13)
	inv.BIOSLanguage = BIOSLanguageInfo{
		CurrentLanguage:      s.BIOSLanguageInformation.CurrentLanguage,
		InstallableLanguages: s.BIOSLanguageInformation.InstallableLanguages,
	}

	return inv, nil
}
