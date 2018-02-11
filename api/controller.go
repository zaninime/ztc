package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

var ErrBadCode = errors.New("unexpected response code from server")

type EpochMilliSeconds struct{ *time.Time }

func (t *EpochMilliSeconds) UnmarshalJSON(jsonValue []byte) error {
	ms, err := strconv.ParseInt(string(jsonValue), 0, 64)

	if err != nil {
		return err
	}

	time := time.Unix(0, ms*1000000)

	t.Time = &time
	return nil
}

type Status2 struct {
	APIVersion int               `json:"apiVersion"`
	Clock      EpochMilliSeconds `json:"clock"`
}

type Status struct {
	Address string            `json:"address"`
	Clock   EpochMilliSeconds `json:"clock"`
	Cluster interface{}       `json:"cluster"`
	Config  struct {
		Physical interface{} `json:"physical"`
		Settings struct {
			PortMappingEnabled    bool   `json:"portMappingEnabled"`
			PrimaryPort           int    `json:"primaryPort"`
			SoftwareUpdate        string `json:"softwareUpdate"`
			SoftwareUpdateChannel string `json:"softwareUpdateChannel"`
		} `json:"settings"`
	} `json:"config"`
	Online               bool   `json:"online"`
	PlanetWorldID        int    `json:"planetWorldId"`
	PlanetWorldTimestamp int64  `json:"planetWorldTimestamp"`
	PublicIdentity       string `json:"publicIdentity"`
	TCPFallbackActive    bool   `json:"tcpFallbackActive"`
	Version              string `json:"version"`
	VersionBuild         int    `json:"versionBuild"`
	VersionMajor         int    `json:"versionMajor"`
	VersionMinor         int    `json:"versionMinor"`
	VersionRev           int    `json:"versionRev"`
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
	endpoint := c.getEndpointURL("/status")

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
