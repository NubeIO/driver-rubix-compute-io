package pigpiod

func (c *Conn) PRRG(gpio int) (uint32, error) {
	cmd := Cmd{
		cmd: 24,
		p1:  uint32(gpio),
	}
	res, err := cmd.ExecuteRes(c.Tcp)
	if err != nil {
		return res.p3, err
	}
	c.DutyCycleRanges[gpio] = res.p3
	return res.p3, nil
}
