package money

import "fmt"

// GBP money
type GBP struct{ Money }

// NewGBP returns a NewGBP money
func NewGBP(amount float64) GBP {
	m := new2DecimalFormatLeft("£", amount)
	return GBP{m}
}

// Add ...
func (m GBP) Add(rhs GBP) Money {
	return m.Clone(m.allDigits + rhs.allDigits)
}

// Subtract ...
func (m GBP) Subtract(rhs GBP) GBP {
	return GBP{m.Money.Clone(m.allDigits - rhs.allDigits)}
}

// GreaterOrEqual ...
func (m GBP) GreaterOrEqual(rhs GBP) bool {
	return m.allDigits >= rhs.allDigits
}

// Greater ...
func (m GBP) Greater(rhs GBP) bool {
	return m.allDigits > rhs.allDigits
}

// Equal ...
func (m GBP) Equal(rhs GBP) bool {
	return m.allDigits == rhs.allDigits
}

// LessThan ...
func (m GBP) LessThan(rhs GBP) bool {
	return m.allDigits < rhs.allDigits
}

// GetMoney ...
func (m GBP) GetMoney() Money {
	return m.Money
}

// Wide ...
func (m GBP) Wide() string {
	return fmt.Sprintf(m.formatWide, m.integralDigits, m.fractionalDigits)
}

// Short ...
func (m GBP) Short() string {
	return fmt.Sprintf(m.formatShort, m.integralDigits, m.fractionalDigits)
}
