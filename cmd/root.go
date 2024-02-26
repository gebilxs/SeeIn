package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "SeeIn",
		Short: "SeeIn is a UPnP scanner",
		Long: `
   _____                     
  / ____|            |---|  
 | (___   ___  ___    | | _    _
  \___ \ / _ \/ ___\  | || |\ | |   
  ____) |  __/ (______| || | \| |
 |_____/ \___|\______|---|_|  \_|

SeeIn is a command-line application for scanning and discovering UPnP devices
on your local network. Use SeeIn to find UPnP devices, their services, and to 
interact with them.`,
	}
)

func init() {
	rootCmd.AddCommand(fetchCmd)
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
