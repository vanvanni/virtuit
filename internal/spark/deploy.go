package spark

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/vanvanni/virtuit/internal/out"
	"github.com/vanvanni/virtuit/internal/support"
)

func Deploy(cfg Spark) error {
	out.Logger.Info(fmt.Sprintf("Preparing Spark: %+v", cfg))

	_, socketPath, err := StartFirecracker(cfg.SocketID)
	if err != nil {
		return fmt.Errorf("failed to start Firecracker: %w", err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
				return net.Dial("unix", socketPath)
			},
		},
	}

	bootArgs := cfg.Cmd
	if bootArgs == "" {
		bootArgs = "console=ttyS0 reboot=k panic=1 pci=off"
	}
	bootSource := map[string]interface{}{
		"kernel_image_path": cfg.Kernel,
		"boot_args":         bootArgs,
	}
	if err := support.SendFirecrackerRequest(client, "PUT", "/boot-source", bootSource); err != nil {
		return fmt.Errorf("boot-source config failed: %w", err)
	}

	for i, cell := range cfg.Cells {
		block := map[string]interface{}{
			"drive_id":       fmt.Sprintf("cell%d", i),
			"path_on_host":   cell.Path,
			"is_root_device": cell.Boot,
			"is_read_only":   false,
		}
		fmt.Println(block)
		if err := support.SendFirecrackerRequest(client, "PUT", fmt.Sprintf("/drives/cell%d", i), block); err != nil {
			return fmt.Errorf("cell config failed: %w", err)
		}
	}

	for _, link := range cfg.Links {
		iface := map[string]interface{}{
			"iface_id":      link.Name,
			"host_dev_name": link.Interface,
			"guest_mac":     cfg.MAC,
		}

		if err := support.SendFirecrackerRequest(client, "PUT", "/network-interfaces/"+link.Name, iface); err != nil {
			return fmt.Errorf("network config failed: %w", err)
		}
	}

	config := map[string]interface{}{
		"vcpu_count":   cfg.Resources.CPU,
		"mem_size_mib": cfg.Resources.Mem,
		"ht_enabled":   "smt",
	}
	if err := support.SendFirecrackerRequest(client, "PUT", "/machine-config", config); err != nil {
		return fmt.Errorf("machine config failed: %w", err)
	}

	if err := support.SendFirecrackerRequest(client, "PUT", "/actions", map[string]interface{}{
		"action_type": "InstanceStart",
	}); err != nil {
		return fmt.Errorf("failed to start instance: %w", err)
	}

	out.Logger.Info("Spark successfully started.")
	return nil
}
