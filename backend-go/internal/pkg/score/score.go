package score

import (
	"fmt"

	ld "github.com/agnivade/levenshtein"
)

func Score(a string, b string) float64 {
	fmt.Println("string a", a)
	fmt.Println("string b", b)
	dist := ld.ComputeDistance(a, b)
	return float64(len(a)-dist) / float64(len(a))
}
