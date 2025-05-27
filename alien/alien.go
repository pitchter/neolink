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
	i := 0

	for i < n {
		curr := symbolValues[s[i]]
		if i+1 < n {
			next := symbolValues[s[i+1]]
			if curr < next {
				total += (next - curr)
				i += 2 
				continue
			}
		}
		total += curr
		i++
	}

	return total
}
