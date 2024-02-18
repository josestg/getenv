package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Func is a function that parses an env string into a value of type T.
type Func[T any] func(v string) (T, error)

// ID returns the given input as the output.
func ID(v string) (string, error) { return v, nil }

// U64 parses a string into an uint64.
func U64(v string) (uint64, error) { return strconv.ParseUint(v, 10, 64) }

// F64 parses a string into a float64.
func F64(v string) (float64, error) { return strconv.ParseFloat(v, 64) }

// Time creates a parser for time.Time with the given layout.
func Time(layout string) Func[time.Time] {
	return func(v string) (time.Time, error) { return time.Parse(layout, v) }
}

// JSON parses a string into a JSON value.
func JSON[T any](v string) (T, error) {
	var p T
	return p, json.Unmarshal([]byte(v), &p)
}

// Parse parses the env in the given key using the parser function.
func Parse[T any](key string, fallback T, p Func[T]) T {
	v, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	pv, err := p(v)
	if err != nil {
		panic(fmt.Errorf("parse key %q, env %q: %w", key, v, err))
	}
	return pv
}
