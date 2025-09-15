package add

import (
	"math/rand"
	"simplemath/internal/generateutil"
)

type Generator struct{}

func (Generator) Generate(r *rand.Rand, digits []int) string {
	ops := make([]int, len(digits))
	for i := range digits {
		ops[i] = generateutil.RandomWithDigits(r, digits[i])
	}
	return generateutil.JoinOperands(ops, "+")
}


