package facade

import "testing"

var expect = "AModule do some things and BModule do some things"

func TestFacadeApi(t *testing.T) {
	api := NEWAPI()
	res := api.DoThing()
	if res != expect {
		t.Fatalf("expect %s, return %s", expect, res)
	}
}
