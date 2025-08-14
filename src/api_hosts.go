package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getHosts(c *gin.Context) {
	c.JSON(http.StatusOK, hostMap)
}

func getHost(c *gin.Context) {
	id := c.Param("id")

	hostID, err := strconv.Atoi(id)
	if err != nil {
		// Respond with an error if the conversion fails
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	value, exists := hostMap[hostID]

	if exists {
		c.JSON(http.StatusOK, value)
	} else {
		c.JSON(http.StatusNotFound, jsonError{Error: "No hosts with ip: " + string(id)})
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

	newHost.ID = len(hostMap)
	hostMap[newHost.ID] = newHost
	c.JSON(http.StatusCreated, newHost)
	b, err := json.Marshal(hostMap)
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

	if _, exists := hostMap[updatedHost.ID]; exists {
		host := hostMap[updatedHost.ID]
		host.IP = updatedHost.IP
		host.Name = updatedHost.Name
		host.Version = updatedHost.Version
		hostMap[updatedHost.ID] = host
		c.JSON(http.StatusOK, hostMap[updatedHost.ID])
	} else {
		c.JSON(http.StatusBadRequest, jsonError{Error: "No host with id: "})
	}
}

// func deleteHost(c *gin.Context){
// 	id := c.Param("id")

// }
