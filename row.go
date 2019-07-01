package paginator

import (
	"fmt"
	"strconv"
	"time"
)

type Row map[string]string

type ErrNoFieldNameFound struct {
	name string
}

func (e *ErrNoFieldNameFound) Error() string {
	return fmt.Sprintf(`No field named "%s" was found in current row`, e.name)
}

func (r Row) String(name string) (string, error) {
	if v, ok := r[name]; ok {
		return v, nil
	}
	return "", &ErrNoFieldNameFound{name}
}

func (r Row) Int(name string) (int, error) {
	if v, ok := r[name]; ok {
		return strconv.Atoi(v)
	}
	return 0, &ErrNoFieldNameFound{name}
}

func (r Row) Float(name string) (float64, error) {
	if v, ok := r[name]; ok {
		return strconv.ParseFloat(v, 64)
	}
	return 0.0, &ErrNoFieldNameFound{name}
}

func (r Row) Time(name string) (time.Time, error) {
	if v, ok := r[name]; ok {
		return time.Parse(time.RFC3339, v)
	}
	return time.Time{}, &ErrNoFieldNameFound{name}
}
