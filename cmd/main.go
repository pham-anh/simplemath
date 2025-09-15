package main

import (
	"math/rand"
	"net/http"
	"strconv"

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

		// Create an instance of FormData
		numQuestions, _ := strconv.Atoi(c.Request().FormValue("numQuestions"))
		numOperands, _ := strconv.Atoi(c.Request().FormValue("numOperands"))

		// Generate a response to display the problems
		response := "<html><body><h1>Generated Addition Problems</h1><ul>"

		for i := 0; i < numQuestions; i++ {
			// Generate a random addition problem
			operands := make([]int, numOperands)
			for j := 0; j < numOperands; j++ {
				operands[j] = rand.Intn(100) // Example: random number between 0 and 99
			}

			// Create the addition problem string
			problem := ""
			for j, operand := range operands {
				if j > 0 {
					problem += " + "
				}
				problem += strconv.Itoa(operand)
			}

			// Add the problem to the response
			response += "<li>" + problem + "</li>"
		}

		// End the HTML response
		response += "</ul></body></html>"

		// Return the response
		return c.HTML(http.StatusOK, response)
	}

	return nil
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
