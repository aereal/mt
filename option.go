package mt

import "time"

type parseConfig struct {
	loc *time.Location
}

type ParseOption interface {
	applyParseOption(c *parseConfig)
}

func WithTimeLocation(loc *time.Location) *OptionWithTime { return &OptionWithTime{loc: loc} }

type OptionWithTime struct{ loc *time.Location }

var _ ParseOption = (*OptionWithTime)(nil)

func (o *OptionWithTime) applyParseOption(c *parseConfig) { c.loc = o.loc }
