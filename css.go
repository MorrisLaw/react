package react

import "github.com/gopherjs/gopherjs/js"

type CSS struct {
	o *js.Object

	Height    string
	MaxHeight string
	MinHeight string
	Overflow  string
	Resize    string
	Width     string
}

// TODO: until we have a resolution on
// https://github.com/gopherjs/gopherjs/issues/236 we define hack() below

func (c *CSS) hack() *CSS {
	if c == nil {
		return nil
	}

	o := object.New()

	o.Set("height", c.Height)
	o.Set("maxHeight", c.MaxHeight)
	o.Set("minHeight", c.MinHeight)
	o.Set("overflow", c.Overflow)
	o.Set("resize", c.Resize)
	o.Set("width", c.Width)

	return &CSS{o: o}
}
