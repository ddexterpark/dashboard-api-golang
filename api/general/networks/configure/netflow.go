package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type NetFlow struct {
	ReportingEnabled bool   `json:"reportingEnabled"`
	CollectorIP      string `json:"collectorIp"`
	CollectorPort    int    `json:"collectorPort"`
}

func GetNetFlow(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/netflow", api.BaseUrl(), networkId)
	var datamodel = NetFlow{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PutNetFlow(networkId, reportingEnabled, collectorIp, collectorPort string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/netflow", api.BaseUrl(), networkId)
	var datamodel = NetFlow{}

	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"reportingEnabled": reportingEnabled,
		"collectorIp": collectorIp,
		"collectorPort": collectorPort}

	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}