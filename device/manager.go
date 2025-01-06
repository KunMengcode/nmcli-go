package device

import (
	"context"

	"github.com/KunMengcode/nmcli-go/utils"
)

const nmcliCmd = "nmcli"

type Manager struct {
	CommandContext func(ctx context.Context, name string, args ...string) utils.Cmd
}
