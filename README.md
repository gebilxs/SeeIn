# SeeIn 公网UPnP设备发现工具
```shell
  _____
  / ____|            |---|
 | (___   ___  ___    | | _    _    
  \___ \ / _ \/ ___\  | || |\ | |   
  ____) |  __/ (______| || | \| |   
 |_____/ \___|\______|---|_|  \_|   

SeeIn is a command-line application for scanning and discovering UPnP devices
on your local network. Use SeeIn to find UPnP devices, their services, and to
interact with them.

Usage:
  SeeIn [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  fetch       Fetches UPnP device information from a given URL
  help        Help about any command

Flags:
  -h, --help   help for SeeIn

Use "SeeIn [command] --help" for more information about a command.

```
```shell
go run main.go fetch --url http://186.134.69.46/
```

## Example Usages
### UPnP
在UPnP设备上执行发现操作
对于公网暴露的UPnP设备，可以通过以下命令发现设备
```shell
SeeIn>go run main.go fetch -u http://121.152.159.108:5531/rss/Starter_desc.xml -s=false  
UPnP Device:
  Device Type: urn:d-link-com:D-Link-Starter:1
  Friendly Name: Wireless AC Dual Band Router
  Model Name: D-Link DIR-806A
  Model Number: Wireless AC Dual Band Router
  Model Description: ADSL
  Model URL: http://www.dlink.com
  Manufacturer: D-Link Corporation
  Manufacturer URL: Http://www.dlink.com/
  Serial Number:
  UDN: uuid:00000000-0008-2fe4-0400-4000009b4001
  Presentation URL: http://121.152.159.108/
  Services Type: urn:d-link-com:service:D-Link-Starter:1

```
对于不直接暴露的UPnP设备，可以通过以下命令发现设备（通过SSDP发送发现设备请求，获得Location）
```shell
SeeIn>go run main.go fetch -u http://128.65.194.253:1900 
HTTP/1.1 200 OK
CACHE-CONTROL: max-age=1900
ST: upnp:rootdevice
USN: uuid:73796E6F-6473-6D00-0000-001132b0b142::upnp:rootdevice
EXT:
SERVER: Synology/DSM/128.65.194.253
LOCATION: http://128.65.194.253:4000/ssdp/desc-DSM-eth0.xml
OPT: "http://schemas.upnp.org/upnp/1/0/"; ns=01
01-NLS: 1
BOOTID.UPNP.ORG: 1
CONFIGID.UPNP.ORG: 1337

UPnP Device:
  Device Type: urn:schemas-upnp-org:device:Basic:1
  Friendly Name: Monolithe (RS819)
  Model Name: RS819
  Model Number: RS819 6.2-25556
  Model Description: Synology NAS
  Model URL: http://www.synology.com
  Manufacturer: Synology
  Manufacturer URL: http://www.synology.com
  Serial Number: 001132b0b142
  UDN: uuid:73796E6F-6473-6D00-0000-001132b0b142
  Presentation URL: http://128.65.194.253:4000/
  Services Type: urn:schemas-dummy-com:service:Dummy:1
```
打包成二进制文件可以直接执行
