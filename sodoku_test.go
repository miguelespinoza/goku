package main

import "testing"

func TestVariables(t *testing.T) {

	if len(squares) != 81 {
		t.Error("squares size is not 81")
	}

	if len(unitList) != 27 {
		t.Error("unitList size is not 27")
	}

	for _, s := range squares {

		if len(units[s]) != 3 {
			t.Errorf("units for %s, are not of size 3 instead of size %d", s, len(units[s]))
		}
	}

	for _, s := range squares {
		if len(peers[s]) != 20 {
			t.Error("peer size is not 20")
		}
	}

	testUnits := map[string]int{
		"C2": 1, "D2": 1, "E2": 1, "F2": 1, "G2": 1, "H2": 1, "I2": 1,
		"C1": 1, "C3": 1, "C4": 1, "C5": 1, "C6": 1, "C7": 1, "C8": 1, "C9": 1,
		"A1": 1, "A2": 1, "A3": 1, "B1": 1, "B2": 1, "B3": 1,
	}

	for _, uVal := range units["C2"] {

		for _, uuVal := range uVal {

			if _, ok := testUnits[uuVal]; !ok {
				t.Errorf("Error with units for C2 with %s", uuVal)
			}
		}
	}

	testPeers := map[string]int{
		"A2": 1, "B2": 1, "D2": 1, "E2": 1, "F2": 1, "G2": 1, "H2": 1, "I2": 1,
		"C1": 1, "C3": 1, "C4": 1, "C5": 1, "C6": 1, "C7": 1, "C8": 1, "C9": 1,
		"A1": 1, "A3": 1, "B1": 1, "B3": 1,
	}

	for _, pVal := range peers["C2"] {
		if _, ok := testPeers[pVal]; !ok {
			t.Errorf("Error with peer for C2 with %s", pVal)
		}
	}
}
