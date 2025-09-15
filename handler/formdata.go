package handler

import (
	"fmt"

	"simplemath/operator"

	"github.com/labstack/echo/v4"
)

type FormData struct {
    Operator     string `form:"operator"`
    NumQuestions int    `form:"numQuestions"`
    NumOperands  int    `form:"numOperands"`
    Digits       []int  `form:"numDigits"`
}

func (f *FormData) validate() error {
    op := operator.Operator(f.Operator)
    if f.Operator == "" || (op != operator.Addition && op != operator.Subtraction && op != operator.Multiplication && op != operator.Division) {
        return fmt.Errorf("Invalid or missing operator")
    }
    if f.NumQuestions < 1 {
        return fmt.Errorf("Number of Questions must be at least 1")
    }
    if f.NumOperands < 2 {
        f.NumOperands = 2
    }
    if f.NumOperands > 3 {
        f.NumOperands = 3
    }
    if len(f.Digits) != f.NumOperands {
        return fmt.Errorf("Provide digits for each operand (%d)", f.NumOperands)
    }
    for _, d := range f.Digits {
        if d < 1 {
            return fmt.Errorf("Digits must be >= 1")
        }
    }
    return nil
}

func FormDataFromRequest(c echo.Context) (*FormData, error) {
    var f FormData
    if err := c.Bind(&f); err != nil {
        return nil, err
    }
    if err := f.validate(); err != nil {
        return nil, err
    }
    return &f, nil
}


