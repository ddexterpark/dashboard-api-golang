package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type VPPAccounts []struct {
	VPPAccount
}

type VPPAccount struct {
	ID              string `json:"id"`
	VppServiceToken string `json:"vppServiceToken"`
}

func GetVPPAccount(organizationId, vppAccountId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/sm/vppAccounts/%s",
		 organizationId, vppAccountId)
	var datamodel = VPPAccount{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetVPPAccounts(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/sm/vppAccounts",
		 organizationId)
	var datamodel = VPPAccounts{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
