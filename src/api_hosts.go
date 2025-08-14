package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getHosts(c *gin.Context) {
	c.JSON(http.StatusOK, hosts)
}

func getHost(c *gin.Context) {
	id := c.Param("id")

	hostID, err := strconv.Atoi(id)
	if err != nil {
		// Respond with an error if the conversion fails
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// value, exists := hosts[hostID]

	if hostID > len(hosts)-1 || hostID < 0 {
		c.JSON(http.StatusNotFound, jsonError{Error: "No hosts with ip: " + string(id)})
	} else {
		c.JSON(http.StatusOK, hosts[hostID])
	}

}

func addHost(c *gin.Context) {
	var newHost host

	if err := c.BindJSON(&newHost); err != nil {
		return
	}

	if ip_exists(newHost.IP) {
		c.JSON(http.StatusBadRequest, jsonError{Error: "Ip allready in use!"})
		return
	}

	newHost.ID = len(hosts)
	// hostMap[newHost.ID] = newHost
	hosts = append(hosts, newHost)
	c.JSON(http.StatusCreated, newHost)
	b, err := json.Marshal(hosts)
	if err != nil {
		return
	}
	write_file(b)

}

// #TODO need to fix the updating
func updateHost(c *gin.Context) {
	var updatedHost host

	if err := c.ShouldBindBodyWithJSON(&updatedHost); err != nil {
		c.JSON(http.StatusBadRequest, jsonError{Error: err.Error()})
		return
	}
	id := updatedHost.ID
	if id < 0 || id > len(hosts)-1 {
		c.JSON(http.StatusBadRequest, jsonError{Error: "No host with id: " + string(id)})
	} else {
		hosts[id].IP = updatedHost.IP
		hosts[id].Name = updatedHost.Name
		hosts[id].Version = updatedHost.Version
		c.JSON(http.StatusOK, hosts[id])
	}

}

// func deleteHost(c *gin.Context){
// 	id := c.Param("id")

// }
