package cmd

import (
	"github.com/spf13/cobra"
)

var (
	ip         string
	port       int
	point      string
	writeValue float64

	ping   bool //ping the api
	toggle bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "rubix-cli",
	Short: "description",
	Long:  `description`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
	}
}

func init() {

	//ip
	RootCmd.PersistentFlags().StringVarP(&ip, "ip", "", "192.168.15.10", "device ip")
	RootCmd.PersistentFlags().IntVarP(&port, "port", "", 5001, "device port")

	//point
	RootCmd.PersistentFlags().StringVarP(&point, "point", "", "coils", "type coil")

}
