package getenv

import (
	"reflect"
	"testing"
	"time"
)

func TestInt(t *testing.T) {
	t.Run("uint", runIntTest[uint])
	t.Run("uint8", runIntTest[uint8])
	t.Run("uint16", runIntTest[uint16])
	t.Run("uint32", runIntTest[uint32])
	t.Run("uint64", runIntTest[uint64])
	t.Run("int", runIntTest[int])
	t.Run("int8", runIntTest[int8])
	t.Run("int16", runIntTest[int16])
	t.Run("int32", runIntTest[int32])
	t.Run("int64", runIntTest[int64])
}

func TestInts(t *testing.T) {
	t.Run("uint", runIntsTest[uint])
	t.Run("uint8", runIntsTest[uint8])
	t.Run("uint16", runIntsTest[uint16])
	t.Run("uint32", runIntsTest[uint32])
	t.Run("uint64", runIntsTest[uint64])
	t.Run("int", runIntsTest[int])
	t.Run("int8", runIntsTest[int8])
	t.Run("int16", runIntsTest[int16])
	t.Run("int32", runIntsTest[int32])
	t.Run("int64", runIntsTest[int64])
}

func TestFloat(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		var fallback float32 = 42.42
		got := Float("test_env_float", fallback)
		if got != fallback {
			t.Errorf("Float() = %v, want %v", got, fallback)
		}

		t.Setenv("test_env_float", "123.123")
		got = Float("test_env_float", fallback)
		if got != float32(123.123) {
			t.Errorf("Float() = %v, want 123.123", got)
		}
	})

	t.Run("float64", func(t *testing.T) {
		var fallback float64 = 42.42
		got := Float("test_env_float", fallback)
		if got != fallback {
			t.Errorf("Float() = %v, want %v", got, fallback)
		}

		t.Setenv("test_env_float", "123.123")
		got = Float("test_env_float", fallback)
		if got != 123.123 {
			t.Errorf("Float() = %v, want 123.123", got)
		}
	})

	t.Run("panics", func(t *testing.T) {
		t.Setenv("test_env_float_panics", "abc")
		defer func() {
			if recover() == nil {
				t.Error("Float() did not panic")
			}
		}()
		Float("test_env_float_panics", 42.42)
	})
}

func TestFloats(t *testing.T) {
	t.Run("float32", func(t *testing.T) {
		fallback := []float32{42.42, 43.43, 44.44}
		got := Floats("test_env_floats", fallback)
		if !reflect.DeepEqual(got, fallback) {
			t.Errorf("Floats() = %v, want %v", got, fallback)
		}

		t.Setenv("test_env_floats", "123.123,124.124")
		got = Floats("test_env_floats", fallback)
		if !reflect.DeepEqual(got, []float32{123.123, 124.124}) {
			t.Errorf("Floats() = %v, want [123.123, 124.124]", got)
		}

		t.Setenv("test_env_floats_panics", "123.123,a")
		defer func() {
			if recover() == nil {
				t.Error("Floats() did not panic")
			}
		}()
		Floats("test_env_floats_panics", fallback)
	})

	t.Run("float64", func(t *testing.T) {
		fallback := []float64{42.42, 43.43, 44.44}
		got := Floats("test_env_floats", fallback)
		if !reflect.DeepEqual(got, fallback) {
			t.Errorf("Floats() = %v, want %v", got, fallback)
		}

		t.Setenv("test_env_floats", "123.123,124.124")
		got = Floats("test_env_floats", fallback)
		if !reflect.DeepEqual(got, []float64{123.123, 124.124}) {
			t.Errorf("Floats() = %v, want [123.123, 124.124]", got)
		}

		t.Setenv("test_env_floats_panics", "123.123,a")
		defer func() {
			if recover() == nil {
				t.Error("Floats() did not panic")
			}
		}()
		Floats("test_env_floats_panics", fallback)
	})
}

func TestString(t *testing.T) {
	fallback := "fallback"
	got := String("test_env_string", fallback)
	if got != fallback {
		t.Errorf("String() = %v, want %v", got, fallback)
	}

	t.Setenv("test_env_string", "value")
	got = String("test_env_string", fallback)
	if got != "value" {
		t.Errorf("String() = %v, want value", got)
	}
}

