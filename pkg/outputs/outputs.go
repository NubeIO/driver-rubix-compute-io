package outputs

import (
	"context"
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/pigpiod"
	log "github.com/sirupsen/logrus"
	"reflect"
	"sync"
	"time"
)

type Outputs struct {
	TestMode      bool    `json:"test_mode,omitempty"`
	IONum         string  `json:"io_num,omitempty"`
	Value         float64 `json:"value"`
	valueOriginal float64
	Pin           int
	Type          string
	SupportsPWM   bool
	DeviceIP      string
	DevicePort    int
	mutex         sync.RWMutex
}

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
	UO1: Outputs{IONum: "UO1", Pin: 21, Type: "UO", SupportsPWM: false},
	UO2: Outputs{IONum: "UO2", Pin: 20, Type: "UO", SupportsPWM: false},
	UO3: Outputs{IONum: "UO3", Pin: 19, Type: "UO", SupportsPWM: true},
	UO4: Outputs{IONum: "UO4", Pin: 12, Type: "UO", SupportsPWM: false},
	UO5: Outputs{IONum: "UO5", Pin: 13, Type: "UO", SupportsPWM: false},
	UO6: Outputs{IONum: "UO6", Pin: 18, Type: "UO", SupportsPWM: true},
	DO1: Outputs{IONum: "DO1", Pin: 22, Type: "DO", SupportsPWM: false},
	DO2: Outputs{IONum: "DO2", Pin: 23, Type: "DO", SupportsPWM: false},
}

type Body struct {
	IONum string  `json:"io_num"`
	Value float64 `json:"value"`
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

//SupportsPWM will check if the output supports PWM
func SupportsPWM(strut interface{}, ioNum string) (exist bool, pin int, err error) {
	t := reflect.TypeOf(strut)
	if kind := t.Kind(); kind != reflect.Struct {
		return false, 0, errors.New("expects a struct")
	}
	v := reflect.ValueOf(strut)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
		if v.Field(i).FieldByName("IONum").String() == ioNum {
			exist = v.Field(i).FieldByName("SupportsPWM").Bool()
			pin = int(v.Field(i).FieldByName("Pin").Int())
		}
	}
	return
}

func (inst *Outputs) getIP() (out string) {
	ip := "127.0.0.1"
	port := 8888
	if inst.DeviceIP != "" {
		ip = "127.0.0.1"
	}
	if inst.DevicePort != 0 {
		port = inst.DevicePort
	}
	out = fmt.Sprintf("%s:%d", ip, port)
	return
}

func (inst *Outputs) logWrite(asDo, doValue bool) {
	voltage := inst.valueOriginal / 10
	percentage := "%" + fmt.Sprintf("%f", inst.valueOriginal)
	if asDo {
		if doValue {
			voltage = 12
		} else {
			voltage = 0
		}
		log.Infoln("rubix.io.outputs.write() WRITE-DO io-name:", inst.IONum, "voltage:", voltage, "on-off:", doValue)
	} else {
		log.Infoln("rubix.io.outputs.write() WRITE-AO io-name:", inst.IONum, "voltage:", voltage, "percentage:", percentage)
	}
}

func (inst *Outputs) write() (out Body, err error) {
	inst.mutex.Lock()
	defer inst.mutex.Unlock()
	var val = inst.Value * 1000000
	io := inst.IONum
	out.IONum = inst.IONum
	out.Value = inst.valueOriginal
	exists, err := ioExists(OutputMaps, io)
	if !exists || err != nil {
		return out, errors.New("IO-number is not valid")
	}
	isPWM, pin, err := SupportsPWM(OutputMaps, io)
	if !exists || err != nil {
		return out, errors.New("error is validating if UO is a PWM")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := pigpiod.Connect(ctx, inst.getIP())
	if err != nil {
		return out, err
	}
	defer c.Close()

	if isPWM { //WRITE AOs
		inst.logWrite(false, false)
		err := c.HardwarePWM(pin, int(val))
		if err != nil {
			return out, err
		}

	} else { //WRITE on/off
		if inst.valueOriginal > 0 {
			inst.logWrite(true, true)
			out.Value = 1
			err = c.WriteOn(pin)
			if err != nil {
				return out, err
			}
		} else {
			inst.logWrite(true, false)
			out.Value = 0
			err = c.WriteOff(pin)
			if err != nil {
				return out, err
			}
		}

	}
	return out, nil
}

type Message struct {
	Message string `json:"message"`
}
