package connection_test

import (
	"context"
	nmcli_go "github.com/KunMengcode/nmcli-go"
	"github.com/KunMengcode/nmcli-go/connection"
	"testing"
)

func TestManager_Up(t *testing.T) {
	m := nmcli_go.NewNMCli()
	out, err := m.Connection.Up(context.Background(), "hotspot", connection.UpOptions{})
	if err != nil {
		return
	}
	t.Log(out)
}

func TestManager_Show(t *testing.T) {
	m := nmcli_go.NewNMCli()
	out, err := m.Connection.Show(context.Background(), "hotspot")
	if err != nil {
		return
	}
	t.Log(out)
}
