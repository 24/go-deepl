package deepl

import "net/http"

// TranslateResponse represents "Translate text" API response
// see: https://www.deepl.com/docs-api/translating-text#response
type TranslateResponse struct {
	ErrorMessage string `json:"message"`
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	}
}

// MonitoringUsageResponse represents "Monitoring usage" API response
// see: https://www.deepl.com/docs-api/other-functions/monitoring-usage
type MonitoringUsageResponse struct {
	ErrorMessage   string `json:"message"`
	CharacterCount int    `json:"character_count"`
	CharacterLimit int    `json:"character_limit"`
}

// APIError represents about the error
type APIError int

func (e APIError) String() string {
	switch e {
	case http.StatusBadRequest:
		return "Bad request. Please check error message and your parameters."
	case http.StatusForbidden:
		return "Authorization failed. Please supply a valid auth_key parameter."
	case http.StatusNotFound:
		return "The requested resource could not be found."
	case http.StatusRequestEntityTooLarge:
		return "The request size exceeds the limit."
	case http.StatusTooManyRequests:
		return "Too many requests. Please wait and resend your request."
	case 456:
		return "Quota exceeded. The character limit has been reached."
	case http.StatusServiceUnavailable:
		return "Resource currently unavailable. Try again later."
	}

	if 500 <= e {
		return "Internal error"
	}

	return ""
}
