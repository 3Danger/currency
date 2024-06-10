package time

import (
	"fmt"
	"strings"
	"time"
)

type Time[T layouter] struct {
	time.Time
}

func NewFromString[T layouter](data string) (Time[T], error) {
	t := Time[T]{}

	if err := t.UnmarshalJSON([]byte(data)); err != nil {
		return Time[T]{}, fmt.Errorf("time unmarshal: %w", err)
	}

	return t, nil
}

func NewFrom[T layouter](t time.Time) Time[T] { return Time[T]{t} }

func Now[T layouter]() Time[T] { return Time[T]{Time: time.Now()} }

func (t Time[T]) UTC() Time[T] {
	return Time[T]{Time: t.Time.UTC()}
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

func (t Time[T]) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(time.DateTime)), nil
}

func (t Time[T]) String() string { return t.Time.Format(time.DateTime) }
