package outputs

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
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
U01-21
U02-20
U03-19(HW-PWM)
U04-12
U05-13(HW-PWM)
U06-18(HW-PWM)
DO1-22
DO2-23
*/

var UO1 = rpio.Pin(21)
var UO2 = rpio.Pin(20)
var UO3 = rpio.Pin(19) //PWM
var UO4 = rpio.Pin(12)
var UO5 = rpio.Pin(13) //PWM
var UO6 = rpio.Pin(18) //PWM
var DO1 = rpio.Pin(22)
var DO2 = rpio.Pin(23)

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
		if io == "UO1" || io == "UO2" || io == "UO4" || io == "DO1" || io == "DO2" {
			if val >= 1 {
				pin.Output()
				pin.High()
				log.Infoln("rubix.io.outputs.write() write as BOOL write High io-name:", inst.IONum, "value:", true)
			} else {
				pin.Low()
				log.Infoln("rubix.io.outputs.write() write as BOOL write LOW io-name:", inst.IONum, "value:", false)
			}
		} else {
			inst.logWrite()
			const cycleLength = 100
			fmt.Println(uint32(val), "VALUE------------------")
			pin.DutyCycle(uint32(val), cycleLength)
		}
	}
	return true, nil
}

func (inst *Outputs) HaltPins() error {
	log.Infoln("rubix.io.outputs.HaltPins()")
	if inst.TestMode {
		return nil

	} else {
		defer func() {
			err := rpio.Close()
			if err != nil {
				log.Errorln("rubix.io.outputs.HaltPins() rpio.Close err:", err)
			}
		}()

	}
	return nil
}

func (inst *Outputs) Init() error {
	if inst.TestMode {

	} else {
		if err := rpio.Open(); err != nil {
			log.Fatalf("Error opening GPIO: %s", err.Error())
		}
		defer func() {
			err := rpio.Close()
			if err != nil {
				log.Errorln("rubix.io.outputs.Init() rpio.Close err:", err)
			}
		}()
		//DOs
		//DO1 = rpio.Pin(22)
		//DO1.Output()
		//DO2 = rpio.Pin(23)
		//DO2.Output()

		//PWMs
		const cycleLength = 100
		const pmwClockFrequency = 50 * cycleLength // 50kHz
		UO3 = rpio.Pin(19)
		UO3.Output()
		UO3.Pwm()
		UO3.Freq(pmwClockFrequency)

		UO5 = rpio.Pin(13)
		UO5.Output()
		UO5.Pwm()
		UO5.Freq(pmwClockFrequency)

		UO6 = rpio.Pin(18)
		UO6.Output()
		UO6.Pwm()
		UO6.Freq(pmwClockFrequency)

	}
	return nil
}

type Message struct {
	Message string `json:"message"`
}
