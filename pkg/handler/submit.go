package handler

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"simplemath/pkg/internal/gen"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type SubmitHandler struct {
	rng *rand.Rand
}

func NewSubmitHandler(r *rand.Rand) *SubmitHandler { return &SubmitHandler{rng: r} }

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

	// Operator symbol for display
	opSymbol := map[string]string{
		"addition":       "+",
		"subtraction":    "-",
		"multiplication": "×",
		"division":       "÷",
	}[operator]
	if opSymbol == "" {
		return c.String(http.StatusBadRequest, "Invalid or missing operator")
	}

	var sb strings.Builder
	sb.WriteString("<html><body><h1>Generated ")
	title := cases.Title(language.Und)
	sb.WriteString(title.String(operator))
	sb.WriteString(" Problems</h1><ul>")

	seen := make(map[string]bool)
	count, attempts, maxAttempts := 0, 0, numQuestions*10
	for count < numQuestions && attempts < maxAttempts {
		ops := make([]int, numOperands)
		for i := 0; i < numOperands; i++ {
			ops[i] = gen.RandomWithDigits(h.rng, digits[i])
		}
		problem := gen.JoinOperands(ops, opSymbol)
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
