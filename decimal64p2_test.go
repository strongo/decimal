package decimal

import (
	"encoding/json"
	"math"
	"testing"
)

func TestParseDecimal64p2ExactAndBounds(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  Decimal64p2
	}{
		{"0.1", 10},
		{"-0", 0},
		{"-0.00", 0},
		{"+1.2", 120},
		{"92233720368547758.07", math.MaxInt64},
		{"-92233720368547758.08", math.MinInt64},
	} {
		got, err := ParseDecimal64p2(tc.input)
		if err != nil || got != tc.want {
			t.Errorf("ParseDecimal64p2(%q) = %d, %v; want %d", tc.input, got, err, tc.want)
		}
	}
	for _, input := range []string{
		"", "+", "-", ".1", "1.", "0.005", "1.234", "1e2", "NaN", "Inf", "-Inf",
		"1.2.3", " 1", "1 ", "1a", "92233720368547758.08", "-92233720368547758.09",
		"9223372036854775808", "-9223372036854775809",
	} {
		if got, err := ParseDecimal64p2(input); err == nil {
			t.Errorf("ParseDecimal64p2(%q) = %d; want error", input, got)
		}
	}
}

func TestDecimal64p2StringBounds(t *testing.T) {
	for _, tc := range []struct {
		value Decimal64p2
		want  string
	}{
		{math.MaxInt64, "92233720368547758.07"},
		{math.MinInt64, "-92233720368547758.08"},
	} {
		if got := tc.value.String(); got != tc.want {
			t.Errorf("String(%d) = %q; want %q", tc.value, got, tc.want)
		}
	}
}

func TestDecimal64p2JSONExactAndCents(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  Decimal64p2
	}{
		{"0.1", 10},
		{"-92233720368547758.08", math.MinInt64},
		{"92233720368547758.07", math.MaxInt64},
		{"9223372036854775807", math.MaxInt64},
		{"-9223372036854775808", math.MinInt64},
	} {
		var got Decimal64p2
		if err := json.Unmarshal([]byte(tc.input), &got); err != nil || got != tc.want {
			t.Errorf("Unmarshal(%q) = %d, %v; want %d", tc.input, got, err, tc.want)
		}
	}
	for _, input := range []string{"0.005", "1e2", "9223372036854775808", "92233720368547758.08", `"1.23"`} {
		got := Decimal64p2(123)
		if err := json.Unmarshal([]byte(input), &got); err == nil || got != 123 {
			t.Errorf("Unmarshal(%q) = %d, %v; want unchanged value and error", input, got, err)
		}
	}
}

func TestNewDecimal64p2(t *testing.T) {
	var d Decimal64p2

	if d = NewDecimal64p2(0, 0); int64(d) != 0 {
		t.Errorf("Expected 0, got: %d", d)
	}

	if d = NewDecimal64p2(1, 0); int64(d) != 100 {
		t.Errorf("Expected 1, got: %d", d)
	}

	if d = NewDecimal64p2(-1, 0); int64(d) != -100 {
		t.Errorf("Expected -1, got: %d", d)
	}

	if d = NewDecimal64p2(1, 23); int64(d) != 123 {
		t.Errorf("Expected 123, got: %d", d)
	}
	if d = NewDecimal64p2(-1, -23); int64(d) != -123 {
		t.Errorf("Expected -123, got: %d", d)
	}
	if d = NewDecimal64p2(0, -23); int64(d) != -23 {
		t.Errorf("Expected -23, got: %d", d)
	}
}

