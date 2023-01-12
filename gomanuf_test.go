package gomanuf

import "testing"

func TestDLink(t *testing.T) {
	manuf, _ := Search("C4:A8:1D:73:D7:8C")
	if manuf == "D-Link International" {
		t.Log("D-Link International MAC address found.")
		return
	}
	t.Errorf("Found %s instead of D-Link International!", manuf)
}

func TestNetgear(t *testing.T) {
	manuf, _ := Search("9C:D3:6D:9A:CA:81")
	if manuf == "Netgear" {
		t.Log("Netgear MAC address found.")
		return
	}
	t.Errorf("Found %s instead of Netgear!", manuf)
}

func TestShanghaiBroadwanCommunications(t *testing.T) {
	manuf, _ := Search("40:ED:98:6F:DB:AC")
	if manuf == "Shanghai Broadwan Communications Co.,Ltd" {
		t.Log("Shanghai Broadwan Communications Co.,Ltd MAC address found.")
		return
	}
	t.Errorf("Found %s instead of Shanghai Broadwan Communications Co.,Ltd!", manuf)
}

func TestPiranhaEMS(t *testing.T) {
	manuf, _ := Search("70:B3:D5:8C:CD:BE")
	if manuf == "Piranha EMS Inc." {
		t.Log("Piranha EMS Inc. MAC address found.")
		return
	}
	t.Errorf("Found %s instead of Piranha EMS Inc.!", manuf)
}

func TestIEEERegistrationAuthority(t *testing.T) {
	manuf, _ := Search("3C:24:F0:F0:BE:CF")
	if manuf == "IEEE Registration Authority" {
		t.Log("IEEE Registration Authority MAC address found.")
		return
	}
	t.Errorf("Found %s instead of IEEE Registration Authority!", manuf)
}

func TestSamsungElectronics(t *testing.T) {
	manuf, _ := Search("24:FC:E5:AD:BB:89")
	if manuf == "Samsung Electronics Co.,Ltd" {
		t.Log("Samsung Electronics Co.,Ltd MAC address found.")
		return
	}
	t.Errorf("Found %s instead of Samsung Electronics Co.,Ltd!", manuf)
}

func TestInvalidAddress(t *testing.T) {
	manuf, _ := Search("G4:FC:E5:AD:BB:89")
	if manuf == "Invalid MAC address" {
		t.Log("Invalid MAC address MAC address found.")
		return
	}
	t.Errorf("Found %s instead of Invalid MAC address!", manuf)
}
