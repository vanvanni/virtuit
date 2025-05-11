package types

type (
	VirtuKernel struct {
	}

	VirtuDrive struct {
	}

	VirtuSocket struct {
	}

	VirtuNetwork struct {
	}

	VirtuMachine struct {
		id         string         `json:"id"`
		Name       string         `json:"name"`
		KernelPath VirtuKernel    `json:"kernel_path"`
		Drives     []VirtuDrive   `json:"drives"`   // Root Disk (rootfs) -- Do we need to add more disks?
		Networks   []VirtuNetwork `json:"networks"` // Tap Devices
		MacAddress string         `json:"mac_address"`
		Cpus       int            `json:"vcpu_count"`
		Memory     int            `json:"memory_mib"`
		Socket     VirtuSocket    `json:"socket"`
		Status     string         `json:"status"`
	}
)
