package upnp

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
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
