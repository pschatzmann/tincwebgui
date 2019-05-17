package service

import (
	"testing"
)

var in = `
docker id 1d6c0791469a at unknown port unknown cipher 0 digest 0 maclength 0 compression 0 options 0 status 0800 nexthop - via - distance 0 pmtu 9018 (min 0 max 9018) rx 0 0 tx 0 0
mcbookair id a3318887e521 at MYSELF port 655 cipher 0 digest 0 maclength 0 compression 0 options 700000c status 0050 nexthop mcbookair via mcbookair distance 0 pmtu 9018 (min 0 max 9018) rx 0 0 tx 0 0
nuc id 80421724e266 at unknown port unknown cipher 0 digest 0 maclength 0 compression 0 options 0 status 0000 nexthop - via - distance 0 pmtu 9018 (min 0 max 9018) rx 0 0 tx 0 0
`

func TestParse(t *testing.T) {
	result := Parse(nodeFields, in)
	if len(result) != 3 {
		t.Errorf("The parsing returned a wrong number of lines")
	}

	if result[0]["tx"] != "0 0" {
		t.Errorf("The parsing returned a wrong tx entry")
	}

	if result[0]["name"] != "docker" {
		t.Errorf("The parsing returned a wrong name entry")
	}

}
