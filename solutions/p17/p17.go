package p17

import "strings"

var tens = map[int]string{
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

var others = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

// LetterCount returns the number of runes in a string that aren't basic hyphens
// or spaces.
func LetterCount(s string) int {
	var res int
	for _, r := range s {
		if r == '-' || r == ' ' {
			continue
		}
		res++
	}
	return res
}

// Words returns the written representation of a number. It only works for 1
// through 1000 inclusive.
func Words(n int) string {
	if n == 1000 {
		return "one thousand"
	}

	if n < 20 {
		return others[n]
	}

	var b strings.Builder

	if n >= 100 {
		b.WriteString(others[n/100])
		b.WriteString(" hundred")
		n %= 100
		if n == 0 {
			return b.String()
		}
		b.WriteString(" and ")
	}

	if n < 20 {
		b.WriteString(others[n])
		return b.String()
	}

	tensPlace, onesPlace := n/10, n%10

	b.WriteString(tens[tensPlace])

	if onesPlace > 0 {
		b.WriteString("-")
		b.WriteString(others[onesPlace])
	}

	return b.String()
}

func Solve() int {
	var res int
	for i := 1; i <= 1000; i++ {
		res += LetterCount(Words(i))
	}
	return res
}
