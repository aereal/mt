package mt

import "time"

type parseConfig struct {
	loc                  *time.Location
	defaultAllowComments bool
	defaultConvertBreaks ConvertBreaks
}

type ParseOption interface {
	applyParseOption(c *parseConfig)
}

func WithTimeLocation(loc *time.Location) *OptionWithTime { return &OptionWithTime{loc: loc} }

type OptionWithTime struct{ loc *time.Location }

var _ ParseOption = (*OptionWithTime)(nil)

func (o *OptionWithTime) applyParseOption(c *parseConfig) { c.loc = o.loc }

func WithDefaultAllowComments(v bool) *OptionWithDefaultAllowComments {
	return &OptionWithDefaultAllowComments{v: v}
}

type OptionWithDefaultAllowComments struct{ v bool }

var _ ParseOption = (*OptionWithDefaultAllowComments)(nil)

func (o *OptionWithDefaultAllowComments) applyParseOption(c *parseConfig) {
	c.defaultAllowComments = o.v
}

func WithDefaultConvertBreaks(cb ConvertBreaks) *OptionWithDefaultConvertBreaks {
	return &OptionWithDefaultConvertBreaks{cb: cb}
}

type OptionWithDefaultConvertBreaks struct{ cb ConvertBreaks }

var _ ParseOption = (*OptionWithDefaultConvertBreaks)(nil)

func (o *OptionWithDefaultConvertBreaks) applyParseOption(c *parseConfig) {
	c.defaultConvertBreaks = o.cb
}