func TestParseDecimal64p2(t *testing.T) {
	d, err := ParseDecimal64p2("0")
	if err != nil {
		t.Error(err)
	}

	if d != 0 {
		t.Errorf("Expected 0, got: %v", d)
	}

	if d, err = ParseDecimal64p2("0.00"); err != nil {
		t.Error(err)
	} else if d != 0 {
		t.Errorf("Expected 0, got: %v", d)
	} else if //goland:noinspection GoDfaNilDereference
	d.String() != "0" {
		t.Errorf("Expected 0, got: %v", d.String())
	}

	if d, err = ParseDecimal64p2("1.00"); err != nil {
		t.Error(err)
	} else if d != NewDecimal64p2(1, 0) {
		t.Errorf("Expected 1, got: %v", d)
	} else if d.String() != "1" {
		t.Errorf("Expected 1, got: %v", d.String())
	}

	if d, err = ParseDecimal64p2("1.23"); err != nil {
		t.Error(err)
	} else if d != NewDecimal64p2(1, 23) {
		t.Errorf("Expected 1.23, got: %d", d)
	} else if d.String() != "1.23" {
		t.Errorf("Expected 1.23, got: %v", d.String())
	}

	if d, err = ParseDecimal64p2("0.03"); err != nil {
		t.Error(err)
	} else if d != NewDecimal64p2(0, 3) {
		t.Errorf("Expected 0.03, got: %d", d)
	} else if d.String() != "0.03" {
		t.Errorf("Expected 0.03, got: %v", d.String())
	}
}

func TestDecimal64p2_String(t *testing.T) {
	m := NewDecimal64p2(0, 0)
	s := m.String()
	if s != "0" {
		t.Errorf("Expected '0', got '%v'", s)
	}

	s = NewDecimal64p2(0, 3).String()
	if s != "0.03" {
		t.Errorf("Expected '0.03', got '%v'", s)
	}

	s = NewDecimal64p2(0, 23).String()
	if s != "0.23" {
		t.Errorf("Expected '0.23', got '%v'", s)
	}

	s = NewDecimal64p2(1, 23).String()
	if s != "1.23" {
		t.Errorf("Expected '1.23', got '%v'", s)
	}

	s = NewDecimal64p2(45, 0).String()
	if s != "45" {
		t.Errorf("Expected '45', got '%v'", s)
	}

	s = NewDecimal64p2(-45, -67).String()
	if s != "-45.67" {
		t.Errorf("Expected '-45.67', got '%v'", s)
	}

	s = NewDecimal64p2FromFloat64(-0.03).String()
	if s != "-0.03" {
		t.Errorf("Expected '-0.03', got '%v'", s)
	}
}

func TestDecimal64p2_IntPart(t *testing.T) {
	m := NewDecimal64p2(23, 45)
	if m.IntPart() != 23 {
		t.Error("m.IntPart() != 23")
	}

	m = NewDecimal64p2(-23, -45)
	if m.IntPart() != -23 {
		t.Error("m.IntPart() != -23")
	}
}

func TestDecimal64p2_DecimalPart(t *testing.T) {
	m := NewDecimal64p2(23, 45)
	if m.DecimalPart() != 45 {
		t.Errorf("m.DecimalPart() != 45, got: %v", m.DecimalPart())
	}

	m = NewDecimal64p2(-23, -45)
	if m.DecimalPart() != 45 {
		t.Errorf("m.DecimalPart() != 45, got: %v", m.DecimalPart())
	}
}

func TestNewDecimal64p2FromFloat64(t *testing.T) {
	var d Decimal64p2
	if d = NewDecimal64p2FromFloat64(0); int64(d) != 0 {
		t.Errorf("Expected 0, got: %d", d)
	}
	if d = NewDecimal64p2FromFloat64(1.23); int64(d) != 123 {
		t.Errorf("Expected 123, got: %d", d)
	} else if d.DecimalPart() != 23 {
		t.Errorf("Decimal part expected to be 23, got: %d", d.DecimalPart())
	}
	if d = NewDecimal64p2FromFloat64(-1.23); int64(d) != -123 {
		t.Errorf("Expected -123, got: %d", d)
	} else if d.DecimalPart() != 23 {
		t.Errorf("Decimal part expected to be 23, got: %d", d.DecimalPart())
	}
	if d = NewDecimal64p2FromFloat64(2333.33); int64(d) != 233333 {
		t.Errorf("Expected 233333, got: %d", d)
	} else if d.DecimalPart() != 33 {
		t.Errorf("Decimal part expected to be 33, got: %d", d.DecimalPart())
	}
}

