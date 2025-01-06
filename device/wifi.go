package device

import (
	"context"
	"fmt"
	"strings"

	"github.com/KunMengcode/nmcli-go/utils"
)

type WiFiListOptions struct {
	IfName string `ArgsPre:"ifname"`
	BSSID  string `ArgsPre:"bssid"`
	Rescan string `ArgsPre:"--rescan"`
}

var WiFiListOptionsRescan = struct {
	Auto string
	Yes  string
	No   string
}{
	Auto: "auto",
	Yes:  "yes",
	No:   "no",
}

var WifiHotspotBand = struct {
	Use2_4G string
	Use5G   string
}{
	Use2_4G: "bg",
	Use5G:   "a",
}

type WiFi struct {
	Name      string
	SSID      string
	SSIDHEX   string
	BSSID     string
	Mode      string
	Chan      string
	Frequency string
	Rate      string
	Signal    string
	Bars      string
	Security  string
	WPAFlags  string
	RSNFlags  string
	Device    string
	Active    string
	InUse     string
	DBusPath  string
}

// WiFiList list available Wi-Fi access points.
// The IfName and BSSID options can be used to list APs for a particular interface, or with a specific BSSID.
// The Rescan flag tells whether a new Wi-Fi scan should be triggered.
func (m Manager) WiFiList(ctx context.Context, args WiFiListOptions) ([]WiFi, error) {
	fields := []string{"NAME", "SSID", "SSID-HEX", "BSSID", "MODE", "CHAN", "FREQ", "RATE", "SIGNAL", "BARS", "SECURITY", "WPA-FLAGS", "RSN-FLAGS", "DEVICE", "ACTIVE", "IN-USE", "DBUS-PATH"}

	cmdArgs := []string{"-g", strings.Join(fields, ",")}
	cmdArgs = append(cmdArgs, "device", "wifi", "list")
	cmdArgs = append(cmdArgs, utils.Marshal(args)...)

	output, err := m.CommandContext(ctx, nmcliCmd, cmdArgs...).Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute nmcli with args %+q: %w", cmdArgs, err)
	}

	parsedOutput, err := utils.ParseCmdOutput(output, len(fields))
	if err != nil {
		return nil, fmt.Errorf("failed to parse nmcli output: %w", err)
	}

	wifis := make([]WiFi, len(parsedOutput))
	for i, fields := range parsedOutput {
		wifis[i] = WiFi{
			Name:      fields[0],
			SSID:      fields[1],
			SSIDHEX:   fields[2],
			BSSID:     fields[3],
			Mode:      fields[4],
			Chan:      fields[5],
			Frequency: fields[6],
			Rate:      fields[7],
			Signal:    fields[8],
			Bars:      fields[9],
			Security:  fields[10],
			WPAFlags:  fields[11],
			RSNFlags:  fields[12],
			Device:    fields[13],
			Active:    fields[14],
			InUse:     fields[15],
			DBusPath:  fields[16],
		}
	}

	return wifis, nil
}

type WiFiConnectOptions struct {
	Password   string                       `ArgsPre:"password"`
	WEPKeyType WiFiConnectOptionsWEPKeyType `ArgsPre:"wep-key-type"`
	IfName     string                       `ArgsPre:"ifname"`
	BSSID      string                       `ArgsPre:"bssid"`
	Name       string                       `ArgsPre:"name"`
	Private    WiFiConnectOptionsPrivate    `ArgsPre:"private"`
	Hidden     WiFiConnectOptionsHidden     `ArgsPre:"hidden"`
}

type WiFiHotspotCreateOptions struct {
	Ifname   string `ArgsPre:"ifname"`
	Con_name string `ArgsPre:"con-name"`
	SSID     string `ArgsPre:"ssid"`
	Band     string `ArgsPre:"band"`
	Password string `ArgsPre:"password"`
	Channel  string `ArgsPre:"channel"`
}

type (
	WiFiConnectOptionsWEPKeyType string
	WiFiConnectOptionsPrivate    string
	WiFiConnectOptionsHidden     string
)

const (
	WiFiConnectOptionsWEPKeyTypeKey    WiFiConnectOptionsWEPKeyType = "key"
	WiFiConnectOptionsWEPKeyTypePhrase WiFiConnectOptionsWEPKeyType = "phrase"
	WiFiConnectOptionsPrivateYes       WiFiConnectOptionsPrivate    = "yes"
	WiFiConnectOptionsPrivateNo        WiFiConnectOptionsPrivate    = "no"
	WiFiConnectOptionsHiddenYes        WiFiConnectOptionsHidden     = "yes"
	WiFiConnectOptionsHiddenNo         WiFiConnectOptionsHidden     = "no"
)

// WiFiConnect Connect to a Wi-Fi network specified by BSSID which could also be a SSID.
// The command finds a matching connection or creates one and then activates it on a device.
// This is a command-line counterpart of clicking an SSID in a GUI client.
// If a connection for the network already exists, it is possible to bring up the existing profile as follows: nmcli con up id <name>.
// Note that only open, WEP and WPA-PSK networks are supported if no previous connection exists.
// It is also assumed that IP configuration is obtained via DHCP.
func (m Manager) WiFiConnect(ctx context.Context, BSSID string, args WiFiConnectOptions) (string, error) {
	cmdArgs := []string{"device", "wifi", "connect", BSSID}
	cmdArgs = append(cmdArgs, utils.Marshal(args)...)
	output, err := m.CommandContext(ctx, nmcliCmd, cmdArgs...).Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute nmcli with args %+q: %w", cmdArgs, err)
	}

	return string(output), nil
}

func (m Manager) WiFiHotspotCreate(ctx context.Context, args WiFiHotspotCreateOptions) (string, error) {
	cmdArgs := []string{"device", "wifi", "hotspot"}
	cmdArgs = append(cmdArgs, utils.Marshal(args)...)
	output, err := m.CommandContext(ctx, nmcliCmd, cmdArgs...).Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute nmcli with args %+q: %w", cmdArgs, err)
	}
	return string(output), nil
}
