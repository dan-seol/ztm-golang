//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayServer(mapping map[string]int) {
	statuses := []string{"Online", "Offline", "Maintenance", "Retired"}
	fmt.Println("There are total", len(mapping), "servers")
	for serverName, serverStatus := range mapping {
		fmt.Println("server", serverName, "is of status", statuses[serverStatus])
	}
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	mapping := make(map[string]int)

	for _, server := range servers {
		mapping[server] = Online
	}
	displayServer(mapping)

	mapping["darkstar"] = Retired
	mapping["aiur"] = Offline
	displayServer(mapping)

	for _, server := range servers {
		mapping[server] = Maintenance
	}

	displayServer(mapping)
}
