package device

import (
	"context"
	"fmt"
	"github.com/KunMengcode/nmcli-go/utils"
	"strings"
)

type Info struct {
	Device          string
	Type            string
	State           string
	IP4Connectivity string
	IP6Connectivity string
	DbusPath        string
	Connection      string
	ConUUID         string
	ConPath         string
}

// Status lists the status for all devices.
func (m Manager) Show(ctx context.Context, DeviceInterfaceName string) (any, error) {
	fields := []string{"GENERAL", "CAPABILITIES", "INTERFACE-FLAGS", "WIFI-PROPERTIES", "AP", "WIRED-PROPERTIES", "WIMAX-PROPERTIES", "NSP", "IP4", "DHCP4", "IP6", "DHCP6", "BOND", "TEAM", "BRIDGE", "VLAN", "BLUETOOTH", "CONNECTIONS"}
	cmdArgs := []string{"-g", strings.Join(fields, ",")}
	cmdArgs = append(cmdArgs, "device", "show")
	if DeviceInterfaceName != "" {
		cmdArgs = append(cmdArgs, DeviceInterfaceName)
	}
	output, err := m.CommandContext(ctx, nmcliCmd, cmdArgs...).Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute nmcli with args %+q: %w", cmdArgs, err)
	}

	if DeviceInterfaceName == "" {
		return utils.ParseCmdsHaveFieldNameOutput(output), nil
	}
	return utils.ParseCmdHaveFieldNameOutput(output), nil
}
