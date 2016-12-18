package decimal

import (
	"testing"
	"encoding/json"
)

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
	if d = NewDecimal64p2(-1, 23); int64(d) != -123 {
		t.Errorf("Expected -123, got: %d", d)
	}
}

func TestParseDecimal64p2(t *testing.T) {
	d, err := ParseMoney64p2("0")
	if err != nil {
		t.Error(err)
	}

	if d != 0 {
		t.Errorf("Expected 0, got: %v", d)
	}

	if d, err = ParseMoney64p2("0.00"); err != nil {
		t.Error(err)
	} else if d != 0 {
		t.Errorf("Expected 0, got: %v", d)
	}

	if d, err = ParseMoney64p2("1.00"); err != nil {
		t.Error(err)
	} else if d != NewDecimal64p2(1, 0) {
		t.Errorf("Expected 1, got: %v", d)
	}

	if d, err = ParseMoney64p2("1.23"); err != nil {
		t.Error(err)
	} else if d != NewDecimal64p2(1, 23) {
		t.Errorf("Expected 1.23, got: %d", d)
	}
}


func TestDecimal64p2_String(t *testing.T) {
	m := NewDecimal64p2(0, 0)
	s := m.String()
	if s != "0" {
		t.Errorf("Expected '0', got '%v'", s)
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

	s = NewDecimal64p2(-45, 67).String()
	if s != "-45.67" {
		t.Errorf("Expected '-45.67', got '%v'", s)
	}
}

func TestDecimal64p2_IntPart(t *testing.T) {
	m := NewDecimal64p2(23, 45)
	if m.IntPart() != 23 {
		t.Error("m.IntPart() != 23")
	}

	m = NewDecimal64p2(-23, 45)
	if m.IntPart() != -23 {
		t.Error("m.IntPart() != -23")
	}
}

func TestDecimal64p2_DecimalPart(t *testing.T) {
	m := NewDecimal64p2(23, 45)
	if m.DecimalPart() != 45 {
		t.Errorf("m.DecimalPart() != 45, got: %v", m.DecimalPart())
	}

	m = NewDecimal64p2(-23, 45)
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
	}
	if d = NewDecimal64p2FromFloat64(-1.23); int64(d) != -123 {
		t.Errorf("Expected -123, got: %d", d)
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
	} else if string(s) != d1.String() {
		t.Errorf("Expected '%v', got: '%v'", d1.String(), string(s))
	}

	d1 = NewDecimal64p2(-1, 23)
	if s, err := json.Marshal(d1); err != nil {
		t.Error(err)
	} else if string(s) != d1.String() {
		t.Errorf("Expected '%v', got: '%v'", d1.String(), string(s))
	}
}

func TestDecimal64p2_UnmarshalJSON(t *testing.T) {
	var d2 Decimal64p2

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