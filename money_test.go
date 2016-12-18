package decimal

import (
	"testing"
)

func TestNewMoney(t *testing.T) {
	m := NewMoney64p2(0, 0)
	if int64(m) != 0 {
		t.Errorf("Expected 0, got: %d", m)
	}
}

func TestParseMoney64p2(t *testing.T) {
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
	} else if d != NewMoney64p2(1, 0) {
		t.Errorf("Expected 1, got: %v", d)
	}

	if d, err = ParseMoney64p2("1.23"); err != nil {
		t.Error(err)
	} else if d != NewMoney64p2(1, 23) {
		t.Errorf("Expected 1.23, got: %d", d)
	}
}


func TestMoney64_String(t *testing.T) {
	m := NewMoney64p2(0, 0)
	s := m.String()
	if s != "0" {
		t.Errorf("Expected '0', got '%v'", s)
	}

	s = NewMoney64p2(0, 23).String()
	if s != "0.23" {
		t.Errorf("Expected '0.23', got '%v'", s)
	}

	s = NewMoney64p2(1, 23).String()
	if s != "1.23" {
		t.Errorf("Expected '1.23', got '%v'", s)
	}

	s = NewMoney64p2(45, 0).String()
	if s != "45" {
		t.Errorf("Expected '45', got '%v'", s)
	}

	s = NewMoney64p2(-45, 67).String()
	if s != "-45.67" {
		t.Errorf("Expected '-45.67', got '%v'", s)
	}
}

func TestMoney64p2_IntPart(t *testing.T) {
	m := NewMoney64p2(23, 45)
	if m.IntPart() != 23 {
		t.Error("m.IntPart() != 23")
	}

	m = NewMoney64p2(-23, 45)
	if m.IntPart() != -23 {
		t.Error("m.IntPart() != -23")
	}
}

func TestMoney64p2_DecimalPart(t *testing.T) {
	m := NewMoney64p2(23, 45)
	if m.DecimalPart() != 45 {
		t.Errorf("m.DecimalPart() != 45, got: %v", m.DecimalPart())
	}

	m = NewMoney64p2(-23, 45)
	if m.DecimalPart() != 45 {
		t.Errorf("m.DecimalPart() != 45, got: %v", m.DecimalPart())
	}
}