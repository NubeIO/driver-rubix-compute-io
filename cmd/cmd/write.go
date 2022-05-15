package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "rubix-io point write",
	Long:  ``,
	Run:   runWrite,
}

func runWrite(cmd *cobra.Command, args []string) {
	client := initClient()
	pnt, res := client.UpdatePointValue(point, int(writeValue))

	if res.GetError() != nil {
		fmt.Println("write-err", res.GetError())
		return
	}

	fmt.Println("WRITE", pnt.Value)

	//if ping {
	//	ping, res := client.Ping()
	//	if res.GetError() != nil {
	//		fmt.Println("ping-err", res.GetError())
	//		return
	//	}
	//	fmt.Println("ping-ok", ping.Ok)
	//}

}

func init() {
	RootCmd.AddCommand(writeCmd)
	writeCmd.PersistentFlags().BoolVarP(&toggle, "toggle", "", false, "write value on/off or o 50 for 5 seconds")
	writeCmd.PersistentFlags().Float64VarP(&writeValue, "value", "", 0, "write value")
}
