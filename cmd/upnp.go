package cmd

import (
	"SeeIn/pkg/error"
	"SeeIn/pkg/upnp"
	"fmt"
	"github.com/spf13/cobra"
)

var url string
var ssdp bool

// fetchCmd 命令用于从指定的 URL 获取 UPnP 设备信息
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches UPnP device information from a given URL",
	Long: `Fetch command retrieves UPnP device details from the specified URL
and displays structured information about the device.`,
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			error.Cry(fmt.Errorf("URL is required"))
			return
		}

		if !ssdp {
			device, err := upnp.FetchAndParseUPnP(url)
			if err != nil {
				error.Cry(err)
				return
			}
			err = upnp.PrintResult(device)
			if err != nil {
				error.Cry(err)
				return
			}
		} else {
			print("1")
			device, err := upnp.FetchAndParseUPnPBySsdp(url)
			if err != nil {
				error.Cry(err)
			}
			err = upnp.PrintResult(device)
			if err != nil {
				error.Cry(err)
				return
			}
		}

	},
}

// 初始化函数用于设置 fetchCmd 命令的 flags
func init() {
	fetchCmd.Flags().StringVarP(&url, "url", "u", "", "URL of the UPnP device to fetch (required)")
	fetchCmd.Flags().BoolVarP(&ssdp, "ssdp", "s", true, "SSDP request for UPnP device to fetch,default is true")
	fetchCmd.MarkFlagRequired("url")

}
