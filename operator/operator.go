package operator

type Operator string

const (
    Addition       Operator = "addition"
    Subtraction    Operator = "subtraction"
    Multiplication Operator = "multiplication"
    Division       Operator = "division"
)

func (o Operator) String() string { return string(o) }

func (o Operator) Symbol() string {
    switch o {
    case Addition:
        return "+"
    case Subtraction:
        return "-"
    case Multiplication:
        return "×"
    case Division:
        return "÷"
    default:
        return ""
    }
}


