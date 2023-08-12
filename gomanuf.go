package gomanuf

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
)

var data map[string]string
var slash28 map[string]string
var slash36 map[string]string

var macRegex, _ = regexp.Compile("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$")

func init() {
	data = make(map[string]string)
	slash28 = make(map[string]string)
	slash36 = make(map[string]string)
	_, location, _, _ := runtime.Caller(0)
	file, err := os.ReadFile(path.Join(path.Dir(location), "manuf.txt"))
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	for i := 0; i < len(lines); i++ {
		line := strings.Replace(lines[i], "\t\t", "\t", -1)
		fields := strings.Split(line, "\t")

		// Ignore the comments and empty lines
		if strings.HasPrefix(fields[0], "#") || line == "" {
			continue
		}

		mac := fields[0]
		manuf := fields[1]
		if len(fields) > 2 {
			manuf = fields[2]
		}
		if strings.Contains(mac, ":00/28") {
			slash28[mac] = manuf
		} else if strings.Contains(mac, ":00/36") {
			slash36[mac] = manuf
		}

		data[mac] = manuf
	}
}

func checkSlash28(mac string) string {
	mac = mac[:10] + "0:00:00/28"
	for address, manuf := range slash28 {
		if address == mac {
			return manuf
		}
	}
	return ""
}

func checkSlash36(mac string) string {
	mac = mac[:13] + "0:00/36"
	for address, manuf := range slash36 {
		if address == mac {
			return manuf
		}
	}
	return ""
}

// Search will return the manufacturer of the given MAC address.
func Search(mac string) (string, error) {
	mac = strings.Replace(strings.ToUpper(mac), "-", ":", -1)
	if !macRegex.MatchString(mac) {
		return "Invalid MAC address", fmt.Errorf("invalid MAC address")
	}
	for address, manuf := range data {
		if strings.HasPrefix(mac, address) {
			// Check if manufacturer is one of those manufacturer that have MACs with /28 or /36
			if manuf == "IEEE Registration Authority" {
				check28 := checkSlash28(mac)
				if check28 != "" {
					return check28, nil
				}
				check36 := checkSlash36(mac)
				if check36 != "" {
					return check36, nil
				}
			}

			return manuf, nil
		}
	}
	return "unknown", nil
}
