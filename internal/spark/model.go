package spark

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/vanvanni/virtuit/internal/cell"
	"github.com/vanvanni/virtuit/internal/link"
	"github.com/vanvanni/virtuit/internal/resources"
	"github.com/vanvanni/virtuit/internal/support"
)

type Spark struct {
	Name      string `toml:"name"`
	Kernel    string `toml:"kernel"`
	Cmd       string `toml:"cmd"` // Command?
	Autostart bool   `toml:"autostart"`

	Resources resources.ResourceSpec `toml:"resources"`
	Links     []link.LinkSpec        `toml:"links"`
	Cells     []cell.CellSpec        `toml:"cells"`

	ID       string `toml:"-"`
	SocketID string `toml:"-"`
	MAC      string `toml:"-"`
}

func CreateSpark(cfg *Spark) (*Spark, error) {
	// >> UUID
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %w", err)
	}
	cfg.ID = uid.String()
	uidParts := strings.Split(cfg.ID, "-")

	// >> Socket
	cfg.SocketID = uidParts[0] + uidParts[1][0:1]

	// >> MAC address
	mac := support.GenerateMac()
	cfg.MAC = mac.String()

	if cfg.Kernel == "default" || cfg.Kernel == "" {
		cfg.Kernel = "/root/staticfyer/test/os/vmlinux"
	}

	for i, c := range cfg.Cells {
		if c.Boot && (c.Path == "default" || c.Path == "") {
			cloned := fmt.Sprintf("/var/virtuit/cells/%s-rootfs.ext4", cfg.SocketID)
			if err := support.CopyFile("/root/staticfyer/test/os/rootfs.ext4", cloned); err != nil {
				return nil, fmt.Errorf("failed to clone boot cell: %w", err)
			}
			cfg.Cells[i].Path = cloned
		}
	}

	return cfg, nil
}
