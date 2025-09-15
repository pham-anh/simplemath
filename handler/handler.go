package handler

import (
	"math/rand"
	"strings"

	"simplemath/gen"
	"simplemath/operator"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type SubmitHandler struct {
	rng *rand.Rand
}

func NewSubmitHandler(r *rand.Rand) *SubmitHandler { return &SubmitHandler{rng: r} }

func (h *SubmitHandler) HandleSubmit(c echo.Context) error {
	form, err := FormDataFromRequest(c)
	if err != nil {
		return c.String(400, err.Error())
	}

	sym := operator.Operator(form.Operator).Symbol()
	var sb strings.Builder
	sb.WriteString("<html><head><style>@media print { .no-print { display:none } body{ margin:12mm } }</style></head><body><h1>Generated ")
	title := cases.Title(language.Und)
	sb.WriteString(title.String(form.Operator))
	sb.WriteString(" Problems</h1><div class=\"no-print\"><button onclick=\"window.print()\">Print</button></div><ol>")

	seen := map[string]bool{}
	count, attempts, maxAttempts := 0, 0, form.NumQuestions*10
	for count < form.NumQuestions && attempts < maxAttempts {
		ops := make([]int, form.NumOperands)
		for i := 0; i < form.NumOperands; i++ {
			ops[i] = gen.RandomWithDigits(h.rng, form.Digits[i])
		}
		problem := gen.JoinOperands(ops, sym)
		if !seen[problem] {
			seen[problem] = true
			sb.WriteString("<li>")
			sb.WriteString(problem)
			sb.WriteString("</li>")
			count++
		}
		attempts++
	}

	sb.WriteString("</ol></body></html>")
	return c.HTML(200, sb.String())
}


