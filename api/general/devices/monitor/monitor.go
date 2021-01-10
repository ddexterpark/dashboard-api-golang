package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)


type DeviceClients []struct {
	Usage struct {
		Sent string `json:"sent"`
		Recv string `json:"recv"`
	} `json:"usage"`
	ID           string      `json:"id"`
	Description  string      `json:"description"`
	Mac          string      `json:"mac"`
	IP           string      `json:"ip"`
	User         string      `json:"user"`
	Vlan         string         `json:"vlan"`
	Switchport   interface{} `json:"switchport"`
	MdnsName     string      `json:"mdnsName"`
	DhcpHostname string      `json:"dhcpHostname"`
}


type LLdpCdp struct {
	SourceMac string `json:"sourceMac"`
	Ports     struct {
		Num8 struct {
			Cdp struct {
				DeviceID   string `json:"deviceId"`
				PortID     string `json:"portId"`
				Address    string `json:"address"`
				SourcePort string `json:"sourcePort"`
			} `json:"cdp"`
		} `json:"8"`
		Num12 struct {
			Cdp struct {
				DeviceID   string `json:"deviceId"`
				PortID     string `json:"portId"`
				Address    string `json:"address"`
				SourcePort string `json:"sourcePort"`
			} `json:"cdp"`
			Lldp struct {
				SystemName        string `json:"systemName"`
				PortID            string `json:"portId"`
				ManagementAddress string `json:"managementAddress"`
				SourcePort        string `json:"sourcePort"`
			} `json:"lldp"`
		} `json:"12"`
	} `json:"ports"`
}

type UplinkLoss []struct {
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	LossPercent string       `json:"lossPercent"`
	LatencyMs   string       `json:"latencyMs"`
}

// Return A Devices Clients
func GetClients(serial, t0, timespan string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/clients",  serial)
	var datamodel = DeviceClients{}
	// Parameters for Request URL
	var parameters = map[string]string{
		"t0": t0,
		"timespan": timespan}
	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List LLDP and CDP information for a device
func GetLLdpCdp(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/lldpCdp",  serial)
	var datamodel = LLdpCdp{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

//  Get the uplink loss percentage and latency in milliseconds for a wired network device.
func GetLossAndLatencyHistory(serial, t0, t1, timespan, resolution, uplink, ip string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/lossAndLatencyHistory",  serial)
	var datamodel = UplinkLoss{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":         t0,
		"t1":         t1,
		"timespan":   timespan,
		"resolution": resolution,
		"uplink":     uplink,
		"ip":         ip,
	}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions

}
