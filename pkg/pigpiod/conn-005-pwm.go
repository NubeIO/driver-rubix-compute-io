package pigpiod

import (
	"fmt"
)

//pigs hp 18 80000 300000 (will write 3vdc)
//HP g f dc        Set hardware PWM frequency and duty-cycle

func (c *Conn) HardwarePWM(gpio int, dutyCycle int) error {
	cmd := Cmd{
		cmd:       86,
		p1:        uint32(gpio),
		p2:        80000,
		p3:        4,
		extension: uint32(dutyCycle),
	}
	_, err := cmd.ExecuteRes(c.Tcp)
	if err != nil {
		return err
	}
	return nil
}

func (c *Conn) PWMRaw(gpio int, dutyCycle uint32) error {
	cmd := Cmd{
		cmd: 5,
		p1:  uint32(gpio),
		p2:  dutyCycle,
	}
	_, err := cmd.ExecuteRes(c.Tcp)
	if err != nil {
		return err
	}
	return nil
}

func (c *Conn) PWM(gpio int, dutyCycle int) error {
	if dutyCycle < 0 || dutyCycle > 100 {
		return fmt.Errorf("dutyCycle=%v out of range [0-100]", dutyCycle)
	}

	rMax, found := c.DutyCycleRanges[gpio]
	if !found {
		res, err := c.PRRG(gpio)
		if err != nil {
			return err
		}
		rMax = res
	}

	r := uint32(float32(rMax) * float32(dutyCycle) / float32(100))
	return c.PWMRaw(gpio, r)
}
