package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type VideoSettings struct {
	ExternalRtspEnabled string `json:"externalRtspEnabled"`
	RtspURL             string `json:"rtspUrl"`
}

func GetVideoSettings(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/video/settings",  serial)
	var datamodel = VideoSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutVideoSettings(serial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/camera/video/settings",  serial)
	var datamodel = VideoSettings{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}