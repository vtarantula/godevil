package netutil

import "net"

func IsValidIP4(val string) bool {
	ipaddr := net.ParseIP(val)
	return ipaddr != nil
}
