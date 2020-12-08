package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type BypassActivationLockAttempts struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Data   struct {
		Num2090938209 struct {
			Success bool     `json:"success"`
			Errors  []string `json:"errors"`
		} `json:"2090938209"`
		Num38290139892 struct {
			Success bool `json:"success"`
		} `json:"38290139892"`
	} `json:"data"`
}

// Bypass Activation Lock Attempt Status
func GetBypassActivationLockAttempts(networkId, attemptId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/bypassActivationLockAttempts/%s",
		api.BaseUrl(), networkId, attemptId)

	var datamodel = BypassActivationLockAttempts{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostBypassActivationLockAttempts(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/bypassActivationLockAttempts",
		api.BaseUrl(), networkId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = BypassActivationLockAttempts{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}