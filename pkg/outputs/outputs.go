package outputs

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type Outputs struct {
	TestMode      bool    `json:"test_mode,omitempty"`
	IONum         string  `json:"io_num,omitempty"`
	Value         float64 `json:"value"`
	valueOriginal float64
	Pin           uint32
	Type          string
}

//type OutputMap struct {
//	IONum string
//
//
//}

/*
pin mapping
U01-21
U02-20
U03-19(HW-PWM)   ///////CAN be a PWM
U04-12
U05-13(HW-PWM)
U06-18(HW-PWM)   ///////CAN be a PWM
DO1-22
DO2-23

On the Raspberry Pi, add dtoverlay=pwm-2chan to /boot/config.txt. This defaults to GPIO_18 as the pin for PWM0 and GPIO_19 as the pin for PWM1.
Alternatively, you can change GPIO_18 to GPIO_12 and GPIO_19 to GPIO_13 using dtoverlay=pwm-2chan,pin=12,func=4,pin2=13,func2=4.
Reboot your Raspberry Pi.
You can check everything is working on running lsmod | grep pwm and looking for pwm_bcm2835
*/

var OutputMaps = struct {
	UO1 Outputs
	UO2 Outputs
	UO3 Outputs
	UO4 Outputs
	UO5 Outputs
	UO6 Outputs
	DO1 Outputs
	DO2 Outputs
}{
	UO1: Outputs{IONum: "UO1", Pin: 21, Type: "UO"},
	UO2: Outputs{IONum: "UO2", Pin: 20, Type: "UO"},
	UO3: Outputs{IONum: "UO3", Pin: 19, Type: "UO"},
	UO4: Outputs{IONum: "UO4", Pin: 12, Type: "UO"},
	UO5: Outputs{IONum: "UO5", Pin: 13, Type: "UO"},
	UO6: Outputs{IONum: "UO6", Pin: 18, Type: "UO"},
	DO1: Outputs{IONum: "DO1", Pin: 22, Type: "DO"},
	DO2: Outputs{IONum: "DO2", Pin: 23, Type: "DO"},
}

type Body struct {
	IONum string  `json:"io_num"`
	Value float64 `json:"value"`
	Debug *bool   `json:"debug"`
}

func ioExists(strut interface{}, item string) (exist bool, err error) {
	t := reflect.TypeOf(strut)
	if kind := t.Kind(); kind != reflect.Struct {
		return false, errors.New("expects a struct")
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Name == item {
			exist = true
		}
	}
	return

}

func (inst *Outputs) logWrite() {
	voltage := inst.valueOriginal / 10
	percentage := "%" + fmt.Sprintf("%f", inst.valueOriginal)
	log.Infoln("rubix.io.outputs.write() io-name:", inst.IONum, "voltage:", voltage, "percentage:", percentage)
}

func (inst *Outputs) write() (ok bool, err error) {
	var val = inst.Value * 1000000
	io := inst.IONum
	fmt.Println(io)
	var pin uint32
	if io == "UO3" {
		pin = 19
	} else if io == "UO4" {
		pin = 12

	} else if io == "UO5" {
		pin = 13
	}
	if inst.TestMode {
		inst.logWrite()
	} else {
		inst.logWrite()
		fmt.Println(1111, pin, val)

	}
	return true, nil
}

func (inst *Outputs) Init() error {
	if inst.TestMode {

	} else {

		//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		//defer cancel()
		//
		//c, err := pigpiod.Connect(ctx, "192.168.15.10:8888")
		//if err != nil {
		//	panic(err)
		//}
		//defer c.Close()
		//Client = c
		//read, err := Client.Read(23)
		//fmt.Println(99999, read, err)

		//c, err := pigpio.New("192.168.15.10:8888")
		//fmt.Println("", err)
		//if err != nil {
		//	fmt.Println("ERR set pig client", err)
		//	return err
		//}
		//client = c

	}
	return nil
}

type Message struct {
	Message string `json:"message"`
}
