package cmd

import (
	"fmt"
	pprint "github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/helpers/print"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "rubix-io points read",
	Long:  ``,
	Run:   runRead,
}

func runRead(cmd *cobra.Command, args []string) {
	client := initClient()
	points, res := client.GetInputs()
	if res.GetError() != nil {
		fmt.Println("ping-err", res.GetError())
		return
	}
	pprint.PrintJOSN(points)

}

func init() {
	RootCmd.AddCommand(readCmd)

}
