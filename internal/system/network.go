package system

import (
	"net"

	"github.com/example/sysinfo-cli/internal/models"
)

// GetNetworkInfo returns network interface information
func GetNetworkInfo() ([]models.NetworkInterface, error) {
	interfaces := make([]models.NetworkInterface, 0)

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		status := "down"
		if iface.Flags&net.FlagUp != 0 {
			status = "up"
		}

		var ips []string
		addrs, err := iface.Addrs()
		if err == nil {
			for _, addr := range addrs {
				ips = append(ips, addr.String())
			}
		}

		interfaces = append(interfaces, models.NetworkInterface{
			Name:        iface.Name,
			IPAddresses: ips,
			MACAddress:  iface.HardwareAddr.String(),
			MTU:         iface.MTU,
			Status:      status,
		})
	}

	return interfaces, nil
}
