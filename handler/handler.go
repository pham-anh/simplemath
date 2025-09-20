package handler

import (
	"math/rand"
	"net/http"
	"net/url"
	"text/template"

	"simplemath/gen"
	"simplemath/operator"

	"github.com/labstack/echo/v4"
)

type Item struct {
	Emoji string
	Text  string
}

type SubmitHandler struct {
	rng *rand.Rand
}

type Result struct {
	Operator    string
	ProblemSets [][]Item
}

func NewSubmitHandler(r *rand.Rand) *SubmitHandler { return &SubmitHandler{rng: r} }

func (h *SubmitHandler) HandleSubmit(c echo.Context) error {
	f, err := FormDataFromRequest(c)
	if err != nil {
		c.SetCookie(&http.Cookie{
			Name:     "flash_error",
			Value:    url.QueryEscape(err.Error()),
			Path:     "/",
			MaxAge:   5,
			HttpOnly: false,
		})
		return c.Redirect(303, "/")
	}

	var sets [][]Item
	sets = append(sets, h.generate(f))
	if f.TwoSided {
		sets = append(sets, h.generate(f))
	}

	// Load and execute the template.
	tpl, err := template.ParseFiles("statics/result.html")
	if err != nil {
		return err
	}

	result := Result{
		Operator:    f.Operator,
		ProblemSets: sets,
	}

	_ = tpl.Execute(c.Response().Writer, result)
	return nil
}

func (h *SubmitHandler) generate(f *FormData) []Item {
	sym := operator.Operator(f.Operator).Symbol()
	seen := map[string]bool{}
	count, attempts, maxAttempts := 0, 0, f.NumQuestions*10
	var items []Item
	for count < f.NumQuestions && attempts < maxAttempts {
		ops := make([]int, f.NumOperands)
		for i := 0; i < f.NumOperands; i++ {
			ops[i] = gen.RandomWithDigits(h.rng, f.Digits[i])
		}
		problem := gen.JoinOperands(ops, sym)
		if !seen[problem] {
			seen[problem] = true
			items = append(items, Item{
				Emoji: getRandomEmoji(),
				Text:  problem,
			})
			count++
		}
		attempts++
	}
	return items
}

// Get a random emoji from the emojis slice.
func getRandomEmoji() string {
	// Select a random index
	randomIndex := rand.Intn(len(emojis))
	// Return the emoji at the random index
	return emojis[randomIndex]
}
