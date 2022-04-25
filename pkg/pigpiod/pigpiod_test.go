package pigpiod

import (
	"context"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/inputs"
	"testing"
	"time"
)

func TestCommands(*testing.T) {

	piaddr := "192.168.15.10:8888"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := Connect(ctx, piaddr)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	d, err := c.ReadI2c(1, 16)
	ins := &inputs.Inputs{}
	data := ins.DecodeData(d)
	fmt.Println(data.UI1.Temp)

	fmt.Println(d, err)

}
