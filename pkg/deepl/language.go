package deepl

import (
	"fmt"
	"strings"
)

// SourceLang is language of the text to be translated
type SourceLang int

const (
	// SourceLangDE points to the German
	SourceLangDE SourceLang = iota
	// SourceLangEN points to the English
	SourceLangEN
	// SourceLangFR points to the French
	SourceLangFR
	// SourceLangIT points to the Italian
	SourceLangIT
	// SourceLangJA points to the Japanese
	SourceLangJA
	// SourceLangES points to the Spanish
	SourceLangES
	// SourceLangNL points to the Dutch
	SourceLangNL
	// SourceLangPL points to the Polish
	SourceLangPL
	// SourceLangPT points to the Portuguese (all Portuguese varieties mixed)
	SourceLangPT
	// SourceLangRU points to the Russian
	SourceLangRU
	// SourceLangZH points to the Chinese
	SourceLangZH
	// SourceLangUnspecified represents omit
	SourceLangUnspecified
)

func (l SourceLang) String() string {
	switch l {
	case SourceLangDE:
		return "DE"
	case SourceLangEN:
		return "EN"
	case SourceLangFR:
		return "FR"
	case SourceLangIT:
		return "IT"
	case SourceLangJA:
		return "JA"
	case SourceLangES:
		return "ES"
	case SourceLangNL:
		return "NL"
	case SourceLangPL:
		return "PL"
	case SourceLangPT:
		return "PT"
	case SourceLangRU:
		return "RU"
	case SourceLangZH:
		return "ZH"
	default:
		return ""
	}
}

// IsSet checks to see the value is set
func (l SourceLang) IsSet() bool {
	return l != SourceLangUnspecified
}

// SourceLangFrom provides SourceLang from string
func SourceLangFrom(l string) SourceLang {
	switch strings.ToUpper(l) {
	case SourceLangDE.String():
		return SourceLangDE
	case SourceLangEN.String():
		return SourceLangEN
	case SourceLangFR.String():
		return SourceLangFR
	case SourceLangIT.String():
		return SourceLangIT
	case SourceLangJA.String():
		return SourceLangJA
	case SourceLangES.String():
		return SourceLangES
	case SourceLangNL.String():
		return SourceLangNL
	case SourceLangPL.String():
		return SourceLangPL
	case SourceLangPT.String():
		return SourceLangPT
	case SourceLangRU.String():
		return SourceLangRU
	case SourceLangZH.String():
		return SourceLangZH
	default:
		return SourceLangUnspecified
	}
}

// TargetLang is language into which the text should be translated
type TargetLang int

const (
	// TargetLangDE points to German
	TargetLangDE TargetLang = iota
	// TargetLangEN points to English
	TargetLangEN
	// TargetLangFR points to French
	TargetLangFR
	// TargetLangIT points to Italian
	TargetLangIT
	// TargetLangJA points to Japanese
	TargetLangJA
	// TargetLangES points to Spanish
	TargetLangES
	// TargetLangNL points to Dutch
	TargetLangNL
	// TargetLangPL points to Polish
	TargetLangPL
	// TargetLangPT points to Portuguese (all Portuguese varieties excluding Brazilian Portuguese)
	TargetLangPT
	// TargetLangPTBR points to Portuguese (Brazilian)
	TargetLangPTBR
	// TargetLangRU points to Russian
	TargetLangRU
	// TargetLangZH points to Chinese
	TargetLangZH
)

func (l TargetLang) String() string {
	switch l {
	case TargetLangDE:
		return "DE"
	case TargetLangFR:
		return "FR"
	case TargetLangIT:
		return "IT"
	case TargetLangJA:
		return "JA"
	case TargetLangES:
		return "ES"
	case TargetLangNL:
		return "NL"
	case TargetLangPL:
		return "PL"
	case TargetLangPT:
		return "PT"
	case TargetLangPTBR:
		return "PT-BR"
	case TargetLangRU:
		return "RU"
	case TargetLangZH:
		return "ZH"
	case TargetLangEN:
		return "EN"
	default:
		return ""
	}
}

// TargetLangFrom provides TargetLang from string
func TargetLangFrom(l string) (TargetLang, error) {
	var (
		t TargetLang
		e error
	)

	switch strings.ToUpper(l) {
	case TargetLangDE.String():
		t = TargetLangDE
	case TargetLangEN.String():
		t = TargetLangEN
	case TargetLangFR.String():
		t = TargetLangFR
	case TargetLangIT.String():
		t = TargetLangIT
	case TargetLangJA.String():
		t = TargetLangJA
	case TargetLangES.String():
		t = TargetLangES
	case TargetLangNL.String():
		t = TargetLangNL
	case TargetLangPL.String():
		t = TargetLangPL
	case TargetLangPT.String():
		t = TargetLangPT
	case TargetLangPTBR.String():
		t = TargetLangPTBR
	case TargetLangRU.String():
		t = TargetLangRU
	case TargetLangZH.String():
		t = TargetLangZH
	default:
		e = fmt.Errorf("undefined target language")
	}

	return t, e
}
