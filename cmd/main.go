package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// Define a struct to hold the form data
type FormData struct {
	Operator    string
	NumQuestions int
	NumOperands  int
	NumDigits   []int
}

// Handle form submission logic
func handleFormSubmission(c echo.Context) error {
	if c.Request().Method == http.MethodPost {
		// Parse the form data
		err := c.Request().ParseForm()
		if err != nil {
			return c.String(http.StatusBadRequest, "Unable to parse form")
		}

		// Parse inputs
		operator := c.Request().FormValue("operator")
		numQuestions, _ := strconv.Atoi(c.Request().FormValue("numQuestions"))
		numOperands, _ := strconv.Atoi(c.Request().FormValue("numOperands"))

		// Collect numDigits (multiple inputs with same name)
		rawDigits := c.Request().Form["numDigits"]
		digits := make([]int, 0, len(rawDigits))
		for _, d := range rawDigits {
			if d == "" {
				continue
			}
			v, convErr := strconv.Atoi(d)
			if convErr != nil {
				return c.String(http.StatusBadRequest, "Invalid numDigits value")
			}
			digits = append(digits, v)
		}

		// Validate inputs
		if operator == "" || (operator != "addition" && operator != "subtraction" && operator != "multiplication" && operator != "division") {
			return c.String(http.StatusBadRequest, "Invalid or missing operator")
		}
		if numQuestions < 1 {
			return c.String(http.StatusBadRequest, "Number of Questions must be at least 1")
		}
		if numOperands < 2 {
			numOperands = 2
		}
		if numOperands > 3 {
			numOperands = 3
		}
		if len(digits) != numOperands {
			return c.String(http.StatusBadRequest, "Provide digits for each operand ("+strconv.Itoa(numOperands)+")")
		}
		for _, d := range digits {
			if d < 1 {
				return c.String(http.StatusBadRequest, "Digits must be >= 1")
			}
		}

		// Operator symbol for display
		opSymbol := map[string]string{
			"addition":       "+",
			"subtraction":    "-",
			"multiplication": "×",
			"division":       "÷",
		}[operator]

		// Build response with deduplication
		var sb strings.Builder
		sb.WriteString("<html><body>")
		sb.WriteString("<h1>Generated ")
		sb.WriteString(strings.Title(operator))
		sb.WriteString(" Problems</h1><ul>")

		generatedProblems := make(map[string]bool)
		problemCount := 0
		maxAttempts := numQuestions * 10
		attempts := 0

		for problemCount < numQuestions && attempts < maxAttempts {
			operands := make([]int, numOperands)
			for j := 0; j < numOperands; j++ {
				operands[j] = randomWithDigits(digits[j])
			}
			problem := joinOperands(operands, opSymbol)
			if !generatedProblems[problem] {
				generatedProblems[problem] = true
				sb.WriteString("<li>")
				sb.WriteString(problem)
				sb.WriteString("</li>")
				problemCount++
			}
			attempts++
		}

		sb.WriteString("</ul></body></html>")
		return c.HTML(http.StatusOK, sb.String())
	}

	return nil
}

func powerOfTen(exp int) int {
	result := 1
	for range exp {
		result *= 10
	}
	return result
}

func randomWithDigits(d int) int {
	if d <= 1 {
		return rand.Intn(9) + 1
	}
	min := powerOfTen(d - 1)
	span := powerOfTen(d) - min
	return min + rand.Intn(span)
}

func joinOperands(operands []int, symbol string) string {
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

func main() {
	e := echo.New()
	e.POST("/submit", handleFormSubmission)
	// Handle GET request to display the input form
	e.GET("/", func(c echo.Context) error {
		return c.File("statics/index.html")
	})

	// Start the server on port 8080
	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
