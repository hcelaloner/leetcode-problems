package src

var symbolValueMapping = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	intValue := 0
	for i := len(s) - 1; i >= 0; i-- {
		valueOfSymbol := symbolValueMapping[s[i]]
		intValue += valueOfSymbol

		if i-1 >= 0 {
			valueOfNextSymbol := symbolValueMapping[s[i-1]]
			if valueOfSymbol > valueOfNextSymbol {
				intValue -= valueOfNextSymbol * 2
			}
		}

	}
	return intValue
}
