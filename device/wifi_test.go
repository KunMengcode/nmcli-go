package device_test

import (
	"context"
	"fmt"
	nmcli_go "github.com/KunMengcode/nmcli-go"
	"github.com/KunMengcode/nmcli-go/device"
	"testing"
)

func TestManager_WiFiList(t *testing.T) {
	m := nmcli_go.NewNMCli()
	wifis, err := m.Device.WiFiList(context.Background(), device.WiFiListOptions{
		Rescan: device.WiFiListOptionsRescan.Auto,
		IfName: "wlo1",
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, wifi := range wifis {
		t.Log(wifi.SSID, wifi.BSSID, wifi.Signal)
	}
}

func TestManager_WiFiConnect(t *testing.T) {
	m := nmcli_go.NewNMCli()
	status, err := m.Device.WiFiConnect(context.Background(), "50:FA:84:3F:09:29", device.WiFiConnectOptions{
		Password: "omat12345",
	})
	if err != nil {
		return
	}
	fmt.Println(status)
}

func TestManager_WiFiHotspotCreate(t *testing.T) {
	m := nmcli_go.NewNMCli()
	status, err := m.Device.WiFiHotspotCreate(context.Background(), device.WiFiHotspotCreateOptions{
		Ifname:   "wlo1",
		Con_name: "hotspot",
		SSID:     "Hello",
		Band:     device.WifiHotspotBand.Use2_4G,
		Password: "123456789",
		Channel:  "1",
	})
	if err != nil {
		return
	}
	println(status)
}
