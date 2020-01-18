package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os"
)

type IpList struct {
	Domain string   `json:"domain"`
	IPs    []string `json:"ips"`
}

func getIpList(domain string) IpList {
	ips, _ := net.LookupIP(domain)

	var ipList IpList
	ipList.Domain = domain

	for _, ip := range ips {
		ipList.IPs = append(ipList.IPs, ip.String())
	}

	return ipList
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ipListArr []IpList

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ipListArr = append(ipListArr, getIpList(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	jsonData, _ := json.MarshalIndent(ipListArr, "", "  ")

	ioutil.WriteFile(os.Args[2], jsonData, 0644)
}
