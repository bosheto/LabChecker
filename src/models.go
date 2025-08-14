package main

type host struct {
	ID       int    `json:"id"`
	IP       string `json:"ip"`
	Status   string `json:"staus"`
	LastUp   string `json:"last_up"`
	returned bool
	Name     string `json:"name"`
	Version  string `json:"version"`
}
