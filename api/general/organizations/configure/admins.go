package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
	"time"
)

type Admins []struct {
	Admin
}

type Admin struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	OrgAccess            string `json:"orgAccess"`
	AccountStatus        string `json:"accountStatus"`
	TwoFactorAuthEnabled string      `json:"twoFactorAuthEnabled"`
	HasAPIKey            string      `json:"hasApiKey"`
	LastActive           time.Time `json:"lastActive"`
	Tags                 []struct {
		Tag    string `json:"tag"`
		Access string `json:"access"`
	} `json:"tags"`
	Networks []struct {
		ID     string `json:"id"`
		Access string `json:"access"`
	} `json:"networks"`
	AuthenticationMethod string `json:"authenticationMethod"`
}

// List the dashboard administrators in this organization
func GetAdmins(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/admins", organizationId)
	var datamodel = Admins{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Revoke all access for a dashboard administrator within this organization
func DelAdmin(organizationId, adminId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/admins/%s",  organizationId, adminId)
	var datamodel = Admin{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update an administrator
func PutAdmin(organizationId, adminId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/admins/%s",  organizationId, adminId)
	var datamodel = Admin{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Create a new dashboard administrator
func PostAdmin(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/admins",  organizationId)
	var datamodel = Admin{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}