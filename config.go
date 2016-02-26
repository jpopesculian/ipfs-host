package main

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Host              string
	Port              uint64
	UserServiceHost   string
	DomainServiceHost string
	IpfsHost          string
	RootHash          string
}

var config Config

func InitConfig() {
	host := flag.String("h", GetOsStringVal("HOST", "localhost"), "Server Host")
	port := flag.Uint64("p", GetOsUint64Val("PORT", 5000), "Port to serve on")

	defaultUserServiceHost := GetOsStringVal("USER_SERVICE_HOST", "localhost:5001")
	userServiceHost := flag.String("u", defaultUserServiceHost, "Host of the User Service")

	defaultDomainServiceHost := GetOsStringVal("DOMAIN_SERVICE_HOST", "localhost:5002")
	domainServiceHost := flag.String("d", defaultDomainServiceHost, "Host of the Domain Mapping Service")

	defaultIpfsHost := GetOsStringVal("IPFS_HOST", "localhost:8080")
	ipfsHost := flag.String("i", defaultIpfsHost, "Host of the IPFS Server")

	defaultRootHash := GetOsStringVal("ROOT_HASH", "QmPXME1oRtoT627YKaDPDQ3PwA8tdP9rWuAAweLzqSwAWT")
	rootHash := flag.String("r", defaultRootHash, "Location of Landing Page")

	flag.Parse()

	config = Config{
		*host,
		*port,
		*userServiceHost,
		*domainServiceHost,
		*ipfsHost,
		*rootHash,
	}
}

func GetOsStringVal(envVar, initial string) string {
	defaultValue := os.Getenv(envVar)
	if len(defaultValue) == 0 {
		defaultValue = initial
	}
	return defaultValue
}

func GetOsUint64Val(envVar string, initial uint64) uint64 {
	defaultValue, err := strconv.ParseUint(os.Getenv(envVar), 0, 16)
	if err != nil {
		defaultValue = initial
	}
	return defaultValue
}
