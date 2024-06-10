package postgres

import (
	"fmt"
)

type ErrPostgres struct {
	err error
	row int
}

func (e *ErrPostgres) Unwrap() error { return e.err }
func (e *ErrPostgres) Row() int      { return e.row }
func (e *ErrPostgres) Error() string {
	return fmt.Sprintf("on row %d: %s", e.row, e.err.Error())
}

type batchProcessor interface {
	Exec(f func(i int, err error))
	Close() error
}

func process(r batchProcessor) error {
	var errExec error

	r.Exec(func(i int, err error) {
		if err != nil && errExec == nil {
			errExec = &ErrPostgres{err: err, row: i}
		}
	})

	if err := r.Close(); err != nil && errExec == nil {
		errExec = err
	}

	return errExec
}
