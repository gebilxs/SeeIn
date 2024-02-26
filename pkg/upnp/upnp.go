package upnp

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type UPnPDevice struct {
	// Define the structure according to the XML response structure
	DeviceType       string `xml:"device>deviceType"`
	FriendlyName     string `xml:"device>friendlyName"`
	ModelName        string `xml:"device>modelName"`
	ModelNumber      string `xml:"device>modelNumber"`
	ModelDescription string `xml:"device>modelDescription"`
	ModelURL         string `xml:"device>modelURL"`
	Manufacturer     string `xml:"device>manufacturer"`
	ManufacturerURL  string `xml:"device>manufacturerURL"`
	SerialNumber     string `xml:"device>serialNumber"`
	UDN              string `xml:"device>UDN"`
	PresentationURL  string `xml:"device>presentationURL"`
	ServicesType     string `xml:"device>serviceList>service>serviceType"`
	// Add more fields as needed
}
type UrlType struct {
	IPAddress string
	Port      string
	Path      string
}

// FetchAndParseUPnP fetches UPnP device details from the specified URL and parses the XML.
func FetchAndParseUPnP(url string) (*UPnPDevice, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching URL: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var device UPnPDevice
	if err := xml.Unmarshal(body, &device); err != nil {
		return nil, fmt.Errorf("error parsing XML: %w", err)
	}

	return &device, nil
}

// FetchAndParseUPnP fetches UPnp device details with ssdp request from the specified URL and parses the XML
func FetchAndParseUPnPBySsdp(url string) (*UPnPDevice, error) {
	pattern := regexp.MustCompile(`https?://(\d+\.\d+\.\d+\.\d+):(\d+)(/[^?]*)?`)
	matches := pattern.FindStringSubmatch(url)
	var urlT UrlType
	if len(matches) >= 3 {
		urlT.IPAddress = matches[1]
		urlT.Port = matches[2]
	}
	if len(matches) > 3 && matches[3] != "" {
		urlT.Path = matches[3]
	}
	//fmt.Println("urlT-Ipaddress is ", urlT.IPAddress)
	//fmt.Println("urlT-Prot is ", urlT.Port)
	//fmt.Println("urlT-Path is ", urlT.Path)
	ssdpAddress := urlT.IPAddress + ":" + urlT.Port
	localAddress := "0.0.0.0:0"
	// 创建udp链接
	location, err := BuildSsdpUdp(ssdpAddress, localAddress)
	if err != nil {
		fmt.Println("create udp fail", err)
		os.Exit(1)
	}

	//print(location)
	locationArray := pattern.FindStringSubmatch(location)
	routerXml := locationArray[3]
	portXml := locationArray[2]
	// Fetch the device details from the location URL
	urlxml := "http://" + urlT.IPAddress + ":" + portXml + routerXml
	//print(urlxml)
	return FetchAndParseUPnP(urlxml)
}

func BuildSsdpUdp(ssdpAddress string, localAddress string) (string, error) {
	localAddr, err := net.ResolveUDPAddr("udp4", localAddress)
	if err != nil {
		return "", fmt.Errorf("Error resolving local address: %w", err)
	}
	ssdpAddr, err := net.ResolveUDPAddr("udp4", ssdpAddress)
	if err != nil {
		return "", fmt.Errorf("Error resolving SSDP address: %w", err)
	}

	// 创建UDP连接
	conn, err := net.DialUDP("udp4", localAddr, ssdpAddr)
	if err != nil {
		return "", fmt.Errorf("Error dialing UDP: %w", err)
	}
	defer conn.Close()

	// 设置超时
	conn.SetDeadline(time.Now().Add(3 * time.Second))

	// SSDP发现消息
	ssdpDiscover := "M-SEARCH * HTTP/1.1\r\n" +
		"HOST: 239.255.255.250:1900\r\n" +
		"MAN: \"ssdp:discover\"\r\n" +
		"MX: 3\r\n" +
		"ST: ssdp:all\r\n" + // 查找所有设备
		"\r\n"

	// 发送SSDP发现消息
	_, err = conn.Write([]byte(ssdpDiscover))
	if err != nil {
		return "", fmt.Errorf("Error sending SSDP discover message: %w", err)
	}

	// 读取响应
	buf := make([]byte, 2048)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		return "", fmt.Errorf("Error reading from UDP: %w", err)
	}

	// 解析响应中的LOCATION字段
	response := string(buf[:n])
	fmt.Println(response)
	for _, line := range strings.Split(response, "\r\n") {
		if strings.HasPrefix(strings.ToLower(line), "location:") {
			location := strings.TrimSpace(line[len("location:"):])
			return location, nil
		}
	}

	return "", errors.New("LOCATION field not found in SSDP response")
}
func PrintResult(device *UPnPDevice) error {
	fmt.Println("UPnP Device:")
	fmt.Printf("  Device Type: %s\n", device.DeviceType)
	fmt.Printf("  Friendly Name: %s\n", device.FriendlyName)
	fmt.Printf("  Model Name: %s\n", device.ModelName)
	fmt.Printf("  Model Number: %s\n", device.ModelNumber)
	fmt.Printf("  Model Description: %s\n", device.ModelDescription)
	fmt.Printf("  Model URL: %s\n", device.ModelURL)
	fmt.Printf("  Manufacturer: %s\n", device.Manufacturer)
	fmt.Printf("  Manufacturer URL: %s\n", device.ManufacturerURL)
	fmt.Printf("  Serial Number: %s\n", device.SerialNumber)
	fmt.Printf("  UDN: %s\n", device.UDN)
	fmt.Printf("  Presentation URL: %s\n", device.PresentationURL)
	fmt.Printf("  Services Type: %s\n", device.ServicesType)
	return nil
}