func TestStrings(t *testing.T) {
	fallback := []string{"fallback"}
	got := Strings("test_env_strings", fallback)
	if !reflect.DeepEqual(got, fallback) {
		t.Errorf("Strings() = %v, want %v", got, fallback)
	}

	t.Setenv("test_env_strings", "value1,value2")
	got = Strings("test_env_strings", fallback)
	if !reflect.DeepEqual(got, []string{"value1", "value2"}) {
		t.Errorf("Strings() = %v, want [value1, value2]", got)
	}
}

func TestDuration(t *testing.T) {
	fallback := time.Second
	got := Duration("test_env_duration", fallback)
	if got != time.Second {
		t.Errorf("Duration() = %v, want 1s", got)
	}

	t.Setenv("test_env_duration", "1m")
	got = Duration("test_env_duration", fallback)
	if got != time.Minute {
		t.Errorf("Duration() = %v, want 1m", got)
	}
}

func TestDurations(t *testing.T) {
	fallback := []time.Duration{time.Second}
	got := Durations("test_env_durations", fallback)
	if !reflect.DeepEqual(got, fallback) {
		t.Errorf("Durations() = %v, want %v", got, fallback)
	}

	t.Setenv("test_env_durations", "1m,2m")
	got = Durations("test_env_durations", fallback)
	if !reflect.DeepEqual(got, []time.Duration{time.Minute, 2 * time.Minute}) {
		t.Errorf("Durations() = %v, want [1m, 2m]", got)
	}
}

func TestBool(t *testing.T) {
	got := Bool("test_env_bool", true)
	if got != true {
		t.Errorf("Bool() = %v, want true", got)
	}

	t.Setenv("test_env_bool", "false")
	got = Bool("test_env_bool", true)
	if got != false {
		t.Errorf("Bool() = %v, want false", got)
	}
}

func TestTime(t *testing.T) {
	fallback := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	got := Time("test_env_time", "2006-01-02", fallback)
	if got != fallback {
		t.Errorf("Time() = %v, want 2021-01-01", got)
	}

	t.Setenv("test_env_time", "2021-01-02")
	got = Time("test_env_time", "2006-01-02", fallback)
	if got != time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC) {
		t.Errorf("Time() = %v, want 2021-01-02", got)
	}
}

func TestJSON(t *testing.T) {
	type data struct {
		A int
		B string
	}

	fallback := data{A: 42, B: "fallback"}
	got := JSON("test_env_json", fallback)
	if got != fallback {
		t.Errorf("JSON() = %v, want %v", got, fallback)
	}

	t.Setenv("test_env_json", `{"A": 123, "B": "value"}`)
	got = JSON("test_env_json", fallback)
	if got != (data{A: 123, B: "value"}) {
		t.Errorf("JSON() = %v, want {123, value}", got)
	}
}

func runIntTest[T intType](t *testing.T) {
	var fallback T = 42
	got := Int("test_env_int", fallback)
	if got != fallback {
		t.Errorf("Int() = %v, want %v", got, fallback)
	}

	t.Setenv("test_env_int", "123")
	got = Int("test_env_int", fallback)
	if got != T(123) {
		t.Errorf("Int() = %v, want 123", got)
	}

	t.Setenv("test_env_int_panics", "abc")
	defer func() {
		if recover() == nil {
			t.Error("Int() did not panic")
		}
	}()
	Int("test_env_int_panics", fallback)
}

func runIntsTest[T intType](t *testing.T) {
	fallback := []T{42, 43, 44}
	got := Ints("test_env_ints", fallback)
	if !reflect.DeepEqual(got, fallback) {
		t.Errorf("Ints() = %v, want %v", got, fallback)
	}

	t.Setenv("test_env_ints", "123,124")
	got = Ints("test_env_ints", fallback)
	if !reflect.DeepEqual(got, []T{123, 124}) {
		t.Errorf("Ints() = %v, want [123, 124]", got)
	}

	t.Setenv("test_env_ints_panics", "123,a")
	defer func() {
		if recover() == nil {
			t.Error("Ints() did not panic")
		}
	}()
	Ints("test_env_ints_panics", fallback)
}
