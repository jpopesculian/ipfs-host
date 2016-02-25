package main

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Port              uint64
	UserServiceHost   string
	DomainServiceHost string
	IpfsHost          string
}

var config Config

func InitConfig() {
	defaultPort, err := strconv.ParseUint(os.Getenv("PORT"), 0, 16)
	if err != nil {
		defaultPort = 5000
	}
	port := flag.Uint64("p", defaultPort, "Port to serve on")

	defaultUserServiceHost := os.Getenv("USER_SERVICE_HOST")
	if len(defaultUserServiceHost) == 0 {
		defaultUserServiceHost = "localhost:5001"
	}
	userServiceHost := flag.String("u", defaultUserServiceHost, "Host of the User Service")

	defaultDomainServiceHost := os.Getenv("DOMAIN_SERVICE_HOST")
	if len(defaultDomainServiceHost) == 0 {
		defaultDomainServiceHost = "localhost:5002"
	}
	domainServiceHost := flag.String("d", defaultDomainServiceHost, "Host of the Domain Mapping Service")

	defaultIpfsHost := os.Getenv("IPFS_HOST")
	if len(defaultIpfsHost) == 0 {
		defaultIpfsHost = "ipfs"
	}
	ipfsHost := flag.String("i", defaultIpfsHost, "Host of the IPFS Server")

	flag.Parse()

	config = Config{
		*port,
		*userServiceHost,
		*domainServiceHost,
		*ipfsHost,
	}
}
