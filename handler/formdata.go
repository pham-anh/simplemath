package handler

import (
	"errors"
	"fmt"

	"simplemath/i18n"
	"simplemath/operator"

	"github.com/labstack/echo/v4"
)

type FormData struct {
	Operator     string `form:"operator"`
	NumQuestions int    `form:"numQuestions"`
	NumOperands  int    `form:"numOperands"`
	Digits       []int  `form:"numDigits"`
	TwoSided     bool   `form:"twoSided"`
	Language     string `form:"language"`
}

func (f *FormData) validate() error {
	trans := i18n.NewLocalizer(f.Language)
	op := operator.Operator(f.Operator)
	if f.Operator == "" || (op != operator.Addition && op != operator.Subtraction && op != operator.Multiplication && op != operator.Division) {
		return errors.New(trans.T("error.invalidOperator"))
	}
	if f.NumQuestions < 1 {
		return errors.New(trans.T("error.minQuestions"))
	}
	if f.NumOperands < 2 {
		f.NumOperands = 2
	}
	if f.NumOperands > 3 {
		f.NumOperands = 3
	}
	if len(f.Digits) != f.NumOperands {
		return fmt.Errorf(trans.T("error.digitsCount"), f.NumOperands)
	}
	for _, d := range f.Digits {
		if d < 1 {
			return errors.New(trans.T("error.minDigits"))
		}
		if d > 2 {
			return errors.New(trans.T("error.maxDigits"))
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
