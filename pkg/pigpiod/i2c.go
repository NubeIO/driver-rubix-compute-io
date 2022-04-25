package pigpiod

///pigs i2crd 1 16

func (c *Conn) ReadI2c(busAddr, length int) ([]byte, error) {
	cmd := Cmd{
		cmd: 56,
		p1:  uint32(busAddr),
		p2:  uint32(length),
	}
	res, err := cmd.ExecuteResData(c.tcp)
	return res.data, err
}
