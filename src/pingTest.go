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

func runPings(c *gin.Context) {
	c.String(http.StatusOK, "ping")
	pingHostMachine()

}

func pingHostMachine() {

	for i := 0; i < len(hostMap); i++ {
		pinger.AddIP(hostMap[i].IP)
	}

	pinger.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		for i := 0; i < len(hostMap); i++ {
			host := hostMap[i]
			if host.IP == addr.String() {
				host.returned = true
			} else {
				host.returned = false
			}
			hostMap[i] = host
		}
	}

	pinger.OnIdle = func() {
		fmt.Println("finish")
		for i := 0; i < len(hostMap); i++ {
			host := hostMap[i]
			if host.returned {
				host.Status = "Up"
				host.LastUp = time.Now().Format(time.UnixDate)
			} else {
				host.Status = "Down"
			}
			hostMap[i] = host
		}
	}

	err := pinger.Run()
	if err != nil {
		fmt.Println(err)
	}
}
