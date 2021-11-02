package bt

import (
	"fmt"
	"strconv"
	"strings"

	"tinygo.org/x/bluetooth"
)

func ParseMACAddress(in string) (bluetooth.MAC, error) {
	inSlice := strings.Split(in, ":")
	mac := bluetooth.MAC{}

	for i := 0; i < len(inSlice); i++ {
		i64, err := strconv.ParseUint(inSlice[len(inSlice)-1- i], 16, 8)
		if err != nil {
			return mac, err
		}

		mac[i] = byte(i64)
	}

	return mac, nil
}

func GetDevice(addr string) (*bluetooth.Device, error) {
	adapter := bluetooth.DefaultAdapter

	if err := adapter.Enable(); err != nil {
		return nil, fmt.Errorf("could not enable the adapter: %v", err)
	}

	mac, err := ParseMACAddress(addr)
	if err != nil {
		return nil, fmt.Errorf("could not parse %q as a MAC address: %v", addr, err)
	}

	btAddr := bluetooth.Address{
		MACAddress: bluetooth.MACAddress{
			MAC: mac,
		},
	}

	return adapter.Connect(btAddr, bluetooth.ConnectionParams{})
}