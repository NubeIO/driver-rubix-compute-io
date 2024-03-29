module github.com/NubeIO/nubeio-rubix-app-pi-gpio-go

go 1.17

//replace github.com/NubeIO/nubeio-rubix-lib-helpers-go => /home/aidan/code/go/nube/nubeio-rubix-lib-helpers-go
//replace github.com/NubeIO/nubeio-rubix-lib-models-go => /home/aidan/code/go/nube/nubeio-rubix-lib-models-go
//replace github.com/NubeIO/nubeio-rubix-lib-rest-go => /home/aidan/code/go/nube/nubeio-rubix-lib-rest-go

require (
	github.com/NubeDev/flow-eng v0.0.8
	github.com/NubeIO/configor v0.0.3
	github.com/NubeIO/nubeio-rubix-lib-helpers-go v0.2.7
	github.com/d2r2/go-i2c v0.0.0-20191123181816-73a8a799d6bc
	github.com/eclipse/paho.mqtt.golang v1.4.1
	github.com/gin-gonic/gin v1.7.7
	github.com/sirupsen/logrus v1.9.0
	github.com/stianeikeland/go-rpio/v4 v4.6.0
	periph.io/x/conn/v3 v3.6.10
	periph.io/x/host/v3 v3.7.2
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/d2r2/go-logger v0.0.0-20210606094344-60e9d1233e22 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-yaml/yaml v2.1.0+incompatible // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	golang.org/x/net v0.0.0-20220127074510-2fabfed7e28f // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
