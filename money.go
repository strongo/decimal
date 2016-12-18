package decimal

import (
	"encoding/json"
	"math"
	"strconv"
	"strings"
)

type Money64p2 int64

const PRECISION_2 = 2

func NewMoney64p2(intPart int64, decimalPart int8) Money64p2 {
	if decimalPart < 0 {
		panic("decimalPart < 0")
	} else if decimalPart > 99 {
		panic("decimalPart > 99")
	}

	if intPart = intPart * 100; intPart >= 0 {
		return Money64p2(intPart + int64(decimalPart))
	} else {
		return Money64p2(intPart - int64(decimalPart))
	}
}

func (d Money64p2) IntPart() int64 {
	return int64(d/100)
}

func (d Money64p2) DecimalPart() int64 {
	result := int64(d - d/100*100)
	if result < 0 {
		result *= -1
	}
	return result
}

func (d Money64p2) String() string {
	if d == 0 {
		return "0"
	}
	s := strconv.FormatInt(int64(d), 10);
	if len(s) <= PRECISION_2 {
		return "0." + s
	} else {
		var left, right string
		left = s[:len(s) - PRECISION_2]
		right = s[len(s) - PRECISION_2:]
		if right == "00" {
			return left
		} else {
			return strings.Join([]string{left, right}, ".")
		}
	}
}

func ParseMoney64p2(s string) (d Money64p2, err error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return d, err
	}
	return FloatToMoney64p2(f), nil
}

func FloatToMoney64p2(f float64) Money64p2 {
	intPart := round(f / 100) * 100
	return Money64p2(intPart + round(toFixed(f - float64(intPart), PRECISION_2)*100))
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
}

func (d *Money64p2) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Money64p2) UnmarshalJSON(data []byte) error {
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	*d = FloatToMoney64p2(f)
	return nil
}