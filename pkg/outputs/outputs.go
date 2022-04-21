package outputs

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
)

type Outputs struct {
	TestMode    bool
	IONum       string
	Value       float64
	WriteAsBool bool
	UO1         gpio.PinIO
}

type OutputMap struct {
	IONum string
	Pin   string
	Type  string
}

/*
pin mapping
U1, U2, U3, U4, U5, U6, D1, D2
[21, 20, 19, 12, 13, 18, 22, 23]
*/

var OutputMaps = struct {
	UO1 OutputMap
	UO2 OutputMap
	UO3 OutputMap
	UO4 OutputMap
	UO5 OutputMap
	UO6 OutputMap
	DO1 OutputMap
	DO2 OutputMap
}{
	UO1: OutputMap{IONum: "UO1", Pin: "21", Type: "UO"},
	UO2: OutputMap{IONum: "UO2", Pin: "20", Type: "UO"},
	UO3: OutputMap{IONum: "UO3", Pin: "19", Type: "UO"},
	UO4: OutputMap{IONum: "UO4", Pin: "12", Type: "UO"},
	UO5: OutputMap{IONum: "UO5", Pin: "13", Type: "UO"},
	UO6: OutputMap{IONum: "UO6", Pin: "18", Type: "UO"},
	DO1: OutputMap{IONum: "DO1", Pin: "22", Type: "DO"},
	DO2: OutputMap{IONum: "DO2", Pin: "23", Type: "DO"},
}

var UO1 = gpioreg.ByName(OutputMaps.UO1.Pin)
var UO2 = gpioreg.ByName(OutputMaps.UO2.Pin)
var UO3 = gpioreg.ByName(OutputMaps.UO3.Pin)

type Body struct {
	IONum       string  `json:"io_num"`
	Value       float64 `json:"value"`
	WriteAsBool *bool   `json:"write_as_bool"`
	Debug       *bool   `json:"debug"`
}

func getBody(ctx *gin.Context) (dto *Body, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Outputs) pinSelect() gpio.PinIO {
	io := inst.IONum
	if io == OutputMaps.UO1.IONum {
		return UO1
	} else if io == OutputMaps.UO2.IONum {
		return UO2
	}
	return nil
}

func (inst *Outputs) Write(ctx *gin.Context) {
	body, err := getBody(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	inst.IONum = body.IONum
	inst.Value = body.Value
	if nils.BoolIsNil(body.Debug) {
		inst.TestMode = true
	}
	if nils.BoolIsNil(body.WriteAsBool) {
		inst.WriteAsBool = true
	}
	ok, err := inst.write()
	reposeHandler(ok, err, ctx)
}

func (inst *Outputs) write() (ok bool, err error) {
	var val = 16777216 * inst.Value
	if inst.TestMode {
		log.Infoln("rubix.io.outputs.write() in-debug io-name:", inst.IONum, "value:", val)
	} else {
		pin := inst.pinSelect()
		if inst.WriteAsBool {
			log.Infoln("rubix.io.outputs.write() write as BOOL io-name:", inst.IONum, "value:", val)
		} else {
			log.Infoln("rubix.io.outputs.write() write as PWM io-name:", inst.IONum, "value:", val)
			if err := pin.PWM(gpio.Duty(val), 10000*physic.Hertz); err != nil {
				log.Errorln(err)
				return false, err
			}
		}

	}
	return true, nil
}

// HaltPin disable the gpio
func (inst *Outputs) haltPin(pin gpio.PinIO) {
	if inst.TestMode {
		log.Infoln("rubix.io.outputs.haltPin() in-debug io-name:", inst.IONum)
	} else {
		log.Infoln("rubix.io.outputs.haltPin() io-name:", inst.IONum)
		if err := pin.Halt(); err != nil {
			log.Errorln(err)
		}
	}
}

func (inst *Outputs) HaltPins() error {
	log.Infoln("rubix.io.outputs.HaltPins()")
	if inst.TestMode {
		return nil

	} else {
		err := UO1.Halt()
		if err != nil {
			log.Errorln("rubix.io.outputs.HaltPins() halt UO1")
			return err
		}
		err = UO2.Halt()
		if err != nil {
			return err
		}
		err = UO3.Halt()
		if err != nil {
			return err
		}

	}
	return nil
}

func (inst *Outputs) Init() error {
	if inst.TestMode {

	} else {
		if _, err := host.Init(); err != nil {
			log.Errorln(err)
			return err
		}
		UO1 = gpioreg.ByName(OutputMaps.UO1.Pin)
		UO2 = gpioreg.ByName(OutputMaps.UO2.Pin)
		if UO1 == nil {
			log.Errorln("rubix.io.outputs.Init() failed to init UO1")
			return errors.New("failed to init pin")
		}
	}
	return nil
}

type Message struct {
	Message string `json:"message"`
}

func reposeHandler(body interface{}, err error, ctx *gin.Context) {
	if err != nil {
		if err == nil {
			ctx.JSON(404, Message{Message: "unknown error"})
		} else {
			if body != nil {
				ctx.JSON(404, body)
			} else {
				ctx.JSON(404, Message{Message: err.Error()})
			}

		}
	} else {
		ctx.JSON(200, body)
	}
}
