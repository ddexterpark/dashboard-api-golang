package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type MeshStatuses struct {
	Serial                string   `json:"serial"`
	MeshRoute             []string `json:"meshRoute"`
	LatestMeshPerformance struct {
		Mbps            int    `json:"mbps"`
		Metric          int    `json:"metric"`
		UsagePercentage string `json:"usagePercentage"`
	} `json:"latestMeshPerformance"`
}

func GetMeshStatuses(networkId, perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/meshStatuses",
		 networkId)
	var datamodel = MeshStatuses{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
