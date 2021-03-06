package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type Packets []struct {
	PortID  string `json:"portId"`
	Packets []struct {
		Desc       string `json:"desc"`
		Total      int    `json:"total"`
		Sent       int    `json:"sent"`
		Recv       int    `json:"recv"`
		RatePerSec struct {
			Total int `json:"total"`
			Sent  int `json:"sent"`
			Recv  int `json:"recv"`
		} `json:"ratePerSec"`
	} `json:"packets"`
}

type Statuses []struct {
	PortID    string   `json:"portId"`
	Enabled   bool     `json:"enabled"`
	Status    string   `json:"status"`
	Errors    []string `json:"errors"`
	Warnings  []string `json:"warnings"`
	Speed     string   `json:"speed"`
	Duplex    string   `json:"duplex"`
	UsageInKb struct {
		Total int `json:"total"`
		Sent  int `json:"sent"`
		Recv  int `json:"recv"`
	} `json:"usageInKb"`
	Cdp struct {
		SystemName          string `json:"systemName"`
		Platform            string `json:"platform"`
		DeviceID            string `json:"deviceId"`
		PortID              string `json:"portId"`
		NativeVlan          int    `json:"nativeVlan"`
		Address             string `json:"address"`
		ManagementAddress   string `json:"managementAddress"`
		Version             string `json:"version"`
		VtpManagementDomain string `json:"vtpManagementDomain"`
		Capabilities        string `json:"capabilities"`
	} `json:"cdp"`
	Lldp struct {
		SystemName         string `json:"systemName"`
		SystemDescription  string `json:"systemDescription"`
		PortID             string `json:"portId"`
		PortDescription    string `json:"portDescription"`
		ChassisID          string `json:"chassisId"`
		ManagementVlan     int    `json:"managementVlan"`
		PortVlan           int    `json:"portVlan"`
		ManagementAddress  string `json:"managementAddress"`
		SystemCapabilities string `json:"systemCapabilities"`
	} `json:"lldp"`
	ClientCount    int     `json:"clientCount"`
	PowerUsageInWh float64 `json:"powerUsageInWh"`
	TrafficInKbps  struct {
		Total float64 `json:"total"`
		Sent  float64 `json:"sent"`
		Recv  int     `json:"recv"`
	} `json:"trafficInKbps"`
}

func GetPacketCounters(serial, t0, timespan string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/ports/statuses/packets",
		 serial)
	var datamodel = Packets{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":       t0,
		"timespan": timespan}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetStatuses(serial, t0, timespan string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/ports/statuses",
		 serial)
	var datamodel = Statuses{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":       t0,
		"timespan": timespan}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
