package chain10

var exponents map[string]uint64 = map[string]uint64{}

func Exponent(chainid uint64) (uint64, bool) {

	var key string = createkey(chainid)

	value, found := exponents[key]
	if !found {
		return 0, false
	}

	return value, true
}
