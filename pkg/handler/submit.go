package handler

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"simplemath/pkg/add"
	"simplemath/pkg/div"
	"simplemath/pkg/mul"
	"simplemath/pkg/sub"

	"github.com/labstack/echo/v4"
)

type Generator interface {
	Generate(r *rand.Rand, digits []int) string
}

type SubmitHandler struct {
	rng  *rand.Rand
	gens map[string]Generator
}

func NewSubmitHandler(r *rand.Rand) *SubmitHandler {
	return &SubmitHandler{
		rng: r,
		gens: map[string]Generator{
			"addition":       add.Generator{},
			"subtraction":    sub.Generator{},
			"multiplication": mul.Generator{},
			"division":       div.Generator{},
		},
	}
}

func (h *SubmitHandler) HandleSubmit(c echo.Context) error {
	if c.Request().Method != http.MethodPost {
		return c.NoContent(http.StatusMethodNotAllowed)
	}
	if err := c.Request().ParseForm(); err != nil {
		return c.String(http.StatusBadRequest, "Unable to parse form")
	}

	operator := c.Request().FormValue("operator")
	numQuestions, _ := strconv.Atoi(c.Request().FormValue("numQuestions"))
	numOperands, _ := strconv.Atoi(c.Request().FormValue("numOperands"))

	if numQuestions < 1 {
		return c.String(http.StatusBadRequest, "Number of Questions must be at least 1")
	}
	if numOperands < 2 {
		numOperands = 2
	}
	if numOperands > 3 {
		numOperands = 3
	}

	rawDigits := c.Request().Form["numDigits"]
	if len(rawDigits) != numOperands {
		return c.String(http.StatusBadRequest, "Provide digits for each operand ("+strconv.Itoa(numOperands)+")")
	}
	digits := make([]int, numOperands)
	for i, d := range rawDigits {
		v, err := strconv.Atoi(d)
		if err != nil || v < 1 {
			return c.String(http.StatusBadRequest, "Digits must be >= 1")
		}
		digits[i] = v
	}

	gen, ok := h.gens[operator]
	if !ok {
		return c.String(http.StatusBadRequest, "Invalid or missing operator")
	}

	var sb strings.Builder
	sb.WriteString("<html><body><h1>Generated ")
	sb.WriteString(strings.Title(operator))
	sb.WriteString(" Problems</h1><ul>")

	seen := make(map[string]bool)
	count, attempts, maxAttempts := 0, 0, numQuestions*10
	for count < numQuestions && attempts < maxAttempts {
		problem := gen.Generate(h.rng, digits)
		if !seen[problem] {
			seen[problem] = true
			sb.WriteString("<li>")
			sb.WriteString(problem)
			sb.WriteString("</li>")
			count++
		}
		attempts++
	}

	sb.WriteString("</ul></body></html>")
	return c.HTML(http.StatusOK, sb.String())
}


