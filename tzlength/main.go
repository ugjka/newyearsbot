package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
)

type tz struct {
	Countries []struct {
		Name   string   `json:"name"`
		Cities []string `json:"cities"`
	} `json:"countries"`
	Offset string `json:"offset"`
}

func (t tz) String() (x string) {
	for i, k := range t.Countries {
		x += fmt.Sprintf("%s", k.Name)
		for i, k1 := range k.Cities {
			if k1 == "" {
				continue
			}
			if i == 0 {
				x += " ("
			}
			x += fmt.Sprintf("%s", k1)
			if i >= 0 && i < len(k.Cities)-1 {
				x += ", "
			}
			if i == len(k.Cities)-1 {
				x += ")"
			}
		}
		if i < len(t.Countries)-1 {
			x += ", "
		}
	}
	return
}

type tzs []tz

func (t tzs) Len() int {
	return len(t)
}

func (t tzs) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t tzs) Less(i, j int) bool {
	x, err := strconv.ParseFloat(t[i].Offset, 64)
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.ParseFloat(t[j].Offset, 64)
	if err != nil {
		log.Fatal(err)
	}
	if x < y {
		return true
	}
	return false
}

func main() {
	var zones tzs
	file, err := os.Open("../tz.json")
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(content, &zones)
	sort.Sort(sort.Reverse(zones))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for _, k := range zones {
		tmp := len(k.String())
		//if tmp > 468 {
		if tmp > 0 {
			fmt.Fprintf(w, "%s\t%d\t\n", k.Offset, tmp)
		}
	}
	w.Flush()
}
