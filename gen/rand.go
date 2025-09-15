package gen

import (
	"math/rand"
	"strconv"
	"strings"
)

func PowerOfTen(exp int) int {
    result := 1
    for range exp {
        result *= 10
    }
    return result
}

func RandomWithDigits(r *rand.Rand, d int) int {
    if d <= 1 {
        return r.Intn(9) + 1
    }
    min := PowerOfTen(d - 1)
    span := PowerOfTen(d) - min
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
    return b.String()
}


