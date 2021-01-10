package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type Devices []struct {
	Device
}
type Device struct {
	Name           string  `json:"name"`
	Lat            float64 `json:"lat"`
	Lng            float64 `json:"lng"`
	Serial         string  `json:"serial"`
	Mac            string  `json:"mac"`
	Model          string  `json:"model"`
	Address        string  `json:"address"`
	Notes          string  `json:"notes"`
	LanIP          string  `json:"lanIp"`
	Tags           string  `json:"tags"`
	NetworkID      string  `json:"networkId"`
	BeaconIDParams struct {
		UUID  string `json:"uuid"`
		Major string `json:"major"`
		Minor string `json:"minor"`
	} `json:"beaconIdParams"`
	Firmware    string `json:"firmware"`
	FloorPlanID string `json:"floorPlanId"`
}

// List the devices in a network
func GetDevices(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/devices",  networkId)
	var datamodel = Devices{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Claim devices into a network
func PostClaimDevices(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/devices/claim",  networkId)
	var datamodel interface{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Remove a single device
func PostRemoveDevices(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/devices/remove",  networkId)
	var datamodel interface{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}