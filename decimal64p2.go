package decimal

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Decimal64p2 is a decimal number implementation based on int64 with 2 digits after point fixed precision
type Decimal64p2 int64

// Defines fixed precision with 2 digits after point
const (
	precision2 = 2
	pointPart2 = 100
)

// NewDecimal64p2FromInt creates Decimal64p2 from integer
func NewDecimal64p2FromInt(intPart int) Decimal64p2 {
	return Decimal64p2(intPart * pointPart2)
}

// NewDecimal64p2 creates Decimal64p2 from integer and decimal parts
func NewDecimal64p2(intPart int64, decimalPart int8) Decimal64p2 {
	switch {
	case decimalPart > 0:
		if decimalPart > 99 {
			panic("decimalPart > 99")
		}
		if intPart < 0 {
			panic("decimalPart > 0 && intPart < 0")
		}
	case decimalPart < 0:
		if decimalPart < -99 {
			panic("decimalPart < -99")
		}
		if intPart > 0 {
			panic("decimalPart < 0 && intPart > 0")
		}
	}

	return Decimal64p2(intPart*pointPart2 + int64(decimalPart))
}

// NewDecimal64p2FromFloat64 creates Decimal64p2 from float64
func NewDecimal64p2FromFloat64(f float64) Decimal64p2 {
	return Decimal64p2(round(f * pointPart2))
}

// AsFloat64 converts decimal to float64
func (d Decimal64p2) AsFloat64() float64 {
	return float64(d) / pointPart2
}

// IntPart returns integer part of the decimal
func (d Decimal64p2) IntPart() int64 {
	return int64(d / pointPart2)
}

// DecimalPart returns part after point
func (d Decimal64p2) DecimalPart() int64 {
	result := int64(d - d/pointPart2*pointPart2)
	if result < 0 {
		result *= -1
	}
	return result
}

// String renders decimal to string. If integer the .00 is NOT rendered.
func (d Decimal64p2) String() string {
	if d == 0 {
		return "0"
	}
	var sign string
	i := uint64(d)
	if d < 0 {
		sign = "-"
		i = uint64(-(int64(d) + 1)) + 1
	}
	s := strconv.FormatUint(i, 10)
	if i <= 9 {
		return sign + "0.0" + s
	} else if i <= 99 {
		return sign + "0." + s
	}

	var left, right string
	left = s[:len(s)-precision2]
	right = s[len(s)-precision2:]
	if right == "00" {
		return sign + left
	}
	return sign + strings.Join([]string{left, right}, ".")
}

// ParseDecimal64p2 creates Decimal64p2 from a string
func ParseDecimal64p2(s string) (d Decimal64p2, err error) {
	original := s
	negative := false
	if len(s) > 0 && (s[0] == '-' || s[0] == '+') {
		negative = s[0] == '-'
		s = s[1:]
	}
	invalid := func() (Decimal64p2, error) {
		return 0, fmt.Errorf("invalid decimal %q", original)
	}
	if len(s) == 0 {
		return invalid()
	}
	limit := uint64(math.MaxInt64)
	if negative {
		limit++
	}
	var cents uint64
	wholeDigits, fractionDigits := 0, 0
	point := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '.' && !point && wholeDigits > 0 {
			point = true
			continue
		}
		if c < '0' || c > '9' {
			return invalid()
		}
		if point {
			fractionDigits++
			if fractionDigits > precision2 {
				return invalid()
			}
		} else {
			wholeDigits++
		}
		digit := uint64(c - '0')
		if cents > (limit-digit)/10 {
			return 0, fmt.Errorf("decimal %q overflows int64 cents", original)
		}
		cents = cents*10 + digit
	}
	if wholeDigits == 0 || (point && fractionDigits == 0) {
		return invalid()
	}
	for ; fractionDigits < precision2; fractionDigits++ {
		if cents > limit/10 {
			return 0, fmt.Errorf("decimal %q overflows int64 cents", original)
		}
		cents *= 10
	}
	if negative {
		if cents == uint64(math.MaxInt64)+1 {
			return Decimal64p2(math.MinInt64), nil
		}
		return Decimal64p2(-int64(cents)), nil
	}
	return Decimal64p2(cents), nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

//func toFixed(num float64, precision int) float64 {
//	output := math.Pow(10, float64(precision))
//	return float64(round(num*output)) / output
//}

// MarshalJSON marshals decimal to JSON
func (d Decimal64p2) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(d), 10)), nil
}

// UnmarshalJSON deserializes JSON to decimal
func (d *Decimal64p2) UnmarshalJSON(data []byte) error {
	if strings.ContainsRune(string(data), '.') {
		parsed, err := ParseDecimal64p2(string(data))
		if err != nil {
			return err
		}
		*d = parsed
		return nil
	}
	cents, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*d = Decimal64p2(cents)
	return nil
}

// Abs returns absolute value for the decimal
func (d Decimal64p2) Abs() Decimal64p2 {
	if d < 0 {
		return d * -1
	}
	return d
}
