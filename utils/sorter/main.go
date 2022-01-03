//Util for sorting the TZ json
//To make it neat and nice after manual edits
package main

import (
	"encoding/json"
	"os"
	"sort"

	nyb "github.com/rhinosf1/newyearsbot/nyb"
)

func main() {
	var v TZS
	err := json.Unmarshal(nyb.Zones, &v)
	if err != nil {
		panic(err)
	}
	sort.Sort(v)
	for i := range v {
		sort.Sort(v[i])
		for j := range v[i].Countries {
			sort.Sort(v[i].Countries[j].Cities)
		}
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(v)
}

//Cities ...
type Cities []string

//TZ holds infor for Time Zone
//TZ holds info for Time Zone
type TZ struct {
	Countries []struct {
		Name   string `json:"name"`
		Cities Cities `json:"cities"`
	} `json:"countries"`
	Offset float64 `json:"offset"`
}

//TZS is a slice of TZ
type TZS []TZ

func (t Cities) Len() int {
	return len(t)
}

func (t Cities) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Cities) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t TZ) Len() int {
	return len(t.Countries)
}

func (t TZ) Swap(i, j int) {
	t.Countries[i], t.Countries[j] = t.Countries[j], t.Countries[i]
}

func (t TZ) Less(i, j int) bool {
	return t.Countries[i].Name < t.Countries[j].Name
}

func (t TZS) Len() int {
	return len(t)
}

func (t TZS) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TZS) Less(i, j int) bool {
	return t[i].Offset < t[j].Offset
}
