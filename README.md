# Package `github.com/strongo/decimal`

[![Build Status](https://travis-ci.org/strongo/decimal.svg?branch=master)](https://travis-ci.org/strongo/decimal)
[![GoDoc](https://godoc.org/github.com/strongo/decimal?status.svg)](https://godoc.org/github.com/strongo/decimal)
[![Go Report Card](https://goreportcard.com/badge/github.com/strongo/decimal)](https://goreportcard.com/report/github.com/strongo/decimal)

Decimal 64 bit numbers implementation to represent money values in GoLang. Based on int64. Supports JSON (un)marshalling.

At the moment provides just a single type `Decimal64p2` with fixed precision of 2 digits after point.
In simple words it stores value as 64 bits integer amount of cents.

The code has <b>100% unit tests coverage</b>.

E.g. `1.43` will be stored as `int64(143)` but when rendered as string will be represented as `"1.43"`.

 
```go
package example

import "github.com/strongo/decimal"

func Example() {
	var amount decimal.Decimal64p2; print(amount)  // 0
	
	amount = decimal.NewDecimal64p2(0, 43); print(amount)  // 0.43
	amount = decimal.NewDecimal64p2(1, 43); print(amount)  // 1.43
	amount = decimal.NewDecimal64p2FromFloat64(23.100001); print(amount)  // 23.10
	amount, _ = decimal.ParseDecimal64p2("2.34"); print(amount)  // 2.34
	amount, _ = decimal.ParseDecimal64p2("-3.42"); print(amount)  // -3.42
}
```

This package originally was developed for <a href="https://debtstrcker.io/"><b>DebtsTracker.io</b></a> - a mobile app & chat bots to <b>split bills & track your debts</b>.

## Reasoning
* Fast
* Compact
* No precision issues with storing values like `0.10`
* By storing with precision to cents there is no ambiguity with rounding. E.g. if you split $10 between 3 persons the amounts will be $3.33, $3.33 & $3.3<b>4</b>.

## <a href="https://github.com/strongo/decimal/blob/master/LICENSE">MIT License</a>
Free to use without restrictions. If cloned please keep links to <a href="https://github.com/strongo/decimal">https://github.com/strongo/decimal</a> and to <a href="https://debtstracker.io/">https://debtstracker.io/</a>.

