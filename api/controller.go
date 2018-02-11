package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
)

var ErrBadCode = errors.New("unexpected response code from server")

type Status struct {
	APIVersion int `json:"apiVersion"`
	Clock      int `json:"clock"`
}

type Controller struct {
	BaseURL   url.URL
	AuthToken string
}

func (c *Controller) getEndpointURL(endpoint string) string {
	newPath := path.Join(c.BaseURL.Path, endpoint)
	query := url.Values{}
	query.Add("auth", c.AuthToken)

	finalURL := c.BaseURL
	finalURL.Path = newPath
	finalURL.RawQuery = query.Encode()

	return finalURL.String()
}

func (c *Controller) GetStatus() (*Status, error) {
	endpoint := c.getEndpointURL("/controller")

	resp, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrBadCode
	}

	decoder := json.NewDecoder(resp.Body)
	var decodedValue Status
	err = decoder.Decode(&decodedValue)

	if err != nil {
		return nil, err
	}

	return &decodedValue, nil
}
