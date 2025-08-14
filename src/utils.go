package main

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
