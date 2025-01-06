package connection

import (
	"context"
	"fmt"
	"github.com/KunMengcode/nmcli-go/utils"
)

type UpOptions struct {
	Ifname      string `json:"ifname"`
	BSSID       string `json:"ap"`
	Passwd_File string `json:"passwd-file"`
}

func (m Manager) Up(ctx context.Context, ID string, args UpOptions) (string, error) {
	cmdArgs := []string{"up", ID}
	cmdArgs = append(cmdArgs, utils.Marshal(args)...)
	output, err := m.CommandContext(ctx, nmcliCmd, cmdArgs...).Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute nmcli with args %+q: %w", cmdArgs, err)
	}
	return string(output), nil
}

func (m Manager) Modify(ctx context.Context, temporary bool, ID string, option map[string]string) (string, error) {
	cmdArgs := []string{"modify"}
	if temporary {
		cmdArgs = append(cmdArgs, "--temporary")
	}
	cmdArgs = append(cmdArgs, ID)
	for k, v := range option {
		cmdArgs = append(cmdArgs, k, v)
	}
	output, err := m.CommandContext(ctx, nmcliCmd, cmdArgs...).Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute nmcli with args %+q: %w", cmdArgs, err)
	}
	return string(output), nil
}
