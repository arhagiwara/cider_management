package main

import (
	"net"
)

type CidrInfo struct {
	Cidr        *net.IPNet
	NetworkIP   net.IP
	BroadcastIP net.IP
}

func MakeCidrInfo(cidr *net.IPNet) CidrInfo {
	var networkIP, broadcastIP net.IP = make(net.IP, 4), make(net.IP, 4)
	for i := 0; 4 > i; i++ {
		networkIP[i] = cidr.IP[i] & cidr.Mask[i]
		broadcastIP[i] = cidr.IP[i] | ^cidr.Mask[i]
	}
	return CidrInfo{
		Cidr:        cidr,
		NetworkIP:   networkIP,
		BroadcastIP: broadcastIP,
	}
}

func OverlapCidr(cidr1, cidr2 CidrInfo) bool {
	if cidr1.Cidr.Contains(cidr2.NetworkIP) || cidr1.Cidr.Contains(cidr2.BroadcastIP) {
		return true
	}
	if cidr2.Cidr.Contains(cidr1.NetworkIP) || cidr2.Cidr.Contains(cidr1.BroadcastIP) {
		return true
	}
	return false
}
