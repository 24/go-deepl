package translation

import (
	"errors"
	"fmt"
	"github.com/tarxzfv/go-deepl/pkg/deepl"
	"github.com/urfave/cli/v2"
)

// Translator translate the text
type Translator interface {
	ProcessWithCLI(*cli.Context) (string, error)
}

// NewTranslator provides new instance that implemented Translator
func NewTranslator(deeplClient deepl.ClientV2) Translator {
	return &translator{deeplClient}
}

type translator struct {
	deepL deepl.ClientV2
}

func (t *translator) ProcessWithCLI(ctx *cli.Context) (string, error) {
	text := ctx.Args().Get(0)

	targetLang, err := deepl.TargetLangFrom(ctx.Value("target_lang").(string))
	if err != nil {
		return "", err
	}
	sourceLang := deepl.SourceLangFrom(ctx.Value("source_lang").(string))

	resp, err := t.deepL.Translate(sourceLang, targetLang, text)
	if err != nil {
		msg := "failed to translate: " + err.Error()

		if resp != nil && resp.ErrorMessage != "" {
			msg = fmt.Sprintf("%s (%s)", msg, resp.ErrorMessage)
		}

		return "", errors.New(msg)
	}

	return resp.Translations[0].Text, nil
}
