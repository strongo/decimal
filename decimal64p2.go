package decimal

import (
	"encoding/json"
	"math"
	"strconv"
	"strings"
)

type Decimal64p2 int64

const PRECISION_2 = 2

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

	return Decimal64p2(intPart * 100 + int64(decimalPart))
}

func NewDecimal64p2FromFloat64(f float64) Decimal64p2 {
	return Decimal64p2(round(f * 100))
}

func (d Decimal64p2) AsFloat64() float64 {
	return float64(d)/100
}

func (d Decimal64p2) IntPart() int64 {
	return int64(d/100)
}

func (d Decimal64p2) DecimalPart() int64 {
	result := int64(d - d/100*100)
	if result < 0 {
		result *= -1
	}
	return result
}

func (d Decimal64p2) String() string {
	if d == 0 {
		return "0"
	}
	var sign string
	i := int64(d)
	if i < 0 {
		sign = "-"
		i *= -1
	}
	s := strconv.FormatInt(i, 10);
	if i <= 9 {
		return sign + "0.0" + s
	} else if i <= 99 {
		return sign + "0." + s
	} else {
		var left, right string
		left = s[:len(s)-PRECISION_2]
		right = s[len(s)-PRECISION_2:]
		if right == "00" {
			return sign + left
		} else {
			return sign + strings.Join([]string{left, right}, ".")
		}
	}
}

func ParseDecimal64p2(s string) (d Decimal64p2, err error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return d, err
	}
	return NewDecimal64p2FromFloat64(f), nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
}

func (d Decimal64p2) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Decimal64p2) UnmarshalJSON(data []byte) error {
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	*d = NewDecimal64p2FromFloat64(f)
	return nil
}

func (d Decimal64p2) Abs() Decimal64p2 {
	if d < 0 {
		return d * -1
	}
	return d
}