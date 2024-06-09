package time

import "time"

type layouter interface {
	Layout() string
}

type LayoutDateTime struct{}

func (LayoutDateTime) Layout() string {
	return time.DateTime
}
