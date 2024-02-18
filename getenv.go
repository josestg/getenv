package getenv

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/josestg/getenv/parser"
)

type (
	intType interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}

	floatType interface {
		~float32 | ~float64
	}
)

// String gets an env as string.
// If key not exists, fallback is returned.
func String(key, fallback string) string {
	return parser.Parse(key, fallback, parser.ID)
}

// Strings gets comma-separated env as string slice.
// If key not exists, fallback is returned.
func Strings(key string, fallback []string) []string {
	return List(key, fallback, parser.ID, identityOf[string])
}

// Int gets env as int variants, if not found, returns the fallback value.
// The type T determined by the fallback type.
func Int[T intType](key string, fallback T) T {
	v := parser.Parse(key, uint64(fallback), parser.U64)
	return u64toi[T](v)
}

// Ints gets comma-separated env as int slice, if not found, returns the fallback value.
// The type T determined by the fallback element type.
func Ints[T intType, S []T](key string, fallback S) S {
	return List(
		key,
		fallback,
		parser.U64,
		u64toi[T],
	)
}

// Float gets env as float, if not found, returns the fallback value.
// The type T determined by the fallback type.
func Float[T floatType](key string, fallback T) T {
	v := parser.Parse(key, float64(fallback), parser.F64)
	return f64tof[T](v)
}

// Floats gets comma-separated env as float slice, if not found, returns the fallback value.
// The type T determined by the fallback element type.
func Floats[T floatType, S []T](key string, fallback S) S {
	return List(
		key,
		fallback,
		parser.F64,
		f64tof[T],
	)
}

// Duration gets env as time.Duration, if not found, returns the fallback value.
func Duration(key string, fallback time.Duration) time.Duration {
	return parser.Parse(key, fallback, time.ParseDuration)
}

// Durations gets comma-separated env as time.Duration slice, if not found, returns the fallback value.
func Durations(key string, fallback []time.Duration) []time.Duration {
	return List(
		key,
		fallback,
		time.ParseDuration,
		identityOf[time.Duration],
	)
}

// Bool gets env as bool, if not found, returns the fallback value.
func Bool(key string, fallback bool) bool {
	return parser.Parse(key, fallback, strconv.ParseBool)
}

// Time gets env as time.Time, if not found, returns the fallback value.
func Time(key, layout string, fallback time.Time) time.Time {
	return parser.Parse(key, fallback, parser.Time(layout))
}

// JSON gets env as JSON, if not found, returns the fallback value.
func JSON[T any](key string, fallback T) T {
	return parser.Parse(key, fallback, parser.JSON[T])
}

// List parses comma-separated env into a slice of type T and maps each element using the converter function.
// If the env not set, returns the fallback value. If parsing fails for any of the values, it will panic.
func List[T, U any](key string, fallback []U, p parser.Func[T], converter func(T) U) []U {
	return parser.Parse(
		key,
		fallback,
		func(s string) ([]U, error) {
			parts := strings.Split(s, ",")
			values := make([]U, 0, len(parts))
			for _, part := range parts {
				v, err := p(part)
				if err != nil {
					return values, fmt.Errorf("parse key %q, part %q: %w", key, part, err)
				}
				values = append(values, converter(v))
			}
			return values, nil
		},
	)
}

func identityOf[T any](v T) T         { return v }
func u64toi[T intType](v uint64) T    { return T(v) }
func f64tof[T floatType](v float64) T { return T(v) }
