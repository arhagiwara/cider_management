package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <config_file>")
		os.Exit(1)
	}
	filepath := os.Args[1]
	config, err := readConfig(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < len(config); i++ {
		for j := i + 1; j < len(config); j++ {
			fmt.Printf("compare %s and %s\n", config[i].CidrInfo.Cidr.String(), config[j].CidrInfo.Cidr.String())
			if OverlapCidr(config[i].CidrInfo, config[j].CidrInfo) {
				fmt.Printf("%s and %s overlap\n", config[i].CidrInfo.Cidr.String(), config[j].CidrInfo.Cidr.String())
				os.Exit(1)
			}
		}
	}
	fmt.Println("cider config is OK")
}
