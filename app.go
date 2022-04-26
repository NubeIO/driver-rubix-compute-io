package main

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/config"
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/inputs"
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/outputs"
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
	output := &outputs.Outputs{
		DeviceIP:   ip,
		DevicePort: 8888,
	}
	input := &inputs.Inputs{
		TestMode: conf.Debug,
	}
	err := output.Init()
	if err != nil {
		log.Errorln("rubix.io.outputs.main() failed to init outputs")
	}
	err = input.Init()
	if err != nil {
		log.Errorln("rubix.io.outputs.main() failed to init inputs")
	}

	router.POST("/api/outputs", output.Write)
	router.POST("/api/outputs/bulk", output.BulkWrite)
	router.GET("/api/outputs/all/:value", output.WriteAll)
	router.GET("/api/inputs/all", input.ReadAll)

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
