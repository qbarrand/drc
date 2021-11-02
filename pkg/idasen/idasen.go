package idasen

import (
	"context"
	"encoding/binary"
	"fmt"

	"tinygo.org/x/bluetooth"
)

// baseHeight is the desk's base height in centimeters.
const baseHeight = 63

type Idasen struct {
	dev *bluetooth.Device
}

func NewIdasen(dev *bluetooth.Device) *Idasen {
	return &Idasen{dev: dev}
}

// GetCurrentHeight returns the desk's current height in centimeters.
func (i *Idasen) GetCurrentHeight(ctx context.Context) (int, error) {
	const (
		serviceUUID = "99fa0020-338a-1024-8a49-009c0215f78a"
		characteristicsUUID = "99fa0021-338a-1024-8a49-009c0215f78a"
	)

	svcUUID, err := bluetooth.ParseUUID(serviceUUID)
	if err != nil {
		return 0, fmt.Errorf("could not parse the service UUID: %v", err)
	}

	charUUID, err := bluetooth.ParseUUID(characteristicsUUID)
	if err != nil {
		return 0, fmt.Errorf("could not parse the characteristics UUID: %v", err)
	}

	huuid := bluetooth.CharacteristicUUIDHeight
	_ = huuid

	svc, err := i.dev.DiscoverServices([]bluetooth.UUID{svcUUID})
	if err != nil {
		return 0, fmt.Errorf("could not discover services: %v", err)
	}

	char, err := svc[0].DiscoverCharacteristics([]bluetooth.UUID{charUUID})
	if err != nil {
		return 0, fmt.Errorf("could not discover the characteristic %s: %v", charUUID, err)
	}

	var rawHeight uint32

	if err = binary.Read(&char[0], binary.LittleEndian, &rawHeight); err != nil {
		return 0, fmt.Errorf("could not read the raw height: %v", err)
	}

	return int(float64(baseHeight) + float64(rawHeight/ 100)), nil
}
