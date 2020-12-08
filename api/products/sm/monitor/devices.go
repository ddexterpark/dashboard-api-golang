package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)

type CellularUsageHistory []struct {
	Received int       `json:"received"`
	Sent     int       `json:"sent"`
	Ts       time.Time `json:"ts"`
}

type Connectivity []struct {
	FirstSeenAt time.Time `json:"firstSeenAt"`
	LastSeenAt  time.Time `json:"lastSeenAt"`
}

type DesktopLogs []struct {
	MeasuredAt    time.Time `json:"measuredAt"`
	User          string    `json:"user"`
	NetworkDevice string    `json:"networkDevice"`
	NetworkDriver string    `json:"networkDriver"`
	WifiChannel   string    `json:"wifiChannel"`
	WifiAuth      string    `json:"wifiAuth"`
	WifiBssid     string    `json:"wifiBssid"`
	WifiSsid      string    `json:"wifiSsid"`
	WifiRssi      string    `json:"wifiRssi"`
	WifiNoise     string    `json:"wifiNoise"`
	DhcpServer    string    `json:"dhcpServer"`
	IP            string    `json:"ip"`
	NetworkMTU    string    `json:"networkMTU"`
	Subnet        string    `json:"subnet"`
	Gateway       string    `json:"gateway"`
	PublicIP      string    `json:"publicIP"`
	DNSServer     string    `json:"dnsServer"`
	Ts            time.Time `json:"ts"`
}

type DeviceCommandLogs []struct {
	Action        string    `json:"action"`
	Name          string    `json:"name"`
	Details       string    `json:"details"`
	DashboardUser string    `json:"dashboardUser"`
	Ts            time.Time `json:"ts"`
}

type PerformanceHistory []struct {
	CPUPercentUsed  float64 `json:"cpuPercentUsed"`
	MemFree         int     `json:"memFree"`
	MemWired        int     `json:"memWired"`
	MemActive       int     `json:"memActive"`
	MemInactive     int     `json:"memInactive"`
	NetworkSent     int     `json:"networkSent"`
	NetworkReceived int     `json:"networkReceived"`
	SwapUsed        int     `json:"swapUsed"`
	DiskUsage       struct {
		C struct {
			Used  int `json:"used"`
			Space int `json:"space"`
		} `json:"c"`
	} `json:"diskUsage"`
	Ts time.Time `json:"ts"`
}

// Return the client's daily cellular data usage history
func GetCellularUsageHistory(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm/devices/%s/cellularUsageHistory",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = CellularUsageHistory{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Returns historical connectivity data (whether a device is regularly checking in to Dashboard).
func GetConnectivity(networkId, deviceId,
	perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm//devices/%s/connectivity",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = Connectivity{}

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

// Return historical records of various Systems Manager network connection details for desktop devices.
func GetDesktopLogs(networkId, deviceId,
	perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm//devices/%s/desktopLogs",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = DesktopLogs{}

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

// Return historical records of commands sent to Systems Manager devices
func GetDeviceCommandLogs(networkId, deviceId,
	perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm//devices/%s/deviceCommandLogs",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = DeviceCommandLogs{}

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

// Return historical records of various Systems Manager client metrics for desktop devices.
func GetPerformanceHistory(networkId, deviceId,
	perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/sm//devices/%s/performanceHistory",
		api.BaseUrl(), networkId, deviceId)
	var datamodel = PerformanceHistory{}

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