func TestNewDecimal64p2FromInt(t *testing.T) {
	var d Decimal64p2
	if d = NewDecimal64p2FromInt(0); int64(d) != 0 {
		t.Errorf("Expected 0, got: %d", d)
	} else if d.DecimalPart() != 0 {
		t.Errorf("Decimal part expected to be 0, got: %d", d.DecimalPart())
	}
	if d = NewDecimal64p2FromInt(123); int64(d) != 12300 {
		t.Errorf("Expected 12345, got: %d", d)
	} else if dp := d.DecimalPart(); dp != 0 {
		t.Errorf("Decimal part expected to be 0, got: %d", dp)
	}
	if d = NewDecimal64p2FromInt(-123); int64(d) != -12300 {
		t.Errorf("Expected -12300, got: %d", d)
	} else if dp := d.DecimalPart(); dp != 0 {
		t.Errorf("Decimal part expected to be 0, got: %d", dp)
	}
}

func TestDecimal64p2_AsFloat64(t *testing.T) {
	var d Decimal64p2
	d = NewDecimal64p2(1, 23)
	if d.AsFloat64() != 1.23 {
		t.Errorf("Expected 1.23, got: %v", d)
	}

	d = NewDecimal64p2(1, 05)
	if d.AsFloat64() != 1.05 {
		t.Errorf("Expected 1.05, got: %v", d)
	}
}

func TestDecimal64p2_MarshalJSON(t *testing.T) {
	var d1 Decimal64p2

	d1 = NewDecimal64p2(1, 23)
	if s, err := json.Marshal(d1); err != nil {
		t.Error(err)
	} else if expected := "123"; string(s) != expected {
		t.Errorf("Expected '%v', got: '%v'", expected, string(s))
	}

	d1 = NewDecimal64p2(-1, -23)
	if s, err := json.Marshal(d1); err != nil {
		t.Error(err)
	} else if expected := "-123"; string(s) != expected {
		t.Errorf("Expected '%v', got: '%v'", expected, string(s))
	}
}

func TestDecimal64p2_UnmarshalJSON(t *testing.T) {
	var d2 Decimal64p2

	if err := json.Unmarshal([]byte("1234"), &d2); err != nil {
		t.Error(err)
	} else if intPart := d2.IntPart(); intPart != 12 {
		t.Errorf("Expected 12 for int part, got %d: %s", intPart, d2.String())
	} else if decimalPart := d2.DecimalPart(); decimalPart != 34 {
		t.Errorf("Expected 0 for decimal part, got %d: %s", decimalPart, d2.String())
	}

	if err := json.Unmarshal([]byte("1.23"), &d2); err != nil {
		t.Error(err)
	} else if d2.IntPart() != 1 || d2.DecimalPart() != 23 {
		t.Errorf("Expected 1.23, got %v", d2)
	}

	if err := json.Unmarshal([]byte("-1.23"), &d2); err != nil {
		t.Error(err)
	} else if d2.IntPart() != -1 || d2.DecimalPart() != 23 {
		t.Errorf("Expected -1.23, got %v", d2)
	}
}

func TestDecimal64p2_Abs(t *testing.T) {
	d1 := NewDecimal64p2FromFloat64(-3.24)
	d2 := d1.Abs()

	if d2 < 0 {
		t.Error("Abs() returned < 0")
	}

	if d2.IntPart() != 3 {
		t.Error("d2.IntPart() != 3")
	}

	if d2.DecimalPart() != 24 {
		t.Error("d2.IntPart() != 24")
	}

	d1 = NewDecimal64p2FromFloat64(3.24)
	d2 = d1.Abs()

	if d1 != d2 {
		t.Error("d1 != d2")
	}
}

func TestNewDecimal64p2_panicPositiveNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("No panic")
		}
	}()
	NewDecimal64p2(1, -1)
}

func TestNewDecimal64p2_panicNegativePositive(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("No panic")
		}
	}()
	NewDecimal64p2(-1, 1)
}

func TestNewDecimal64p2_panicDecimalGreaterPlus99(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("No panic")
		}
	}()
	NewDecimal64p2(1, 100)
}

func TestNewDecimal64p2_panicDecimalLessMinus99(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("No panic")
		}
	}()
	NewDecimal64p2(-1, -100)
}
