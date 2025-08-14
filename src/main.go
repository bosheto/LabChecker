package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

// # We are storing the hosts in two places
// var hosts = []host{}

// var hostMap = make(map[int]host)
var hosts = []host{
	// {ID: 0, IP: "10.238.105.130", Status: "Down", returned: false, LastUp: ""},
	// {ID: 1, IP: "10.238.105.145", Status: "Down", returned: false, LastUp: ""},
}

// var hostMapIP = make(map[string]int)

func initHosts() {

	jsonData, err := os.ReadFile("hosts.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonData, &hosts)
	if err != nil {
		panic(err)
	}
}

func main() {

	initHosts()

	go pingAllHostsPeriodically()

	// Intialize Router
	router := gin.Default()

	// Routes

	// GET
	router.GET("/hosts", getHosts)
	router.GET("/hosts/:id", getHost)
	// router.GET("/ping", runPings)

	// POST
	router.POST("/hosts", addHost)

	//PUT
	router.PUT("/hosts", updateHost)

	router.Run("localhost:8080")
}
