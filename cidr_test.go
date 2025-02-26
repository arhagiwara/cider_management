package main

import (
	"net"
	"testing"
)

func testCidr(t *testing.T, cidrString, minString, maxString string) {
	_, cidr, _ := net.ParseCIDR(cidrString)
	got := MakeCidrInfo(cidr)
	want := CidrInfo{
		Cidr:        cidr,
		NetworkIP:   net.ParseIP(minString),
		BroadcastIP: net.ParseIP(maxString),
	}
	if got.Cidr.String() != want.Cidr.String() {
		t.Errorf("MakeCidrInfo() = %v, want %v", got.Cidr, want.Cidr)
	}
	if got.NetworkIP.String() != want.NetworkIP.String() {
		t.Errorf("MakeCidrInfo() = %v, want %v", got.NetworkIP, want.NetworkIP)
	}
	if got.BroadcastIP.String() != want.BroadcastIP.String() {
		t.Errorf("MakeCidrInfo() = %v, want %v", got.BroadcastIP, want.BroadcastIP)
	}
}
func TestMakeCidrInfo(t *testing.T) {
	testCidr(t, "10.10.10.10/24", "10.10.10.0", "10.10.10.255")
	testCidr(t, "10.10.10.10/32", "10.10.10.10", "10.10.10.10")
	testCidr(t, "10.10.10.10/31", "10.10.10.10", "10.10.10.11")
	testCidr(t, "10.10.10.10/8", "10.0.0.0", "10.255.255.255")
}

func TestOverlapCidr(t *testing.T) {
	_, cidr1, _ := net.ParseCIDR("10.10.10.0/24")
	_, cidr2, _ := net.ParseCIDR("10.10.0.0/16")
	got := OverlapCidr(MakeCidrInfo(cidr1), MakeCidrInfo(cidr2))
	if got != true {
		t.Errorf("OverlapCidr() = %v, want %v", got, true)
	}
	_, cidr1, _ = net.ParseCIDR("10.10.10.0/24")
	_, cidr2, _ = net.ParseCIDR("10.11.0.0/16")
	got = OverlapCidr(MakeCidrInfo(cidr1), MakeCidrInfo(cidr2))
	if got != false {
		t.Errorf("OverlapCidr() = %v, want %v", got, false)
	}
	_, cidr1, _ = net.ParseCIDR("10.10.10.0/26")
	_, cidr2, _ = net.ParseCIDR("10.10.10.100/30")
	got = OverlapCidr(MakeCidrInfo(cidr1), MakeCidrInfo(cidr2))
	if got != false {
		t.Errorf("OverlapCidr() = %v, want %v", got, true)
	}
	_, cidr1, _ = net.ParseCIDR("10.10.10.0/26")
	_, cidr2, _ = net.ParseCIDR("10.10.10.63/32")
	got = OverlapCidr(MakeCidrInfo(cidr1), MakeCidrInfo(cidr2))
	if got != true {
		t.Errorf("OverlapCidr() = %v, want %v", got, true)
	}
	_, cidr1, _ = net.ParseCIDR("10.10.10.0/26")
	_, cidr2, _ = net.ParseCIDR("10.10.10.64/32")
	got = OverlapCidr(MakeCidrInfo(cidr1), MakeCidrInfo(cidr2))
	if got != false {
		t.Errorf("OverlapCidr() = %v, want %v", got, true)
	}
}
