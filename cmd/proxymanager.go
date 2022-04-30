package cmd

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"
)

var proxyRegex = `((http|https|socks4|socks5):\/\/)?([\w-]+\.)+[\w-]+(\/[\w- .\/?%&=]*)?`

//LoadProxies is a function to load each line of 'appconfig/proxies' into a slice of strings
func LoadProxies() []string {
	var proxies []string
	file, err := os.Open("appconfig/proxies")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return proxies
}

//GetProxy is a function to return a random proxy from the slice of proxies
func GetProxy() string {
	proxies := LoadProxies()
	//seed the random number generator with the current time in nanoseconds for best results :> )
	rand.Seed(int64(time.Now().Nanosecond()))
	proxy := proxies[rand.Intn(len(proxies))]
	return proxy
}

//CheckProxy is a function to check if a proxy is valid using the proxyRegex
func CheckProxy(proxy string) bool {
	valid, _ := regexp.MatchString(proxyRegex, proxy)
	return valid
}
