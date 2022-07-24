package network

import (
	"net"
)

func LookupDomain(ip string) ([]net.IP, error) {
	addr, err := net.LookupIP(ip)
	if err != nil {
		return nil, err
	}
	return addr, nil
}
