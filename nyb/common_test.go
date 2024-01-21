package nyb

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIrcChans(t *testing.T) {
	var ch Channels
	cases := "#test1 #test2 #test3"
	if err := ch.Set(cases); err != nil {
		t.Error(err)
	}
	if !(ch.String() == fmt.Sprintf("%v", ch)) {
		t.Error("channels string doesn't match")
	}
	if err := ch.Set("already set"); err == nil {
		t.Error("did not catch already set")
	}
}

func TestNominatimUnmarshal(t *testing.T) {
	type Re struct {
		Lat         string
		Lon         string
		DisplayName string `json:"Display_name"`
	}
	type Res []Re
	v := make(Res, 0)
	v = append(v, Re{
		Lat:         "fail",
		Lon:         "1.00",
		DisplayName: "test",
	})
	//Lat fail
	data, err := json.Marshal(v)
	if err != nil {
		t.Error("could not marshal the test case")
	}
	test := make(NominatimResults, 0)
	err = json.Unmarshal(data, &test)
	if err == nil {
		t.Error("lat did not fail")
	}
	//Lon fail
	v[0].Lat = "1.00"
	v[0].Lon = "fail"
	data, err = json.Marshal(v)
	if err != nil {
		t.Error("could not marshal the test case")
	}
	test = make(NominatimResults, 0)
	err = json.Unmarshal(data, &test)
	if err == nil {
		t.Error("lon did not fail")
	}
}
