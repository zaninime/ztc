package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type IPNetString struct {
	*net.IPNet
}

func (ins *IPNetString) UnmarshalJSON(jsonValue []byte) error {
	var rawString string
	err := json.Unmarshal(jsonValue, &rawString)

	if err != nil {
		return err
	}

	_, net, err := net.ParseCIDR(rawString)

	if err != nil {
		return err
	}

	ins.IPNet = net
	return nil
}

func (ins *IPNetString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, ins.String())), nil
}

type Rule struct {
	Type      string `json:"type"`
	Not       bool   `json:"not"`
	Or        bool   `json:"or"`
	EtherType int    `json:"etherType,omitempty"`
}

type V4AssignMode struct {
	Zt bool `json:"zt"`
}

type V6AssignMode struct {
	Zt       bool `json:"zt"`
	Rfc4193  bool `json:"rfc4193"`
	SixPlane bool `json:"6plane"`
}

type Route struct {
	Target IPNetString `json:"target"`
	Via    *net.IP     `json:"via"`
}

type IPAssignmentPool struct {
	IPRangeStart net.IP `json:"ipRangeStart"`
	IPRangeEnd   net.IP `json:"ipRangeEnd"`
}

type EditableNetwork struct {
	IPAssignmentPools []IPAssignmentPool `json:"ipAssignmentPools"`
	MulticastLimit    int                `json:"multicastLimit"`
	Routes            []Route            `json:"routes"`
	Tags              []string           `json:"tags"`
	V4AssignMode      V4AssignMode       `json:"v4AssignMode"`
	V6AssignMode      V6AssignMode       `json:"v6AssignMode"`
	Rules             []Rule             `json:"rules"`
	EnableBroadcast   bool               `json:"enableBroadcast"`
	Name              string             `json:"name"`
	Private           bool               `json:"private"`
}

type Network struct {
	*EditableNetwork
	ID           string            `json:"id"`
	Revision     int               `json:"revision"`
	CreationTime EpochMilliSeconds `json:"creationTime"`
}

func (c *Controller) GetNetwork(networkID string) (*Network, error) {
	endpoint := c.getEndpointURL("/controller/network/" + networkID)

	resp, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrBadCode
	}

	decoder := json.NewDecoder(resp.Body)
	var decodedValue Network
	err = decoder.Decode(&decodedValue)

	if err != nil {
		return nil, err
	}

	return &decodedValue, nil

}

func (c *Controller) EditNetwork(networkID string, config *EditableNetwork) (*Network, error) {
	endpoint := c.getEndpointURL("/controller/network/" + networkID)

	requestBody, err := json.Marshal(config)

	requestBodyReader := bytes.NewBuffer(requestBody)

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(endpoint, "application/json", requestBodyReader)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// duplicate from above
	if resp.StatusCode != http.StatusOK {
		return nil, ErrBadCode
	}

	decoder := json.NewDecoder(resp.Body)
	var decodedValue Network
	err = decoder.Decode(&decodedValue)

	if err != nil {
		return nil, err
	}

	return &decodedValue, nil
	// end duplicate from above
}

func (c *Controller) GetNetworkList() ([]string, error) {
	endpoint := c.getEndpointURL("/controller/network/")

	resp, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrBadCode
	}

	decoder := json.NewDecoder(resp.Body)
	var decodedValue []string
	err = decoder.Decode(&decodedValue)

	if err != nil {
		return nil, err
	}

	return decodedValue, nil
}
