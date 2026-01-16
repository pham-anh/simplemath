package i18n

import (
	"embed"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

//go:embed *.yaml
var translations embed.FS

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	// Load English translations
	if _, err := bundle.LoadMessageFileFS(translations, "en.yaml"); err != nil {
		panic(fmt.Sprintf("Failed to load English translations: %v", err))
	}

	// Load Vietnamese translations
	if _, err := bundle.LoadMessageFileFS(translations, "vi.yaml"); err != nil {
		panic(fmt.Sprintf("Failed to load Vietnamese translations: %v", err))
	}
}

// Localizer is a wrapper around i18n.Localizer for template usage
type Localizer struct {
	localizer *i18n.Localizer
}

// NewLocalizer creates a new localizer for the given language
func NewLocalizer(lang string) *Localizer {
	var tag language.Tag
	switch lang {
	case "vi":
		tag = language.Vietnamese
	default:
		tag = language.English
	}

	return &Localizer{
		localizer: i18n.NewLocalizer(bundle, tag.String()),
	}
}

// T translates a message with the given ID
func (l *Localizer) T(messageID string, templateData ...map[string]interface{}) string {
	config := &i18n.LocalizeConfig{
		MessageID: messageID,
	}

	if len(templateData) > 0 {
		config.TemplateData = templateData[0]
	}

	message, err := l.localizer.Localize(config)
	if err != nil {
		// Return the message ID if translation is not found
		return messageID
	}

	return message
}

// For compatibility with existing code
func GetBundle() *i18n.Bundle {
	return bundle
}
