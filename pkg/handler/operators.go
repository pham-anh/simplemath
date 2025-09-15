package handler

type Operator string

const (
    OperatorAddition       Operator = "addition"
    OperatorSubtraction    Operator = "subtraction"
    OperatorMultiplication Operator = "multiplication"
    OperatorDivision       Operator = "division"
)

func (o Operator) String() string { return string(o) }

func (o Operator) Symbol() string {
    switch o {
    case OperatorAddition:
        return "+"
    case OperatorSubtraction:
        return "-"
    case OperatorMultiplication:
        return "×"
    case OperatorDivision:
        return "÷"
    default:
        return ""
    }
}


