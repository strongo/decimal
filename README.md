# decimal
Decimal 64 bit implementation to represent money values in GoLang.

Based on int64 with precision of 2 digits after point.

In simple words it stores value as integer amount of cents.

E.g. `1.43` will be stored as `int64(143)` but when rendered as string will be represented as `"1.43"`.
 
```go
package example

import "github.com/strongo/decimal"

func Example() {
	var amount decimal.Money64p2; print(amount)  // 0
	
	amount = decimal.NewMoney64p2(0, 43); print(amount)  // 0.43
	amount = decimal.NewMoney64p2(1, 43); print(amount)  // 1.43
	amount = decimal.NewMoney64p2(23, 00); print(amount)  // 23
	amount, _ = decimal.ParseMoney64p2("2.34"); print(amount)  // 2.34
	amount, _ = decimal.ParseMoney64p2("-3.42"); print(amount)  // -3.42
}
```

This package originally was developed for <a href="https://debtstrcker.io/">DebtsTracker.io</a> - an app to track your personal debts. 