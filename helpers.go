package gomanuf

import (
	"fmt"
	"strconv"
	"strings"
)

func macToUint64(mac string) (uint64, error) {
	parts := strings.Split(mac, ":")

	var hexStr strings.Builder
	for _, p := range parts {
		if p != "" {
			hexStr.WriteString(p)
		}
	}

	hex := hexStr.String()

	switch len(hex) {
	case 6:
		hex += "000000"
	case 12:
	default:
		return 0, fmt.Errorf("invalid MAC format: %s", mac)
	}

	val, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return 0, err
	}

	return val, nil
}

func maskMac(mac uint64, cidr uint8) uint64 {
	var mask uint64

	switch cidr {
	case 24:
		mask = 0xFFFFFF000000
	case 28:
		mask = 0xFFFFFFF00000
	case 36:
		mask = 0xFFFFFFFFF000
	default:
		mask = 0xFFFFFFFFFFFF
	}

	return mac & mask
}
