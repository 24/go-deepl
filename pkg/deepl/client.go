package deepl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// ClientV2 represents DeepL API (v2) endpoint
// see: https://www.deepl.com/docs-api/
type ClientV2 interface {
	Translate(s SourceLang, t TargetLang, text string) (*TranslateResponse, error)
	MonitoringUsage() (*MonitoringUsageResponse, error)
}

// Config is configuration of DeepL API client
type Config struct {
	URL     string
	AuthKey string
}

// NewDeepLClient provides the instance that implemented Client
func NewDeepLClient(c *Config, h *http.Client) (ClientV2, error) {
	_, err := url.Parse(c.URL)
	if err != nil {
		return nil, errors.New("failed to parse url: " + c.URL)
	}
	if c.AuthKey == "" {
		return nil, errors.New("empty authentication key")
	}

	return &client{c, h}, nil
}

type client struct {
	config *Config
	http   *http.Client
}

func (c *client) Translate(s SourceLang, t TargetLang, text string) (*TranslateResponse, error) {
	if text == "" {
		return nil, errors.New("empty text")
	}

	u, err := c.buildURLPath("/v2/translate")
	if err != nil {
		return nil, fmt.Errorf("failed to build URL (%s)", err)
	}
	body := strings.NewReader(c.buildTranslateRequestBody(s, t, text).Encode())

	req, err := http.NewRequest(http.MethodPost, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request (%s)", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	v := new(TranslateResponse)

	resp, err := c.sendRequest(req, v)
	if err != nil {
		return nil, fmt.Errorf("failed to request (%s)", err)
	}

	if resp.StatusCode != http.StatusOK {
		return v, errors.New(APIError(resp.StatusCode).String())
	}

	return v, nil
}

func (c *client) MonitoringUsage() (*MonitoringUsageResponse, error) {
	u, err := c.buildURLPath("/v2/usage")
	if err != nil {
		return nil, fmt.Errorf("failed to build URL (%s)", err)
	}
	body := strings.NewReader(c.buildRequestBody().Encode())

	req, err := http.NewRequest(http.MethodPost, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request (%s)", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	v := new(MonitoringUsageResponse)

	resp, err := c.sendRequest(req, v)
	if err != nil {
		return nil, fmt.Errorf("failed to request (%s)", err)
	}

	if resp.StatusCode != http.StatusOK {
		return v, errors.New(APIError(resp.StatusCode).String())
	}

	return v, nil
}

func (c *client) buildURLPath(v string) (*url.URL, error) {
	u, err := url.Parse(c.config.URL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, v)

	return u, nil
}

// see: https://www.deepl.com/docs-api/translating-text#request
func (c *client) buildTranslateRequestBody(s SourceLang, t TargetLang, text string) url.Values {
	body := c.buildRequestBody()

	body.Add("text", text)
	body.Add("target_lang", t.String())

	if s.IsSet() {
		body.Add("source_lang", s.String())
	}

	return body
}

func (c *client) buildRequestBody() url.Values {
	body := url.Values{}
	body.Add("auth_key", c.config.AuthKey)

	return body
}

func (c *client) sendRequest(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request (%s)", err)
	}
	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return nil, fmt.Errorf("failed to parse response (%s)", err)
	}
	return resp, nil
}
