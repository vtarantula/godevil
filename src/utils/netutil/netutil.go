package netutil

import (
	"errors"
	"fmt"
	"net"
)

func IsValidIP4(val string) bool {
	ipaddr := net.ParseIP(val)
	return ipaddr != nil
}

// Valid only for /24 subnets
// TODO: Support other subnets too
func GetNetworkIPv4(ip string) (string, error) {
	if !IsValidIP4(ip) {
		str_val := fmt.Sprintf("Invalid IP: %s", ip)
		return "", errors.New(str_val)
	}
	ipaddr := net.ParseIP(ip)
	mask := ipaddr.DefaultMask()
	network := ipaddr.Mask(mask)
	return network.String(), nil
}
