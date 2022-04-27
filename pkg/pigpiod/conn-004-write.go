package pigpiod

func (c *Conn) WriteOn(gpio int) error {
	cmd := Cmd{
		cmd: 4,
		p1:  uint32(gpio),
		p2:  uint32(LevelHigh),
	}
	_, err := cmd.ExecuteRes(c.Tcp)
	if err != nil {
		return err
	}
	return nil
}

func (c *Conn) WriteOff(gpio int) error {
	cmd := Cmd{
		cmd: 4,
		p1:  uint32(gpio),
		p2:  uint32(LevelLow),
	}
	_, err := cmd.ExecuteRes(c.Tcp)
	if err != nil {
		return err
	}
	return nil
}
