package main

import (
	"net"
	"testing"
)

func CompareConfig(got, want Config) bool {
	if got.CidrInfo.Cidr.String() != want.CidrInfo.Cidr.String() {
		return false
	}
	if got.CidrInfo.NetworkIP.String() != want.CidrInfo.NetworkIP.String() {
		return false
	}
	if got.CidrInfo.BroadcastIP.String() != want.CidrInfo.BroadcastIP.String() {
		return false
	}
	if got.Date != want.Date {
		return false
	}
	if got.Use != want.Use {
		return false
	}
	if got.Comment != want.Comment {
		return false
	}
	return true
}

func TestReadConfig(t *testing.T) {
	got, err := readConfig("test_data/test_config.csv")
	if err != nil {
		t.Errorf("readConfig error: %v", err)
	}
	want := Config{
		CidrInfo: CidrInfo{
			Cidr:        &net.IPNet{IP: net.ParseIP("10.120.0.0"), Mask: net.CIDRMask(16, 32)},
			NetworkIP:   net.ParseIP("10.120.0.0"),
			BroadcastIP: net.ParseIP("10.120.255.255"),
		},
		Date:    "2020/01/11",
		Use:     true,
		Comment: "office1 cidr",
	}
	if CompareConfig(got[0], want) == false {
		t.Errorf("readConfig() = %v, want %v", got[0], want)
	}
	if len(got) != 2 {
		t.Errorf("readConfig() = %v, want %v", len(got), 2)
	}
}
