package inputs

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/thermistor"
	"github.com/d2r2/go-i2c"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Inputs struct {
	TestMode   bool
	DeviceIP   string
	DevicePort int
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

func (inst *Inputs) getIP() (out string) {
	ip := "0.0.0.0"
	port := 8888
	if inst.DeviceIP != "" {
		ip = inst.DeviceIP
	}
	if inst.DevicePort != 0 {
		port = inst.DevicePort
	}
	out = fmt.Sprintf("%s:%d", ip, port)
	return
}

func (inst *Inputs) ReadAll(ctx *gin.Context) {
	testBytes := []byte{248, 182, 248, 176, 248, 168, 248, 184, 248, 174, 248, 178, 248, 177, 248, 177}
	if inst.TestMode {
		data := inst.DecodeData(testBytes)
		reposeHandler(data, nil, ctx)
	} else {
		i2, err := i2c.NewI2C(0x33, 1)
		if err != nil {
			log.Errorln("pig-io.inputs.failed.ReadAll() failed to open i2c")
			reposeHandler(nil, errors.New("failed to open i2c"), ctx)
			return
		}
		defer i2.Close()
		err = i2.WriteRegU8(0x33, 0xDA)
		if err != nil {
			log.Errorln("pig-io.inputs.failed.ReadAll() failed to write i2c")
			reposeHandler(nil, errors.New("failed write to inputs board"), ctx)
			return
		}

		bytes, _, err := i2.ReadRegBytes(0xF, 16)
		fmt.Println(bytes, err)
		if err != nil {
			log.Errorln("pig-io.inputs.failed.ReadAll() failed to read i2c")
			reposeHandler(nil, errors.New("failed to read i2c"), ctx)
			return
		}
		data := inst.DecodeData(bytes)
		reposeHandler(data, nil, ctx)
		//ip := inst.getIP()
		//ct, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		//defer cancel()
		//
		//c, err := pigpiod.Connect(ct, ip)
		//if err != nil {
		//	panic(err)
		//}
		//defer c.Close()
		//
		//busInst, err := c.InitI2c(1, 0x33)
		//if err != nil {
		//	log.Errorln("pig-io.inputs.failed.ReadAll() failed to open i2c")
		//	reposeHandler(nil, errors.New("failed to open i2c"), ctx)
		//	return
		//}
		//
		////_, err = c.WriteI2c(int(busInst), 0x33, 0xDA)
		////if err != nil {
		////	log.Errorln("pig-io.inputs.failed.ReadAll() WriteI2c")
		////	reposeHandler(nil, errors.New("failed to write I2c"), ctx)
		////	return
		////}
		//
		//d, err := c.ReadI2c(int(busInst), 0xF, 16)
		//fmt.Println(d)
		//if err != nil {
		//	log.Errorln("pig-io.inputs.failed.ReadAll() failed to do an i2c read")
		//	reposeHandler(nil, errors.New("failed to read i2c"), ctx)
		//	return
		//}
		//if len(d) == 16 {
		//	err := c.CloseI2c(int(busInst))
		//	if err != nil {
		//		log.Errorln("pig-io.inputs.failed.ReadAll() to close ic2 bus")
		//		reposeHandler(nil, errors.New("failed to close i2c"), ctx)
		//		return
		//	}
		//	data := inst.DecodeData(d)
		//	reposeHandler(data, nil, ctx)
		//	return
		//} else {
		//	err := c.CloseI2c(int(busInst))
		//	if err != nil {
		//		log.Errorln("pig-io.inputs.failed.ReadAll() to close ic2 bus")
		//		reposeHandler(nil, errors.New("failed to close i2c"), ctx)
		//		return
		//	}
		//	if err != nil {
		//		log.Errorln("pig-io.inputs.failed.ReadAll() i2c bytes is incorrect len leg is:", len(d))
		//		reposeHandler(nil, errors.New("incorrect data return"), ctx)
		//		return
		//	}
		//}

		//example of actual hardware read
		//github.com/reef-pi/rpi/i2c
		//bus, err := i2c.New()
		//fmt.Println(bus, err)
		//if err != nil {
		//	return
		//}
		//b := make([]byte, 16)
		//err = bus.ReadFromReg(0x33, 0xDA, b)
		//if err != nil {
		//	return
		//}

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

func (inst *Inputs) DecodeData(bytes []byte) *Data {
	inputs := &Data{}
	for i := 0; i < 16; i = i + 2 {
		data := binary.BigEndian.Uint16(bytes[i:i+2]) & 0xFFF
		if i == 0 {
			voltage := getVoltage(data)
			decodeBool := getBool(data)
			decodeTemp := getTemp(data)
			decodeAmps := getAmps(data)
			inputs.UI1.Raw = data
			inputs.UI1.Temp = decodeTemp
			inputs.UI1.Volt = voltage
			inputs.UI1.Amps = decodeAmps
			inputs.UI1.Bool = decodeBool

		}
		if i == 2 {
			voltage := getVoltage(data)
			decodeBool := getBool(data)
			decodeTemp := getTemp(data)
			decodeAmps := getAmps(data)
			inputs.UI2.Raw = data
			inputs.UI2.Temp = decodeTemp
			inputs.UI2.Volt = voltage
			inputs.UI2.Amps = decodeAmps
			inputs.UI2.Bool = decodeBool

		}
		if i == 4 {
			voltage := getVoltage(data)
			decodeBool := getBool(data)
			decodeTemp := getTemp(data)
			decodeAmps := getAmps(data)
			inputs.UI3.Raw = data
			inputs.UI3.Temp = decodeTemp
			inputs.UI3.Volt = voltage
			inputs.UI3.Amps = decodeAmps
			inputs.UI3.Bool = decodeBool
		}
		if i == 6 {
			voltage := getVoltage(data)
			decodeBool := getBool(data)
			decodeTemp := getTemp(data)
			decodeAmps := getAmps(data)
			inputs.UI4.Raw = data
			inputs.UI4.Temp = decodeTemp
			inputs.UI4.Volt = voltage
			inputs.UI4.Amps = decodeAmps
			inputs.UI4.Bool = decodeBool
		}
		if i == 8 {
			voltage := getVoltage(data)
			decodeBool := getBool(data)
			decodeTemp := getTemp(data)
			decodeAmps := getAmps(data)
			inputs.UI5.Raw = data
			inputs.UI5.Temp = decodeTemp
			inputs.UI5.Volt = voltage
			inputs.UI5.Amps = decodeAmps
			inputs.UI5.Bool = decodeBool
		}
		if i == 10 {
			voltage := getVoltage(data)
			decodeBool := getBool(data)
			decodeTemp := getTemp(data)
			decodeAmps := getAmps(data)
			inputs.UI6.Raw = data
			inputs.UI6.Temp = decodeTemp
			inputs.UI6.Volt = voltage
			inputs.UI6.Amps = decodeAmps
			inputs.UI6.Bool = decodeBool
		}
		if i == 12 {
			voltage := getVoltage(data)
			decodeBool := getBool(data)
			decodeTemp := getTemp(data)
			decodeAmps := getAmps(data)
			inputs.UI7.Raw = data
			inputs.UI7.Temp = decodeTemp
			inputs.UI7.Volt = voltage
			inputs.UI7.Amps = decodeAmps
			inputs.UI7.Bool = decodeBool
		}
		if i == 14 {
			voltage := getVoltage(data)
			decodeBool := getBool(data)
			decodeTemp := getTemp(data)
			decodeAmps := getAmps(data)
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
