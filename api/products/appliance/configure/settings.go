package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type Settings struct {
	ClientTrackingMethod string `json:"clientTrackingMethod"`
}

func GetSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/settings",  networkId)
	var datamodel = Settings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
