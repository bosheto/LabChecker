package main

import (
	"time"

	"net"

	"github.com/tatsushid/go-fastping"
)

func pingHostFast(ip string) bool {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ip)
	if err != nil {
		return false
	}
	p.AddIPAddr(ra)
	result := false
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		result = true
	}
	p.OnIdle = func() {}
	p.MaxRTT = 2 * time.Second
	_ = p.Run()
	return result
}

func pingAllHostsPeriodically() {
	for {
		for i := range hosts {
			status := "Down"
			if pingHostFast(hosts[i].IP) {
				status = "Up"
				hosts[i].LastUp = time.Now().Format(time.RFC3339)
			}
			hosts[i].Status = status
		}
		time.Sleep(30 * time.Second)
	}
}
