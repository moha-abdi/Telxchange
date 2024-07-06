package exchange

import (
	"fmt"
	"strconv"
)

// Currency is a custom type for defining currency codes
type Currency int

// Define constants for each currency code
const (
	USD  Currency = 840
	SLSH Currency = 706
)

// String method to get the string representation of the currency
func (c Currency) String() string {
	switch c {
	case USD:
		return "USD"
	case SLSH:
		return "SLSH"
	default:
		return fmt.Sprintf("Unknown currency code (%d)", c)
	}
}

// Integer to ascii unlike String() this converts the datatype to string
func (c Currency) Itoa() string {
	return strconv.FormatInt(int64(c), 10)
}
