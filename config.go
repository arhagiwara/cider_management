package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net"
	"os"
)

type Config struct {
	CidrInfo CidrInfo
	Date     string
	Use      bool
	Comment  string
}

func readConfig(csv_file string) ([]Config, error) {
	file, err := os.Open(csv_file)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.Comment = '#'

	var config []Config

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		_, ipv4Net, err := net.ParseCIDR(line[0])
		if err != nil {
			return nil, err
		}
		if line[2] != "true" && line[2] != "false" {
			return nil, fmt.Errorf("invalid use value: %s", line[2])
		}
		config = append(config, Config{
			CidrInfo: MakeCidrInfo(ipv4Net),
			Date:     line[1],
			Use:      line[2] == "true",
			Comment:  line[3],
		})
	}
	return config, nil
}
