package main

import (
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
	testMode := false
	output := &outputs.Outputs{
		TestMode: testMode,
	}
	input := &inputs.Inputs{
		TestMode: testMode,
	}
	input.InitBus()
	err := output.Init()
	if err != nil {
		log.Errorln("rubix.io.outputs.main() failed to init outputs")
	}

	router.POST("/api/write", output.Write)
	router.GET("/api/inputs", input.ReadAll)

	server := &http.Server{
		Addr:    ":5001",
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}
	err = output.HaltPins()
	if err != nil {
		log.Errorln("rubix.io.outputs.main() failed to halt all outputs", err)
	}
	log.Println("Server exiting")
}
