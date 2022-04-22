package inputs

import (
	"encoding/binary"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/thermistor"
	"github.com/d2r2/go-i2c"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Inputs struct {
	TestMode bool
}

type inputsMap struct {
	Raw  uint16  `json:"raw"`
	Temp float64 `json:"temp_10_k"`
	Volt float64 `json:"volt"`
	Amps float64 `json:"amps"`
	Bool float64 `json:"bool"`
}

type Data struct {
	UI1 inputsMap `json:"UI1"`
	UI2 inputsMap `json:"UI2"`
	UI3 inputsMap `json:"UI3"`
	UI4 inputsMap `json:"UI4"`
	UI5 inputsMap `json:"UI5"`
	UI6 inputsMap `json:"UI6"`
	UI7 inputsMap `json:"UI7"`
	UI8 inputsMap `json:"UI8"`
}

func (inst *Inputs) ReadAll(ctx *gin.Context) {
	testBytes := []byte{240, 0, 249, 43, 249, 157, 241, 18, 240, 0, 240, 0, 240, 0, 240, 0}
	if inst.TestMode {
		data := inst.decodeData(testBytes)
		reposeHandler(data, nil, ctx)
	} else {
		bus, err := i2c.NewI2C(0x33, 1)
		if err != nil {
			log.Errorln(err, "NewI2C")
			reposeHandler(nil, err, ctx)
			return
		}
		defer bus.Close()
		bytes, _, err := bus.ReadRegBytes(0xDA, 16)
		if err != nil {
			log.Errorln(err, "ReadRegBytes")
			reposeHandler(nil, err, ctx)
			return
		}
		data := inst.decodeData(bytes)
		reposeHandler(data, nil, ctx)
	}
}

func getResistance(data uint16) (resistance float64) {
	vin := 2.048
	out := float64(data) * (vin / 4096.0)
	if !(vin-out == 0) {
		r1 := 10000.0
		r2 := (out * r1) / (vin - out)
		resistance = r2
	}
	return
}

func getVoltage(data uint16) (voltage float64) {
	x := 10.0 / 4096.0
	voltage = float64(data) * x
	return
}

func getBool(data uint16) (out float64) {
	if data > 500 {
		out = 0
	} else {
		out = 1
	}
	return
}

func getAmps(data uint16) (voltage float64) {
	x := 10.0 / 4096.0
	voltage = float64(data) * x
	return
}

func getTemp(data uint16) (temp float64) {
	res := getResistance(data)
	temp, err := thermistor.ResistanceToTemperature(res, thermistor.T210K)
	if err != nil {
		return -5555555
	}
	return
}

func (inst *Inputs) decodeData(bytes []byte) *Data {
	inputs := &Data{}

	for i := 0; i < 16; i = i + 2 {
		data := binary.BigEndian.Uint16(bytes[i:i+2]) & 0xFFF
		voltage := getVoltage(data)
		decodeBool := getBool(data)
		decodeTemp := getTemp(data)
		decodeAmps := getAmps(data)
		if i == 0 {
			inputs.UI1.Raw = data
			inputs.UI1.Temp = decodeTemp
			inputs.UI1.Volt = voltage
			inputs.UI1.Amps = decodeAmps
			inputs.UI1.Bool = decodeBool

		}
		if i == 2 {
			inputs.UI2.Raw = data
			inputs.UI2.Temp = decodeTemp
			inputs.UI2.Volt = voltage
			inputs.UI2.Amps = decodeAmps
			inputs.UI2.Bool = decodeBool

		}
		if i == 4 {
			inputs.UI3.Raw = data
			inputs.UI3.Temp = decodeTemp
			inputs.UI3.Volt = voltage
			inputs.UI3.Amps = decodeAmps
			inputs.UI3.Bool = decodeBool
		}
		if i == 6 {
			inputs.UI4.Raw = data
			inputs.UI4.Temp = decodeTemp
			inputs.UI4.Volt = voltage
			inputs.UI4.Amps = decodeAmps
			inputs.UI4.Bool = decodeBool
		}
		if i == 8 {
			inputs.UI5.Raw = data
			inputs.UI5.Temp = decodeTemp
			inputs.UI5.Volt = voltage
			inputs.UI5.Amps = decodeAmps
			inputs.UI5.Bool = decodeBool
		}
		if i == 10 {
			inputs.UI6.Raw = data
			inputs.UI6.Temp = decodeTemp
			inputs.UI6.Volt = voltage
			inputs.UI6.Amps = decodeAmps
			inputs.UI6.Bool = decodeBool
		}
		if i == 12 {
			inputs.UI7.Raw = data
			inputs.UI7.Temp = decodeTemp
			inputs.UI7.Volt = voltage
			inputs.UI7.Amps = decodeAmps
			inputs.UI7.Bool = decodeBool
		}
		if i == 14 {
			inputs.UI8.Raw = data
			inputs.UI8.Temp = decodeTemp
			inputs.UI8.Volt = voltage
			inputs.UI8.Amps = decodeAmps
			inputs.UI8.Bool = decodeBool
		}

	}
	return inputs
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
