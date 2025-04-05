package proxy

import (
	"bufio"
	"fmt"
	"os"
)

func loadHosts() []string {
	path := "proxy/forbidden-hosts.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error Processing Forbidden Hosts", err)
		os.Exit(1)
	}
	defer file.Close()

	var hosts []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		host := scanner.Text()
		hosts = append(hosts, host)
	}
	return hosts
}

func isForbiddenHost(host string) bool {
	forbiddenHosts := loadHosts()

	for i := 0; i < len(forbiddenHosts); i++ {
		if forbiddenHosts[i] == host {
			return true
		}
	}
	return false
}

// func containsBannedWord(host string) bool {

// }
