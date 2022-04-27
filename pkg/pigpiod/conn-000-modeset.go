package pigpiod

func (c *Conn) ModeSet(gpio int, mode GpioMode) error {
	cmd := Cmd{
		cmd: 0,
		p1:  uint32(gpio),
		p2:  uint32(mode),
	}
	_, err := cmd.ExecuteRes(c.Tcp)
	if err != nil {
		return err
	}
	return nil
}
