package main

import (
	"github.com/gin-gonic/gin"
)

// # We are storing the hosts in two places
// var hosts = []host{}
var hostMap = make(map[int]host)

// var hostMapIP = make(map[string]int)

func initHosts() {

	testHosts := []host{
		{ID: 0, IP: "10.238.105.130", Status: "Down", returned: false, LastUp: ""},
		{ID: 1, IP: "10.238.105.145", Status: "Donw", returned: false, LastUp: ""},
	}

	// hosts = append(hosts, testHosts...)

	for i := 0; i < len(testHosts); i++ {
		hostMap[testHosts[i].ID] = testHosts[i]
	}

	// for i := 0; i < len(hosts); i++ {
	// 	hostMapIP[testHosts[i].IP] = hosts[i].ID
	// }
}

func main() {

	initHosts()

	// Intialize Router
	router := gin.Default()

	// Routes

	// GET
	router.GET("/hosts", getHosts)
	router.GET("/hosts/:id", getHost)
	router.GET("/ping", runPings)

	// POST
	router.POST("/hosts", addHost)

	//PUT
	router.PUT("/hosts", updateHost)

	router.Run("localhost:8080")
}
