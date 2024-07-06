package network

import "strconv"

type NetworkCode int

func (n NetworkCode) Itoa() string {
	return strconv.FormatInt(int64(n), 10)
}

// Network represents a payment network
type Network struct {
	Name string
	Code NetworkCode
}

// Define constants for network codes
const (
	sabCode NetworkCode = iota + 106
	evcCode
	zaadCode
	sahalCode
	cashPlusCode
)

// Define available networks
var (
	Zaad     = Network{Name: "ZAAD", Code: zaadCode}
	Evc      = Network{Name: "EVC", Code: evcCode}
	Sahal    = Network{Name: "SAHAL", Code: sahalCode}
	CashPlus = Network{Name: "CASHPLUS", Code: cashPlusCode}
	Sab      = Network{Name: "Sab", Code: sabCode}
)

// AllNetworks is a slice containing all available networks
var AllNetworks = []Network{Zaad, Evc, Sahal, CashPlus, Sab}
