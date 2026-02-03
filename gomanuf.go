package gomanuf

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type MacKey struct {
	Masked uint64
	CIDR   uint8
}

var data map[MacKey]string

var macRegex, _ = regexp.Compile("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$")

func init() {
	data = make(map[MacKey]string)
	_, location, _, _ := runtime.Caller(0)
	file, err := os.ReadFile(path.Join(path.Dir(location), "manuf.txt"))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for i := 0; i < len(lines); i++ {
		line := strings.ReplaceAll(lines[i], "\t\t", "\t")
		fields := strings.Split(line, "\t")

		// Ignore the comments and empty lines
		if strings.HasPrefix(fields[0], "#") || line == "" {
			continue
		}

		mac := fields[0]
		manuf := fields[1]

		if strings.Contains(mac, "/") {
			parts := strings.SplitN(mac, "/", 2)
			macPrefix := parts[0]

			cidrVal, err := strconv.ParseUint(parts[1], 10, 8)
			if err != nil {
				continue
			}
			cidr := uint8(cidrVal)

			if cidr == 28 || cidr == 36 {
				macVal, err := macToUint64(macPrefix)
				if err != nil {
					continue
				}

				masked := maskMac(macVal, cidr)
				data[MacKey{Masked: masked, CIDR: cidr}] = manuf
			}
		} else {
			macVal, err := macToUint64(mac)
			if err != nil {
				continue
			}

			cidr := uint8(24)
			masked := maskMac(macVal, cidr)
			data[MacKey{Masked: masked, CIDR: cidr}] = manuf
		}
	}
}

// Lookup will return the manufacturer of the given MAC address.
func Lookup(mac string) (string, error) {
	newMac := strings.ReplaceAll(strings.ToUpper(mac), "-", ":")
	if !macRegex.MatchString(newMac) {
		return "", fmt.Errorf("invalid MAC address")
	}
	macVal, err := macToUint64(newMac)
	if err != nil {
		return "", fmt.Errorf("invalid MAC format")
	}

	cidrs := []uint8{36, 28, 24}
	for _, cidr := range cidrs {
		masked := maskMac(macVal, cidr)
		if manuf, ok := data[MacKey{Masked: masked, CIDR: cidr}]; ok {
			return manuf, nil
		}
	}

	return "unknown", nil
}

// Deprecated: use Lookup instead.
func Search(mac string) (string, error) {
	return Lookup(mac)
}
