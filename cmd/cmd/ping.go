package cmd

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/nube/rubixio"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "rubix-io points read and write",
	Long:  ``,
	Run:   runPing,
}

func initClient() *rubixio.Client {
	restService := &rest.Service{}
	restService.Url = ip
	restService.Port = port
	restOptions := &rest.Options{}
	restService.Options = restOptions
	restService = rest.New(restService)

	nubeProxy := &rest.NubeProxy{}
	nubeProxy.UseRubixProxy = false
	nubeProxy.RubixUsername = "admin"
	nubeProxy.RubixPassword = "N00BWires"
	nubeProxy.RubixProxyPath = nube.Services.BacnetServer.Proxy
	restService.NubeProxy = nubeProxy
	client := rubixio.New(restService)
	return client
}

func runPing(cmd *cobra.Command, args []string) {
	client := initClient()
	p, res := client.Ping()
	if res.GetError() != nil {
		fmt.Println("ping-err", res.GetError())
		return
	}
	fmt.Println("ping-ok", p.Ok)

}

func init() {
	RootCmd.AddCommand(pingCmd)
	pingCmd.PersistentFlags().BoolVarP(&ping, "ping", "", true, "ping the api")

}
