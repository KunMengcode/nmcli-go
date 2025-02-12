English | [简体中文](./README_cn.md)
# nmcli-go

go wrapper for command line tool `nmcli`, The warehouse referenced [leberKleber/go-nmcli](https://github.com/leberKleber/go-nmcli)

## General

| original command            | library path                     | implemented                   |
|-----------------------------|----------------------------------|-------------------------------|
| `nmcli general status`      | `NMCli.General.Status(...)`      | :negative_squared_cross_mark: |
| `nmcli general hostname`    | `NMCli.General.Hostname(...)`    | :heavy_check_mark:            |
| `nmcli general permissions` | `NMCli.General.Permissions(...)` | :heavy_check_mark:            |
| `nmcli general logging`     | `NMCli.General.Logging(...)`     | :negative_squared_cross_mark: |

## Networking

| original command       | library path        | implemented                   |
|------------------------|---------------------|-------------------------------|
| `nmcli networking ...` | not implemented yet | :negative_squared_cross_mark: |

## Radio

| original command  | library path        | implemented                   |
|-------------------|---------------------|-------------------------------|
| `nmcli radio ...` | not implemented yet | :negative_squared_cross_mark: |

## Device

| original command                  | library path                         | implemented               |
|-----------------------------------|--------------------------------------|---------------------------|
| `nmcli device status`             | `NMCli.Device.Status(...)`           | :heavy_check_mark:        |
| `nmcli device show`               | `NMCli.Device.Show(...)`             | :heavy_check_mark:        |
| `nmcli device set`                | `NMCli.Device.Set(...)`              | :negative_squared_cross_mark: |
| `nmcli device reapply`            | `NMCli.Device.Reapply(...)`          | :negative_squared_cross_mark: |
| `nmcli device modify`             | `NMCli.Device.Modify(...)`           | :negative_squared_cross_mark: |
| `nmcli device disconnect`         | `NMCli.Device.Disconnect(...)`       | :negative_squared_cross_mark: |
| `nmcli device wifi list`          | `NMCli.Device.WiFiList(...)`         | :heavy_check_mark:        |
| `nmcli device wifi connect`       | `NMCli.Device.WiFiConnect(...)`      | :heavy_check_mark:        |
| `nmcli device wifi hotspot`       | `NMCli.Device.WiFiHotspot(...)`      | :heavy_check_mark:        |
| `nmcli device wifi rescan`        | `NMCli.Device.WiFiRescan(...)`       | :negative_squared_cross_mark: |
| `nmcli device wifi show-password` | `NMCli.Device.WiFiShowPassword(...)` | :negative_squared_cross_mark: |
| `nmcli device wifi lldp`          | `NMCli.Device.WiFiLLDP(...)`         | :negative_squared_cross_mark: |

## Connection

| original command                  | library path             | implemented                   |
|-----------------------------------|--------------------------|-------------------------------|
| `nmcli device up`                 | `NMCli.Connection.Up(...)` | :heavy_check_mark:            |

## Agent

| original command  | library path        | implemented                   |
|-------------------|---------------------|-------------------------------|
| `nmcli agent ...` | not implemented yet | :negative_squared_cross_mark: |

## Monitor

| original command    | library path        | implemented                   |
|---------------------|---------------------|-------------------------------|
| `nmcli monitor ...` | not implemented yet | :negative_squared_cross_mark: |

