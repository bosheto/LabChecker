package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tatsushid/go-fastping"
)

var pinger = fastping.NewPinger()

type host struct {
	IP       string `json:"ip"`
	Status   string `json:"staus"`
	LastUp   string `json:"last_up"`
	returned bool
}

type jsonError struct {
	Error string `json:"error"`
}

var hosts = []host{}
var hostsMap = make(map[string]host)

func initHosts() {
	testHosts := []host{
		{IP: "10.238.105.130", Status: "Down", returned: false, LastUp: ""},
		{IP: "10.238.105.145", Status: "Donw", returned: false, LastUp: ""},
	}

	hosts = append(hosts, testHosts...)

	for i := 0; i < len(hosts); i++ {
		hostsMap[hosts[i].IP] = hosts[i]
	}
}

func getHostsStatus(c *gin.Context) {
	c.JSON(http.StatusOK, hosts)
}

func runPings(c *gin.Context) {
	c.String(http.StatusOK, "ping")
	pingHostMachine()
}

func pingHostMachine() {

	for i := 0; i < len(hosts); i++ {
		pinger.AddIP(hosts[i].IP)
	}

	pinger.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		// fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		for i := 0; i < len(hosts); i++ {
			if hosts[i].IP == addr.String() {
				hosts[i].returned = true
			}
		}
	}

	pinger.OnIdle = func() {
		fmt.Println("finish")
		for i := 0; i < len(hosts); i++ {
			if hosts[i].returned == true {
				hosts[i].Status = "Up"
				hosts[i].LastUp = time.Now().Format(time.UnixDate)
			} else {
				hosts[i].Status = "Down"
			}
		}
	}

	err := pinger.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func addHost(c *gin.Context) {
	var newHost host

	if err := c.BindJSON(&newHost); err != nil {
		return
	}

	_, exists := hostsMap[newHost.IP]

	if exists {
		c.JSON(http.StatusBadRequest, jsonError{Error: "Ip allready in use!"})
		return
	}

	hosts = append(hosts, newHost)
	hostsMap[newHost.IP] = newHost
	c.JSON(http.StatusCreated, newHost)
}

func main() {

	initHosts()

	// Intialize Router
	router := gin.Default()

	// Routes

	// GET
	router.GET("/hosts", getHostsStatus)
	router.GET("/ping", runPings)

	// POST
	router.POST("/hosts", addHost)

	router.Run("localhost:8080")
}
