package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type VpnStats []struct {
	NetworkID      string `json:"networkId"`
	NetworkName    string `json:"networkName"`
	MerakiVpnPeers []struct {
		NetworkID    string `json:"networkId"`
		NetworkName  string `json:"networkName"`
		UsageSummary struct {
			ReceivedInKilobytes string `json:"receivedInKilobytes"`
			SentInKilobytes     string `json:"sentInKilobytes"`
		} `json:"usageSummary"`
		LatencySummaries []struct {
			SenderUplink   string `json:"senderUplink"`
			ReceiverUplink string `json:"receiverUplink"`
			AvgLatencyMs   string `json:"avgLatencyMs"`
			MinLatencyMs   string `json:"minLatencyMs"`
			MaxLatencyMs   string `json:"maxLatencyMs"`
		} `json:"latencySummaries"`
		LossPercentageSummaries []struct {
			SenderUplink      string  `json:"senderUplink"`
			ReceiverUplink    string  `json:"receiverUplink"`
			AvgLossPercentage string     `json:"avgLossPercentage"`
			MinLossPercentage string     `json:"minLossPercentage"`
			MaxLossPercentage float64 `json:"maxLossPercentage"`
		} `json:"lossPercentageSummaries"`
		JitterSummaries []struct {
			SenderUplink   string  `json:"senderUplink"`
			ReceiverUplink string  `json:"receiverUplink"`
			AvgJitter      float64 `json:"avgJitter"`
			MinJitter      string     `json:"minJitter"`
			MaxJitter      float64 `json:"maxJitter"`
		} `json:"jitterSummaries"`
		MosSummaries []struct {
			SenderUplink   string  `json:"senderUplink"`
			ReceiverUplink string  `json:"receiverUplink"`
			AvgMos         float64 `json:"avgMos"`
			MinMos         string     `json:"minMos"`
			MaxMos         float64 `json:"maxMos"`
		} `json:"mosSummaries"`
	} `json:"merakiVpnPeers"`
}

type VPNStatus []struct {
	NetworkID    string `json:"networkId"`
	NetworkName  string `json:"networkName"`
	DeviceSerial string `json:"deviceSerial"`
	DeviceStatus string `json:"deviceStatus"`
	Uplinks      []struct {
		Interface string `json:"interface"`
		PublicIP  string `json:"publicIp"`
	} `json:"uplinks"`
	VpnMode         string `json:"vpnMode"`
	ExportedSubnets []struct {
		Subnet string `json:"subnet"`
		Name   string `json:"name"`
	} `json:"exportedSubnets"`
	MerakiVpnPeers []struct {
		NetworkID    string `json:"networkId"`
		NetworkName  string `json:"networkName"`
		Reachability string `json:"reachability"`
	} `json:"merakiVpnPeers"`
	ThirdPartyVpnPeers []struct {
		Name         string `json:"name"`
		PublicIP     string `json:"publicIp"`
		Reachability string `json:"reachability"`
	} `json:"thirdPartyVpnPeers"`
}

func GetVpnStats(organizationId, t0, t1, timespan, perPage, startingAfter, endingBefore,
	networkIds string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/appliance/vpn/stats",  organizationId)
	var datamodel = VpnStats{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":            t0,
		"t1":            t1,
		"timespan":      timespan,
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore,
		"networkIds":    networkIds}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetVPNStatus(organizationId, perPage, startingAfter, endingBefore,
	networkIds string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/appliance/vpn/statuses",  organizationId)
	var datamodel = VPNStatus{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore,
		"networkIds":    networkIds}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
