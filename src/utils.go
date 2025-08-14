package main

import "os"

type jsonError struct {
	Error string `json:"error"`
}

func ip_exists(ip string) bool {
	for i := 0; i < len(hostMap); i++ {
		if hostMap[i].IP == ip {
			return true
		}
	}
	return false
}

func write_file(content []byte) {
	file, err := os.Create("hosts.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		panic(err)
	}
}
