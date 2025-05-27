package alien

var symbolValues = map[byte]int{
	'A': 1,
	'B': 5,
	'Z': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'R': 1000,
}

func AlienToInt(s string) int {
	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		value := symbolValues[s[i]]
		if i+1 < n && symbolValues[s[i]] < symbolValues[s[i+1]] {
			total -= value
		} else {
			total += value
		}
	}

	return total
}
