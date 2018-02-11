package api

import (
	"encoding/json"
	"net"
	"net/http"
)

type EditableMember struct {
	Authorized      bool     `json:"authorized"`
	ActiveBridge    bool     `json:"activeBridge"`
	IPAssignments   []net.IP `json:"ipAssignments"`
	NoAutoAssignIps bool     `json:"noAutoAssignIps"`
}

type Member struct {
	*EditableMember
	Address              string            `json:"address"`
	AuthHistory          []interface{}     `json:"authHistory"`
	Capabilities         []string          `json:"capabilities"`
	Clock                EpochMilliSeconds `json:"clock"`
	CreationTime         EpochMilliSeconds `json:"creationTime"`
	ID                   string            `json:"id"`
	Identity             string            `json:"identity"`
	LastAuthorizedTime   int               `json:"lastAuthorizedTime"`
	LastDeauthorizedTime int               `json:"lastDeauthorizedTime"`
	LastModified         EpochMilliSeconds `json:"lastModified"`
	LastRequestMetaData  string            `json:"lastRequestMetaData"`
	RecentLog            []struct {
		Auth   bool   `json:"auth"`
		AuthBy string `json:"authBy"`
		Ts     int64  `json:"ts"`
		VMajor int    `json:"vMajor"`
		VMinor int    `json:"vMinor"`
		VProto int    `json:"vProto"`
		VRev   int    `json:"vRev"`
	} `json:"recentLog"`
	Revision int      `json:"revision"`
	Tags     []string `json:"tags"`
}

func (c *Controller) GetMemberList(networkID string) ([]string, error) {
	endpoint := c.getEndpointURL("/controller/network/" + networkID + "/member")

	resp, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrBadCode
	}

	decoder := json.NewDecoder(resp.Body)
	var decodedValue map[string]int
	err = decoder.Decode(&decodedValue)

	if err != nil {
		return nil, err
	}

	var keys []string

	for key := range decodedValue {
		keys = append(keys, key)
	}

	return keys, nil
}

func (c *Controller) GetMember(networkID string, nodeID string) (*Member, error) {
	endpoint := c.getEndpointURL("/controller/network/" + networkID + "/member/" + nodeID)

	resp, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrBadCode
	}

	decoder := json.NewDecoder(resp.Body)
	var decodedValue Member
	err = decoder.Decode(&decodedValue)

	if err != nil {
		return nil, err
	}

	return &decodedValue, nil
}
