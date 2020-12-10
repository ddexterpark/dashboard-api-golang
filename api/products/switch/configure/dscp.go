package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type DscpToCosMappings struct {
	Mappings []struct {
		Dscp  int    `json:"dscp"`
		Cos   int    `json:"cos"`
		Title string `json:"title"`
	} `json:"mappings"`
}

func GetDscpToCosMappings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/dscpToCosMappings",
		api.BaseUrl(), networkId)
	var datamodel = DscpToCosMappings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutDscpToCosMappings(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/switch/dscpToCosMappings",
		api.BaseUrl(), networkId)
	var datamodel = DscpToCosMappings{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}