package xip

import (
	"fmt"
	"net"
)

func GenerateAllIP(chanIP chan string) {
	for i := 0; i <= 255; i++ {
		for j := 0; j <= 255; j++ {
			for k := 0; k <= 255; k++ {
				for l := 0; l <= 255; l++ {
					ip := fmt.Sprintf("%d.%d.%d.%d", i, j, k, l)
					chanIP <- ip
				}
			}
		}
	}
	close(chanIP)
}

func GenerateRangeIP(startIP, endIP string, chanIP chan string) error {
	ipStart := net.ParseIP(startIP)
	if ipStart == nil {
		return fmt.Errorf("invalid start IP address: %s", startIP)
	}
	ipEnd := net.ParseIP(endIP)
	if ipEnd == nil {
		return fmt.Errorf("invalid end IP address: %s", endIP)
	}

	ip := ipStart
	for {
		if ipToUint32(ip) > ipToUint32(ipEnd) {
			break
		}

		chanIP <- ip.String()

		ip = nextIP(ip)
	}
	fmt.Println("ip done")
	close(chanIP)
	return nil
}

func ipToUint32(ip net.IP) uint32 {
	ip = ip.To4()
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
}

func nextIP(ip net.IP) net.IP {
	next := make(net.IP, len(ip))
	copy(next, ip)
	for i := len(next) - 1; i >= 0; i-- {
		if next[i] < 255 {
			next[i]++
			break
		} else {
			next[i] = 0
		}
	}
	return next
}
