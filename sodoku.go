package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

var digits = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var digStr = strings.Join(digits, "")
var rows = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

var cols = digits
var squares = cross(rows, cols)
var unitList = getUnitList()
var units = getUnits()
var peers = getPeers()

func cross(a, b []string) []string {
	crossPr := []string{}

	for _, valA := range a {
		for _, valB := range b {
			crossPr = append(crossPr, fmt.Sprint(valA, valB))
		}
	}
	return crossPr
}

// rows, cols
func getUnitList() [][]string {
	uList := make([][]string, len(rows)*3)

	p := 0
	for _, cVal := range cols {
		cr := cross(rows, []string{cVal})
		uList[p] = cr
		p++
	}

	for _, rVal := range rows {
		cr := cross([]string{rVal}, cols)
		uList[p] = cr
		p++
	}

	rs := []string{`A B C`, `D E F`, `G H I`}
	cs := []string{`1 2 3`, `4 5 6`, `7 8 9`}

	for _, rVal := range rs {
		for _, cVal := range cs {
			// for i := 0; i < len(rs); i++ {
			// 	for j := 0; i < len(cs); i++ {
			// 		rVal, cVal := string(rs[i]), string(cs[j])

			cr := cross(strings.Fields(rVal), strings.Fields(cVal))
			uList[p] = cr
			p++
		}
	}

	return uList
}

// squares, unitList
func getUnits() map[string][][]string {
	uns := make(map[string][][]string)

	for _, s := range squares {
		un := make([][]string, 3)
		i := 0
		for _, u := range unitList {
			for _, uu := range u {
				if s == uu {
					un[i] = u
					uns[s] = un
					i++
					break
				}
			}

		}
	}

	return uns
}

func getPeers() map[string][]string {

	peerSize := 20

	unions := make(map[string][]string)
	for _, s := range squares {
		union := make(map[string]interface{}, peerSize)
		// union[sVal]
		for _, u := range units[s] {
			for _, uu := range u {
				if s != uu {
					union[uu] = 1
				}
			}
		}

		prs := []string{}
		for val := range union {
			prs = append(prs, val)
		}
		unions[s] = prs
	}

	return unions
}

// ParseGrid : Convert grid to a dict of possible values, {square: digits},
// or return False if a contradiction is detected.
func ParseGrid(grid string) map[string]string {
	values := make(map[string]string, len(squares))
	for _, s := range squares {
		values[s] = digStr
	}
	gridNew, err := gridValues(grid)
	if err != nil {
		log.Fatal(err)
	}

	for s, d := range gridNew {

		if strings.Contains(digStr, d) {
			//&& err != nil {
			err := assign(values, s, d)
			if err != nil {
				return nil
			}

		}
	}

	return values
}

// Convert grid into a dict of {square: char} with '0' or '.' for empties.
func gridValues(grid string) (map[string]string, error) {
	chars := []string{}

	for i := 0; i < len(grid); i++ {
		test := string(grid[i : i+1])
		if strings.Contains(digStr, test) || test == "0" || test == "." {
			chars = append(chars, test)
		}
	}

	if len(chars) != 81 {
		return nil, errors.New("Grid provided not of length 81")
	}

	gridNew := make(map[string]string)
	i := 0
	for _, s := range squares {
		gridNew[s] = chars[i]
		i++
	}

	return gridNew, nil
}

func assign(values map[string]string, s string, d string) error {
	otherValues := strings.Replace(values[s], d, "", -1)
	for _, d2 := range otherValues {
		if err := eliminate(values, s, string(d2)); err != nil {
			return err // TODO: should this exact error be thrown?
		}
	}

	return nil

}

// eliminate d from values[s]; propagate when values or places <= 2.
// Return values, except return False if a contradiction is detected
func eliminate(values map[string]string, s string, d string) error {

	if !strings.Contains(values[s], d) {
		return nil // Already elminated
	}
	values[s] = strings.Replace(values[s], d, "", -1)

	// (1) If a square s is reduced to one value d2, then eliminate d2 from the peers.
	if len(values[s]) == 0 {
		return errors.New("Removed last value from values map")
	} else if len(values[s]) == 1 {
		d2 := values[s]
		allElminiated := true
		for _, s2 := range peers[s] {
			err := eliminate(values, s2, d2)
			if err != nil {
				allElminiated = false
			}
		}
		if !allElminiated {
			return errors.New("Not all eliminated")
		}
		// d2 := values[s]
		// for _, s2 := range peers[s] {
		// 	if err := eliminate(values, s2, d2); err != nil {
		// 		return errors.New("Not all eliminated")
		// 	}
		// }
	}

	// (2) If a unit u is reduced to only one place for a value d, then put it there.

	for _, u := range units[s] {
		dplaces := []string{}
		for _, sVal := range u {
			if strings.Contains(values[sVal], d) {
				dplaces = append(dplaces, sVal)
			}
		}

		if len(dplaces) == 0 {
			return fmt.Errorf("Contradiction, no place for this value %s", d)
		}

		if len(dplaces) == 1 {
			if err := assign(values, dplaces[0], d); err != nil {
				return errors.New("d can only be in one place in unit, assign it there")
			}
			// TODO: might need val... check out..
		}
	}

	return nil
}

func Solve(grid string) map[string]string {
	return Search(ParseGrid(grid))
}

// Using depth-first search and propagation, try all possible values.
func Search(values map[string]string) map[string]string {
	//fmt.Println("deep:",deep)
	if values == nil {
		return nil
	}
	solved := true
	for s, _ := range values {
		if len(values[s]) != 1 {
			solved = false
		}
	}
	if solved {
		return values
	}
	// Chose the unfilled square s with the fewest possibilities
	min := len(digits) + 1
	sq := ""
	for _, s := range squares {
		l := len(values[s])
		if l > 1 {
			if l < min {
				sq = s
				min = l
			}
		}
	}

	for _, d := range values[sq] {

		testValues := copyMap(values)
		err := assign(testValues, sq, string(d))
		if err != nil {
			continue
		}

		val := checkNoZero(Search(testValues))

		if val == nil {
			continue
		}
		return val
	}
	return nil
}

func checkNoZero(test map[string]string) map[string]string {
	for s := range test {
		if s == "" {
			return nil
		}
	}

	return test
}

func copyMap(values map[string]string) map[string]string {
	copyVal := make(map[string]string, len(values))
	for k, v := range values {
		copyVal[k] = v
	}

	return copyVal
}

// Display ...
func Display(values map[string]string) {
	for r, row := range rows {
		for c, col := range digits {
			if c == 3 || c == 6 {
				fmt.Printf("| ")
			}
			fmt.Printf("%v ", values[string(row)+string(col)])
		}
		fmt.Println()
		if r == 2 || r == 5 {
			fmt.Println("------+-------+-------")
		}
	}
}
