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
	Operator string
	Items    []Item
}

func NewSubmitHandler(r *rand.Rand) *SubmitHandler { return &SubmitHandler{rng: r} }

func (h *SubmitHandler) HandleSubmit(c echo.Context) error {
	form, err := FormDataFromRequest(c)
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

	sym := operator.Operator(form.Operator).Symbol()
	seen := map[string]bool{}
	count, attempts, maxAttempts := 0, 0, form.NumQuestions*10

	var items []Item
	for count < form.NumQuestions && attempts < maxAttempts {
		ops := make([]int, form.NumOperands)
		for i := 0; i < form.NumOperands; i++ {
			ops[i] = gen.RandomWithDigits(h.rng, form.Digits[i])
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

	// Load and execute the template.
	tpl, err := template.ParseFiles("statics/result.html")
	if err != nil {
		return err
	}
	_ = tpl.Execute(c.Response().Writer, Result{
		Operator: form.Operator,
		Items:    items,
	})
	return nil
}

// Get a random emoji from the emojis slice.
func getRandomEmoji() string {
	// Select a random index
	randomIndex := rand.Intn(len(emojis))
	// Return the emoji at the random index
	return emojis[randomIndex]
}

// Large slice of emojis covering the requested categories.
var emojis = []string{
	// Animals & Nature
	"🐵", "🐒", "🦍", "🦧", "🐶", "🐕", "🦮", "🐩", "🐺", "🦊",
	"🦝", "🐱", "🐈", "🦁", "🐯", "🐅", "🐆", "🐴", "🐎", "🦌",
	"🐮", "🐂", "🐃", "🐄", "🐷", "🐖", "🐗", "🐽", "🐏", "🐑",
	"🐐", "🐪", "🐫", "🦙", "🦒", "🐘", "🦣", "🦏", "🦛", "🐭",
	"🐁", "🐀", "🐹", "🐰", "🐇", "🐿️", "🦫", "🦡", "🦔", "🦦",
	"🦇", "🐻", "🐻‍❄️", "🐨", "🐼", "🦥", "🐾", "🦃", "🐔", "🐓",
	"🐣", "🐤", "🐥", "🐦", "🐧", "🕊️", "🦅", "🦆", "🦢", "🦉",
	"🦤", "🦩", "🦜", "🐸", "🐊", "🐢", "🦎", "🐍", "🐲", "🐉",
	"🦕", "🦖", "🐳", "🐋", "🐬", "🦭", "🐠", "🐟", "🐡", "🦈",
	"🐙", "🐚", "🐌", "🦋", "🐛", "🐜", "🐝", "🪲", "🐞", "🦗",
	"🪳", "🕷️", "🕸️", "🦂", "🦟", "🪰", "🪱", "🦠", "💐", "🌸",
	"💮", "🪷", "🪻", "🌷", "🌹", "🥀", "🌺", "🌻", "🌼", "🍂",
	"🍁", "🌾", "🌿", "🌱", "🌲", "🌳", "🌴", "🌵", "🪴", "🪹",
	"🪺", "🌰", "🍄", "🌎", "🌍", "🌏", "🌑", "🌒", "🌓", "🌔",
	"🌕", "🌖", "🌗", "🌘", "🌙", "🌚", "🌛", "🌜", "🌝", "🌞",
	"🌟", "🌠", "🌌", "☄️", "🪐", "☀️", "🌡️", "🌤️", "🌥️", "🌦️",
	"☁️", "🌧️", "⛈️", "🌩️", "⚡", "🔥", "💥", "❄️", "🌨️", "☃️",
	"⛄", "🌬️", "💨", "🌪️", "🌫️", "🌈", "☂️", "☔", "💧", "🌊",
	"🪨", "🪵", "🏔️", "⛰️", "🌋", "🗻", "🏕️", "🏞️", "🛣️", "🛤️",
	"🌅", "🌄", "🏙️", "🌉", "🌃", "🌆", "🌇",

	// Food & Drink
	"🍇", "🍈", "🍉", "🍊", "🍋", "🍌", "🍍", "🥭", "🍎", "🍏",
	"🍐", "🍑", "🍒", "🍓", "🫐", "🥝", "🍅", "🫒", "🥥", "🥑",
	"🍆", "🥔", "🥕", "🌽", "🌶️", "🫑", "🥒", "🥬", "🥦", "🧄",
	"🧅", "🥜", "🌰", "🫚", "🫛", "🫘", "🍞", "🥐", "🥖", "🫓",
	"🥨", "🧀", "🥚", "🍳", "🧈", "🥓", "🥩", "🍗", "🍖", "🦴",
	"🌭", "🍔", "🍟", "🍕", "🫔", "🥪", "🫕", "🥙", "🧆", "🌮",
	"🌯", "🫙", "🫛", "🍝", "🍜", "🍲", "🍛", "🍣", "🍱", "🥟",
	"🦪", "🍚", "🍘", "🍙", "🍢", "🍡", "🍧", "🍡", "🍨", "🍦",
	"🍩", "🍪", "🎂", "🍰", "🧁", "🥧", "🍫", "🍬", "🍭", "🍮",
	"🍯", "🍶", "🍼", "🥛", "☕", "🍵", "🫖", "🧋", "🍾", "🍷",
	"🍸", "🍹", "🍻", "🥂", "🥃", "🫗", "🥤", "🧊", "🥄", "🍴",
	"🍽️", "🔪", "🥢", "🧂",

	// Travel & Places
	"🚗", "🚕", "🚙", "🚌", "🚎", "🏎️", "🚓", "🚑", "🚒", "🚐",
	"🛻", "🚚", "🚛", "🚜", "🛵", "🏍️", "🛺", "🚲", "🛴", "🦼",
	"🦽", "🩼", "🛞", "🚨", "🚔", "🚍", "🚖", "🚡", "🚠", "🚟",
	"🚃", "🚄", "🚅", "🚂", "🚆", "🚇", "🚈", "🚝", "🚟", "🛗",
	"✈️", "🛫", "🛬", "🚁", "🚀", "🛸", "🛶", "⛵", "🚤", "🛥️",
	"🛳️", "⛴️", "🚢", "⚓", "🚧", "🚦", "🚥", "⛽", "🚏", "🗺️",
	"🗾", "🧭", "💒", "⛪", "🕌", "🛕", "🕍", "⛩️", "🕋", "🏰",
	"🏯", "🏟️", "🗼", "🗽", "🏠", "🏡", "🏘️", "🛖", "🏚️", "🏢",
	"🏣", "🏤", "🏥", "🏦", "🏨", "🏩", "🏪", "🏫", "🏭", "🏯",
	"🏯", "🏰", "🗼", "🗽", "🏠", "🏡", "🏘️", "🏚️", "🏢", "🏣",
	"🏤", "🏥", "🏦", "🏨", "🏩", "🏪", "🏫", "🏭", "🏯", "🏯",
	"🌃", "🌆", "🌇", "🌉", "🛕", "🕍", "⛩️", "🕋", "🏛️", "🛖",
	"🏞️", "🛣️", "🛤️", "🌅", "🌄", "🏙️", "🌉", "🌃", "🌆", "🌇",
}
