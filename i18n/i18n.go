package i18n

type Language string

const (
	Vietnamese Language = "vi"
	English    Language = "en"
)

type Translator struct {
	lang Language
}

func NewTranslator(lang string) *Translator {
	l := Language(lang)
	if l != English && l != Vietnamese {
		l = English // Default to English
	}
	return &Translator{lang: l}
}

func (t *Translator) T(key string) string {
	if t.lang == English {
		return translations[English][key]
	}
	return translations[Vietnamese][key]
}

var translations = map[Language]map[string]string{
	Vietnamese: {
		// Page titles and headers
		"page.title":        "Bài Tập Toán",
		"page.result.title": "Bài Tập Đã Tạo",
		"header.create":     "Tạo Bài Tập Toán",
		"header.result":     "Bài Tập Cộng",

		// Form labels
		"label.operator":     "Chọn Phép Toán:",
		"label.numQuestions": "Số Lượng Câu:",
		"label.numOperands":  "Số Lượng Số Hạng:",
		"label.digits":       "Số Chữ Số Cho Mỗi Số Hạng:",
		"label.twoSided":     "In 2 mặt",

		// Operator names
		"operator.addition":       "Phép Cộng",
		"operator.subtraction":    "Phép Trừ",
		"operator.multiplication": "Phép Nhân",
		"operator.division":       "Phép Chia",

		// Operand labels
		"operand.label": "Số Hạng",

		// Buttons
		"button.create":      "Tạo Bài Tập!",
		"button.back":        "Quay Lại",
		"button.generateNew": "Tạo Bài Tập Mới",
		"button.print":       "In",

		// Validation messages
		"error.invalidOperator": "Phép toán không hợp lệ hoặc bị thiếu",
		"error.minQuestions":    "Số lượng câu hỏi phải ít nhất là 1",
		"error.digitsCount":     "Vui lòng cung cấp số chữ số cho mỗi số hạng (%d)",
		"error.minDigits":       "Số chữ số phải >= 1",
		"error.maxDigits":       "Số chữ số phải <= 2",
	},
	English: {
		// Page titles and headers
		"page.title":        "Math Exercise",
		"page.result.title": "Generated Exercise",
		"header.create":     "Create Math Exercise",
		"header.result":     "Addition Exercise",

		// Form labels
		"label.operator":     "Select Operation:",
		"label.numQuestions": "Number of Questions:",
		"label.numOperands":  "Number of Operands:",
		"label.digits":       "Number of Digits for Each Operand:",
		"label.twoSided":     "Print on Both Sides",

		// Operator names
		"operator.addition":       "Addition",
		"operator.subtraction":    "Subtraction",
		"operator.multiplication": "Multiplication",
		"operator.division":       "Division",

		// Operand labels
		"operand.label": "Operand",

		// Buttons
		"button.create":      "Create Exercise!",
		"button.back":        "Back",
		"button.generateNew": "Generate New Exercise",
		"button.print":       "Print",

		// Validation messages
		"error.invalidOperator": "Invalid or missing operation",
		"error.minQuestions":    "Number of questions must be at least 1",
		"error.digitsCount":     "Please provide number of digits for each operand (%d)",
		"error.minDigits":       "Number of digits must be >= 1",
		"error.maxDigits":       "Number of digits must be <= 2",
	},
}
