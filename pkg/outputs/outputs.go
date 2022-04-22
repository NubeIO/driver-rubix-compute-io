package outputs

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	gpio "github.com/stianeikeland/go-rpio/v4"
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

/*
pin mapping
U1,  U2, U3,         U4, U5,         U6,         D1, D2
[21, 20, 19(HW-PWM), 12, 13(HW-PWM), 18(HW-PWM), 22, 23]
*/
var UO3 = gpio.Pin(19)
var UO5 = gpio.Pin(13)

//var UO1 = gpioreg.ByName(OutputMaps.UO1.Pin)
//var UO2 = gpioreg.ByName(OutputMaps.UO2.Pin)
//var UO3 = gpioreg.ByName(OutputMaps.UO3.Pin)
//var UO4 = gpioreg.ByName(OutputMaps.UO4.Pin)
//var UO5 = gpioreg.ByName(OutputMaps.UO5.Pin)
//var UO6 = gpioreg.ByName(OutputMaps.UO6.Pin)
//var DO1 = gpioreg.ByName(OutputMaps.DO1.Pin)
//var DO2 = gpioreg.ByName(OutputMaps.DO2.Pin)

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
	var val = inst.Value
	io := inst.IONum
	if inst.TestMode {
		inst.logWrite()
	} else {
		pin := inst.pinSelect()
		if io == OutputMaps.DO1.IONum || io == OutputMaps.DO2.IONum {
			if val >= 1 {
				//log.Infoln("rubix.io.outputs.write() write as BOOL write High io-name:", inst.IONum, "value:", true)
				//if err := pin.Out(gpio.High); err != nil {
				//	log.Fatal(err)
				//}
			} else {
				//log.Infoln("rubix.io.outputs.write() write as BOOL write LOW io-name:", inst.IONum, "value:", false)
				//if err := pin.Out(gpio.Low); err != nil {
				//	log.Fatal(err)
				//}
			}
		} else {
			inst.logWrite()
			const cycleLength = 100
			const pmwClockFrequency = 50 * cycleLength // 50kHz
			log.Println("10%")
			pin.Output()
			pin.Pwm()
			pin.Freq(pmwClockFrequency)
			fmt.Println(uint32(val), "VALUE------------------")
			pin.DutyCycle(uint32(val), cycleLength)
			//time.Sleep(3 * time.Second)

			//if err := pin.PWM(gpio.Duty(val), 8000*physic.Hertz); err != nil {
			//	log.Errorln(err)
			//	return false, err
			//}
		}
	}
	return true, nil
}

// HaltPin disable the gpio
//func (inst *Outputs) haltPin(pin gpio.PinIO) {
//	if inst.TestMode {
//	} else {
//		log.Infoln("rubix.io.outputs.haltPin() io-name:", pin.Name())
//		if err := pin.Halt(); err != nil {
//			log.Errorln(err)
//		}
//	}
//}

func (inst *Outputs) HaltPins() error {
	log.Infoln("rubix.io.outputs.HaltPins()")
	if inst.TestMode {
		return nil

	} else {
		//err := UO1.Halt()
		//if err != nil {
		//	log.Errorln("rubix.io.outputs.HaltPins() halt UO1")
		//	return err
		//}
		//err = UO2.Halt()
		//if err != nil {
		//	log.Errorln("rubix.io.outputs.HaltPins() halt UO2")
		//	return err
		//}
		//err = UO3.Halt()
		//if err != nil {
		//	log.Errorln("rubix.io.outputs.HaltPins() halt UO3")
		//	return err
		//}
		//err = UO4.Halt()
		//if err != nil {
		//	log.Errorln("rubix.io.outputs.HaltPins() halt UO4")
		//	return err
		//}
		//err = UO5.Halt()
		//if err != nil {
		//	log.Errorln("rubix.io.outputs.HaltPins() halt UO5")
		//	return err
		//}
		//err = UO6.Halt()
		//if err != nil {
		//	log.Errorln("rubix.io.outputs.HaltPins() halt UO6")
		//	return err
		//}
		//err = DO1.Halt()
		//if err != nil {
		//	log.Errorln("rubix.io.outputs.HaltPins() halt DO1")
		//	return err
		//}
		//err = DO2.Halt()
		//if err != nil {
		//	log.Errorln("rubix.io.outputs.HaltPins() halt DO2")
		//	return err
		//}
	}
	return nil
}

func (inst *Outputs) Init() error {
	if inst.TestMode {

	} else {
		if err := gpio.Open(); err != nil {
			log.Fatalf("Error opening GPIO: %s", err.Error())
		}
		//if _, err := host.Init(); err != nil {
		//	log.Errorln(err)
		//	return err
		//}
		UO3 = gpio.Pin(19)
		UO5 = gpio.Pin(13)
		//UO1 = gpioreg.ByName(OutputMaps.UO1.Pin)
		//UO2 = gpioreg.ByName(OutputMaps.UO2.Pin)
		//UO3 = gpioreg.ByName(OutputMaps.UO3.Pin)
		//UO4 = gpioreg.ByName(OutputMaps.UO4.Pin)
		//UO5 = gpioreg.ByName(OutputMaps.UO5.Pin)
		//UO6 = gpioreg.ByName(OutputMaps.UO6.Pin)
		//DO1 = gpioreg.ByName(OutputMaps.DO1.Pin)
		//DO2 = gpioreg.ByName(OutputMaps.DO2.Pin)
		//if UO1 == nil {
		//	log.Errorln("rubix.io.outputs.Init() failed to init UO1")
		//	return errors.New("failed to init pin")
		//}
	}
	return nil
}

type Message struct {
	Message string `json:"message"`
}
