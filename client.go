package fullstory

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// BaseURL is the base URL for the fullstory.com API.
const BaseURL = "https://fullstory.com/api/v1"

var _ error = StatusError{}

// StatusError is returned when the HTTP request succeeds, but the response status
// does not equal http.StatusOK.
type StatusError struct {
	Status     string
	StatusCode int
	Body       []byte // Data from response body.
}

func (e StatusError) Error() string {
	return fmt.Sprintf("fullstory: response error: %d %s", e.StatusCode, e.Status)
}

// Client represents a HTTP client for making requests to the FullStory API.
type Client struct {
	HTTPClient *http.Client
	Config     Config
}

// Config is configuration for Client.
type Config struct {
	APIToken string
}

// NewClient returns a Client initialized with http.DefaultClient and the
// supplied Config.
func NewClient(cfg Config) *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		Config:     cfg,
	}
}

// doReq performs the supplied HTTP request and returns the data in the response.
// Necessary authentication headers are added to the request if not already set.
//
// If the error is nil, the caller is responsible for closing the returned data.
func (c *Client) doReq(req *http.Request) (io.ReadCloser, error) {
	req.Header.Set("Authorization", "Basic "+c.Config.APIToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()
		return nil, StatusError{
			Body:       b,
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
		}

	}

	return resp.Body, nil
}
