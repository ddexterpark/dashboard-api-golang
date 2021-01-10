package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
	"time"
)

type Licenses []struct {
	License
}

type License struct {
	ID                        string      `json:"id"`
	LicenseType               string      `json:"licenseType"`
	LicenseKey                string      `json:"licenseKey"`
	OrderNumber               string      `json:"orderNumber"`
	DeviceSerial              string      `json:"deviceSerial"`
	NetworkID                 string      `json:"networkId"`
	State                     string      `json:"state"`
	SeatCount                 interface{} `json:"seatCount"`
	TotalDurationInDays       string         `json:"totalDurationInDays"`
	DurationInDays            string         `json:"durationInDays"`
	PermanentlyQueuedLicenses []struct {
		ID             string `json:"id"`
		LicenseType    string `json:"licenseType"`
		LicenseKey     string `json:"licenseKey"`
		OrderNumber    string `json:"orderNumber"`
		DurationInDays string `json:"durationInDays"`
	} `json:"permanentlyQueuedLicenses"`
	ClaimDate      time.Time `json:"claimDate"`
	ActivationDate time.Time `json:"activationDate"`
	ExpirationDate time.Time `json:"expirationDate"`
}

type AssignSeats struct {
	ResultingLicenses []struct {
		ID                        string        `json:"id"`
		LicenseType               string        `json:"licenseType"`
		LicenseKey                string        `json:"licenseKey"`
		OrderNumber               string        `json:"orderNumber"`
		DeviceSerial              interface{}   `json:"deviceSerial"`
		NetworkID                 string        `json:"networkId"`
		State                     string        `json:"state"`
		SeatCount                 string           `json:"seatCount"`
		TotalDurationInDays       string           `json:"totalDurationInDays"`
		DurationInDays            string           `json:"durationInDays"`
		PermanentlyQueuedLicenses []interface{} `json:"permanentlyQueuedLicenses"`
		ClaimDate                 time.Time     `json:"claimDate"`
		ActivationDate            time.Time     `json:"activationDate"`
		ExpirationDate            time.Time     `json:"expirationDate"`
	} `json:"resultingLicenses"`
}

type MoveLicenses struct {
	DestOrganizationID string   `json:"destOrganizationId"`
	LicenseIds         []string `json:"licenseIds"`
}

type MoveSeats struct {
	DestOrganizationID string `json:"destOrganizationId"`
	LicenseID          string `json:"licenseId"`
	SeatCount          string `json:"seatCount"`
}

type RenewSeats struct {
	ResultingLicenses []struct {
		ID                        string        `json:"id"`
		LicenseType               string        `json:"licenseType"`
		LicenseKey                string        `json:"licenseKey"`
		OrderNumber               string        `json:"orderNumber"`
		DeviceSerial              interface{}   `json:"deviceSerial"`
		NetworkID                 string        `json:"networkId"`
		State                     string        `json:"state"`
		SeatCount                 string           `json:"seatCount"`
		TotalDurationInDays       string           `json:"totalDurationInDays"`
		DurationInDays            string           `json:"durationInDays"`
		PermanentlyQueuedLicenses []interface{} `json:"permanentlyQueuedLicenses"`
		ClaimDate                 time.Time     `json:"claimDate"`
		ActivationDate            time.Time     `json:"activationDate"`
		ExpirationDate            time.Time     `json:"expirationDate"`
	} `json:"resultingLicenses"`
}


// List the licenses for an organization
func GetLicenses(organizationId, perPage, startingAfter,
	endingBefore, deviceSerial, networkId, state string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/licenses",
		organizationId)
	var datamodel = Licenses{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage":          perPage,
		"startingAfter":    startingAfter,
		"endingBefore":		endingBefore,
		"deviceSerial": 	deviceSerial,
		"networkId":		networkId,
		"state":			state,
	}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Display a license
func GetLicense(organizationId, licenseId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/licenses/%s",
		organizationId, licenseId)
	var datamodel = License{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


// Update a license
func PutLicense(organizationId, licenseId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/licenses/%s",
		organizationId, licenseId)
	var datamodel = License{}
	sessions, err := api.Sessions(baseurl, "PUT", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Assign SM seats to a network. This will increase the managed SM device limit of the network
func PostAssignSeats(organizationId, deviceSerial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/licenses/assignSeats",
		organizationId)
	var datamodel = AssignSeats{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Move licenses to another organization. This will also move any devices that the licenses are assigned to
func PostMoveLicenses(organizationId, deviceSerial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/licenses/move",
		organizationId)
	var datamodel = AssignSeats{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Move SM seats to another organization
func PostMoveSeats(organizationId, deviceSerial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/licenses/moveSeats",
		organizationId)
	var datamodel = MoveSeats{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Renew SM seats of a license.
// This will extend the license expiration date of managed SM devices covered by this license
func PostRenewSeats(organizationId, deviceSerial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/licenses/renewSeats",
		organizationId)
	var datamodel = RenewSeats{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
