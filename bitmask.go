package bitmask

func Reset(bitmask uint16) uint16 {
	return 0x0000
}

func Add(bitmask uint16, value uint16) uint16 {
	return (bitmask | value)
}

func Remove(bitmask uint16, value uint16) uint16 {
	return ((bitmask &^ value) & bitmask)
}

func Has(bitmask, value uint16) bool {
	return (bitmask & value) == value
}
