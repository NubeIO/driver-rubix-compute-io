package inputs

import (
	"encoding/binary"
	"fmt"
	"github.com/d2r2/go-i2c"
	"log"
)

type inputsMap struct {
	Raw  int     `json:"raw"`
	TEMP float64 `json:"temp"`
	Volt float64 `json:"volt"`
	Amps float64 `json:"amps"`
}

type Data struct {
	UI1 inputsMap `json:"ui_1"`
}

func Inputs2() {

	testBytes := [16]byte{240, 0, 249, 43, 249, 157, 241, 18, 240, 0, 240, 0, 240, 0, 240, 0}
	fmt.Println(testBytes, "testBytes")

	bus, err := i2c.NewI2C(0x33, 1)
	if err != nil {
		log.Fatal(err)
	}

	defer bus.Close()
	bytes, _int, err := bus.ReadRegBytes(0xF, 16)
	if err != nil {
		fmt.Println(bytes)
		return
	}
	inputs := &Data{}
	x := 10.0 / 4096.0
	for i := 0; i < 16; i = i + 2 {
		data := binary.BigEndian.Uint16(bytes[i:i+2]) & 0xFFF
		voltage := float64(data) * x
		if i == 0 {
			inputs = &Data{
				UI1: inputsMap{
					Volt: voltage,
				},
			}
		}
		fmt.Println(data, "data")
		fmt.Println(voltage, "voltage")
		fmt.Println(" ")
	}
	fmt.Println(inputs)
	fmt.Println("_int", _int)
	fmt.Println("bytes", bytes)
}
