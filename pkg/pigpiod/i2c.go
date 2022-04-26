package pigpiod

//pigs i2co 1 0x33 0

func (c *Conn) InitI2c(busAddr, device int) (handel uint32, err error) {
	cmd := Cmd{
		cmd: 54,
		p1:  uint32(busAddr),
		p2:  uint32(device),
	}
	res, err := cmd.ExecuteRes(c.tcp)
	return res.p3, err
}

func (c *Conn) CloseI2c(handel int) (err error) {
	cmd := Cmd{
		cmd: 55,
		p1:  uint32(handel),
	}
	_, err = cmd.ExecuteRes(c.tcp)
	return err
}

///pigs I2CRI 48 0xF 16

func (c *Conn) ReadI2c(handel, register, length int) ([]byte, error) {
	cmd := Cmd{
		cmd: 67,
		p1:  uint32(handel),
		p2:  uint32(register),
		p3:  uint32(length),
	}
	res, err := cmd.ExecuteResData(c.tcp)
	return res.data, err
}

//pigs I2CWB 48 0x33 0xDA
//              addr  byte
//I2CWB h r byte   SMBus Write Byte Data: write byte to register

func (c *Conn) WriteI2c(handel, busAddr, register int) ([]byte, error) {
	cmd := Cmd{
		cmd: 56,
		p1:  uint32(handel),
		p2:  uint32(busAddr),
		p3:  uint32(register),
	}
	res, err := cmd.ExecuteResData(c.tcp)
	return res.data, err
}
