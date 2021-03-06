package tinkoff

import (
	"fmt"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.String())), nil
}

func (t Time) String() string {
	original := time.Time(t)
	if original.IsZero() {
		return ""
	}
	return original.Format(time.RFC3339)
}

func serializeBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func errorConcat(first, second error) error {
	if first == nil {
		return second
	}
	if second == nil {
		return first
	}
	return fmt.Errorf("%s: %s", second, first)
}
