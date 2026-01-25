package gen

import (
	"math/rand"
	"strconv"
	"strings"
)

func Random(min int, max int) int {
	r := rand.New(rand.NewSource(rand.Int63()))
	span := max - min
	return min + r.Intn(span)
}

func JoinOperands(operands []int, symbol string) string {
	var b strings.Builder
	for idx, v := range operands {
		if idx > 0 {
			b.WriteString(" ")
			b.WriteString(symbol)
			b.WriteString(" ")
		}
		b.WriteString(strconv.Itoa(v))
	}
	return b.String() + " = "
}
