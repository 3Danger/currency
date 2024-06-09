package time

import (
	"fmt"
	"strings"
	"time"
)

type Time[T layouter] struct {
	time.Time
}

func (t *Time[T]) UnmarshalJSON(data []byte) error {
	var layouter T

	tt, err := time.Parse(layouter.Layout(), strings.Trim(string(data), "\""))
	if err != nil {
		return fmt.Errorf("parsing time: %w", err)
	}

	t.Time = tt

	return nil
}

func (t *Time[T]) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(time.DateTime)), nil
}
