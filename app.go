package main

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/config"
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/inputs"
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/outputs"
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/ping"
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/runner"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	conf := config.CreateApp()
	if err := os.MkdirAll(conf.GetAbsConfigDir(), 0755); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(conf.GetAbsDataDir(), 0755); err != nil {
		panic(err)
	}

	router := gin.Default()
	ip := conf.Server.Address
	pings := &ping.Ping{}
	testMod := false
	output := &outputs.Outputs{
		TestMode:   testMod,
		DeviceIP:   ip,
		DevicePort: 8888,
	}

	input := &inputs.Inputs{
		TestMode: testMod,
	}
	err := input.Init()
	if err != nil {
		log.Errorln("rubix.io.outputs.main() failed to init inputs")
	}
	router.GET("/api/system/ping", pings.Ping)
	router.POST("/api/outputs", output.Write)
	router.GET("/api/outputs/:io/:value", output.WriteOne)
	router.POST("/api/outputs/bulk", output.BulkWrite)
	router.GET("/api/outputs/all/:value", output.WriteAll)

	inputLoop := runner.NewRunner(&runner.InputRunner{
		Enable:      true,
		TestMode:    testMod,
		LoopDelayMs: 0,
	})

	if !inputLoop.Enable {
		router.GET("/api/inputs/all", input.ReadAll)
	}

	go inputLoop.Runner()

	port := conf.Server.Port
	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Infoln("rubix.io.main() interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("rubix.io.main() Server Close", err)
		}
	}()
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Infoln("rubix.io.main() Server closed as request")
		} else {
			log.Fatal("rubix.io.main() Server unexpect Close", err)
		}
	}
	log.Infoln("rubix.io.main() Server exiting")
}
