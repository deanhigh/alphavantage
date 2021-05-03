package alphavantage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	defaultBaseURL = "https://www.alphavantage.co/query"
)

// Client common struct encapsultating AV client operations
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client
	// Base URL for API requests. Defaults to the public AlphaVantage APIs
	baseURL *url.URL

	apiKey string

	FundamentalService *FundamentalService
}

// NewClient new client default key from env variable AV_API_KEY
func NewClient() (*Client, error) {
	apikey := os.Getenv("AV_API_KEY")
	if len(apikey) == 0 {
		return nil, fmt.Errorf("Configure an alphavantage API key in AV_API_KEY environment variable")
	}

	return NewClientWithKey(apikey)
}

// NewClientWithKey given an apikey will create a client struct
func NewClientWithKey(apiKey string) (*Client, error) {
	c := &Client{
		apiKey: apiKey,
		client: http.DefaultClient}

	err := c.setBaseURL(defaultBaseURL)

	if err != nil {
		return nil, err
	}

	c.FundamentalService = &FundamentalService{client: c}

	return c, nil
}

// setBaseURL sets the base URL for API requests to a custom endpoint.
func (c *Client) setBaseURL(urlStr string) error {
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	// Update the base URL of the client.
	c.baseURL = baseURL

	return nil
}

// NewQuery creates an API request.
func (c *Client) NewQuery(params map[string]string) (*http.Request, error) {
	u := *c.baseURL

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}

	for key, value := range params {
		q.Add(key, value)
	}
	q.Add("apikey", c.apiKey)

	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

type errorResponse struct {
	ErrorMessage string `json:"Error Message,omitempty"`
	Information  string `json:"Information,omitempty"`
}

// Do run a request specialised for alphavantage which doesn't return error codes but rather different payloads dependent on your error. Silly!
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	er := errorResponse{}
	err = json.Unmarshal([]byte(b), &er)
	if err != nil {
		return resp, err
	}

	if er.ErrorMessage != "" {
		return resp, fmt.Errorf("ERROR: %s", er.ErrorMessage)
	}

	if er.Information != "" {
		return resp, fmt.Errorf("ERROR INFO: %s", er.Information)
	}

	if v != nil {
		err = json.Unmarshal([]byte(b), &v)
		if err != nil {
			err = fmt.Errorf("%w; %s", err, fmt.Errorf("%s", string([]byte(b))))
		}
	}

	return resp, err
}
