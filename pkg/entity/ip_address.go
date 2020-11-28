package entity

import (
	"fmt"
	"net"
	"strconv"
)

// IPAddress represents the ip address includes host and port.
type IPAddress struct {
	host string
	port string
}

// NewIPAddress creates the valdi instance of ip address.
func NewIPAddress(addr string) (IPAddress, error) {
	fmt.Println(addr)
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return IPAddress{}, fmt.Errorf("net split host port: %w", err)
	}
	fmt.Println(host, port)
	ip := net.ParseIP(host)
	if ip == nil {
		return IPAddress{}, fmt.Errorf("invlid host")
	}
	if _, err := strconv.ParseUint(port, 10, 16); err != nil {
		return IPAddress{}, fmt.Errorf("invalid port: %w", err)
	}
	return IPAddress{
		host: host,
		port: port,
	}, nil
}

// Host returns the host of ip address.
func (ip IPAddress) Host() string {
	return ip.host
}

// Port returns the port of ip address.
func (ip IPAddress) Port() string {
	return ip.port
}
