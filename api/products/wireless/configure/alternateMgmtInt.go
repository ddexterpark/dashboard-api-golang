package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type AlternateManagementInterface struct {
	Enabled      string `json:"enabled"`
	VlanID       string      `json:"vlanId"`
	Protocols    []string `json:"protocols"`
	AccessPoints []struct {
		Serial                string `json:"serial"`
		AlternateManagementIP string `json:"alternateManagementIp"`
		SubnetMask            string `json:"subnetMask"`
		Gateway               string `json:"gateway"`
		DNS1                  string `json:"dns1"`
		DNS2                  string `json:"dns2"`
	} `json:"accessPoints"`
}

func GetAlternateManagementInterface(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/alternateManagementInterface",
		 networkId)
	var datamodel = AlternateManagementInterface{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutAlternateManagementInterface(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/alternateManagementInterface",
		 networkId)
	var datamodel = AlternateManagementInterface{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}