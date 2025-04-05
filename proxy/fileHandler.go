package proxy

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadHosts() []string {
	file, err := os.Open(FORBIDDEN_HOSTS_PATH)
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

func loadWords() []string {
	file, err := os.Open(BANNED_WORDS_PATH)
	if err != nil {
		fmt.Println("Error Processing Banned Words", err)
		os.Exit(1)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
	}
	return words
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

func containsBannedWord(host string) bool {
	bannedWords := loadWords()

	for i := 0; i < len(bannedWords); i++ {
		if strings.Contains(host, bannedWords[i]) {
			return true
		}
	}
	return false
}
