package outputs

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
)

type Outputs struct {
	TestMode      bool    `json:"test_mode,omitempty"`
	IONum         string  `json:"io_num,omitempty"`
	Value         float64 `json:"value"`
	valueOriginal float64
}

type OutputMap struct {
	IONum string
	Pin   string
	Type  string
}

/*
pin mapping
U01-21
U02-20
U03-19(HW-PWM)
U04-12
U05-13(HW-PWM)
U06-18(HW-PWM)
DO1-22
DO2-23
*/

var outputsArr = []string{"UO1", "UO2", "UO3", "UO4", "UO5", "UO6", "DO1", "DO1"}

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
var UO4 = gpioreg.ByName(OutputMaps.UO4.Pin)
var UO5 = gpioreg.ByName(OutputMaps.UO5.Pin)
var UO6 = gpioreg.ByName(OutputMaps.UO6.Pin)
var DO1 = gpioreg.ByName(OutputMaps.DO1.Pin)
var DO2 = gpioreg.ByName(OutputMaps.DO2.Pin)

type Body struct {
	IONum string  `json:"io_num"`
	Value float64 `json:"value"`
	Debug *bool   `json:"debug"`
}

func (inst *Outputs) logWrite() {
	voltage := inst.valueOriginal / 10
	percentage := "%" + fmt.Sprintf("%f", inst.valueOriginal)
	log.Infoln("rubix.io.outputs.write() io-name:", inst.IONum, "voltage:", voltage, "percentage:", percentage)
}

func (inst *Outputs) write() (ok bool, err error) {
	var val = 16777216 * inst.Value
	io := inst.IONum
	if inst.TestMode {
		inst.logWrite()
	} else {
		pin := inst.pinSelect()
		if pin == nil {
			return false, errors.New("no valid io num was selected try UO1 or DO1")
		}
		if io == OutputMaps.DO1.IONum || io == OutputMaps.DO2.IONum {
			if val >= 1 {
				log.Infoln("rubix.io.outputs.write() write as BOOL write High io-name:", inst.IONum, "value:", true)
				if err := pin.Out(gpio.High); err != nil {
					log.Fatal(err)
				}
			} else {
				log.Infoln("rubix.io.outputs.write() write as BOOL write LOW io-name:", inst.IONum, "value:", false)
				if err := pin.Out(gpio.Low); err != nil {
					log.Fatal(err)
				}
			}
		} else {

			inst.logWrite()
			//defer func(pin gpio.PinIO) {
			//	err := pin.Halt()
			//	fmt.Println("------PIN HALT")
			//	if err != nil {
			//		fmt.Println("------PIN HALT ERROR", err)
			//	}
			//}(pin)
			if err := pin.PWM(gpio.Duty(val), 8000*physic.Hertz); err != nil {
				log.Errorln(err)
				return false, err
			}
			err := pin.Halt()
			if err != nil {
				fmt.Println("------PIN HALT ERROR", err)
			} else {
				fmt.Println("------PIN HALT")
			}
		}
	}
	return true, nil
}

// HaltPin disable the gpio
func (inst *Outputs) haltPin(pin gpio.PinIO) {
	if inst.TestMode {
	} else {
		log.Infoln("rubix.io.outputs.haltPin() io-name:", pin.Name())
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
			log.Errorln("rubix.io.outputs.HaltPins() halt UO2")
			return err
		}
		err = UO3.Halt()
		if err != nil {
			log.Errorln("rubix.io.outputs.HaltPins() halt UO3")
			return err
		}
		err = UO4.Halt()
		if err != nil {
			log.Errorln("rubix.io.outputs.HaltPins() halt UO4")
			return err
		}
		err = UO5.Halt()
		if err != nil {
			log.Errorln("rubix.io.outputs.HaltPins() halt UO5")
			return err
		}
		err = UO6.Halt()
		if err != nil {
			log.Errorln("rubix.io.outputs.HaltPins() halt UO6")
			return err
		}
		err = DO1.Halt()
		if err != nil {
			log.Errorln("rubix.io.outputs.HaltPins() halt DO1")
			return err
		}
		err = DO2.Halt()
		if err != nil {
			log.Errorln("rubix.io.outputs.HaltPins() halt DO2")
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
		UO3 = gpioreg.ByName(OutputMaps.UO3.Pin)
		UO4 = gpioreg.ByName(OutputMaps.UO4.Pin)
		UO5 = gpioreg.ByName(OutputMaps.UO5.Pin)
		UO6 = gpioreg.ByName(OutputMaps.UO6.Pin)
		DO1 = gpioreg.ByName(OutputMaps.DO1.Pin)
		DO2 = gpioreg.ByName(OutputMaps.DO2.Pin)
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
