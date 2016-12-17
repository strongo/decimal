package decimal

import strconv

type Money int64

func (d Money) String() string {
  return strconv.FormatInt(d, 10)
}
