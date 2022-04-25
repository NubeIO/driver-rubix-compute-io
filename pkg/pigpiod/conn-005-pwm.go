package pigpiod

import (
	"fmt"
)

func (c *Conn) HardwarePWM(gpio int, dutyCycle uint32) error {
	cmd := Cmd{
		cmd: 5,
		p1:  uint32(gpio),
		p2:  4,
		p3:  dutyCycle,
	}
	_, err := cmd.ExecuteRes(c.tcp)
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
	_, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return err
	}
	return nil
}

func (c *Conn) PWM(gpio int, dutyCycle int) error {
	if dutyCycle < 0 || dutyCycle > 100 {
		return fmt.Errorf("dutyCycle=%v out of range [0-100]", dutyCycle)
	}

	rMax, found := c.dutyCycleRanges[gpio]
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
