package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/ugjka/newyearsbot/nyb"
	"gopkg.in/ugjka/go-tz.v2/tz"
)

// scraping https://www.timeanddate.com/time/map/
const DATASET_JSON_URL = "https://c.tadst.com/gfx/tzmap/worldclockcities.en.json"

var email *string
var nominatim *string

func main() {
	_, err := os.Stat("tz.json")
	if err == nil {
		postprocess()
		return
	}
	email = flag.String("email", "", "nominatim email")
	nominatim = flag.String("nominatim", "http://nominatim.openstreetmap.org", "nominatim server")
	flag.Parse()
	if *email == "" {
		fmt.Fprintf(os.Stderr, "%s", "provide email with -email flag\n")
		return
	}
	var p places
	resp, err := http.Get(DATASET_JSON_URL)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
	var zones nyb.TZS
	for i, loc := range p.Places {
		if loc.Country == "Western Sahara" {
			continue
		}
		fmt.Printf("%d ", len(p.Places)-i)
		loc.Country = fixes(loc.Country)
		loc.Name = fixes(loc.Name)
		remoteOffset, err := timeZone(loc.Country, loc.Name)
		if err != nil {
			fmt.Printf("\n%s, %s: %s\n", loc.Country, loc.Name, err)
			continue
		}
		if !zones.Exists(remoteOffset, loc.Country, loc.Name) {
			zones = zones.Insert(remoteOffset, loc.Country, loc.Name)
		}
		time.Sleep(time.Second)
	}
	data, err := json.Marshal(zones)
	if err != nil {
		log.Fatal(err)
	}
	data = sorter(data)
	err = os.WriteFile("tz.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func sorter(data []byte) []byte {
	var v TZS
	err := json.Unmarshal(data, &v)
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
	data, err = json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// POST PROCESSING
func postprocess() {
	data, err := os.ReadFile("tz.json")
	if err != nil {
		log.Fatal(err)
	}
	fresh := nyb.TZS{}
	json.Unmarshal(data, &fresh)
	old := nyb.TZS{}
	json.Unmarshal(nyb.Zones, &old)
	var diff int
	for i, z := range fresh {
		if old[i+diff].Offset != z.Offset {
			diff++
		}
		if 396-len(old[i+diff].String()) > 63 && z.Offset != 0 {
			for _, co := range z.Countries {
				if len(co.Cities) == 0 {
					if !old.Exists(z.Offset, co.Name, "") {
						old.Insert(z.Offset, co.Name, "")
					}
					continue
				}
				for _, ci := range co.Cities {
					if !old.Exists(z.Offset, co.Name, ci) {
						old.Insert(z.Offset, co.Name, ci)
					}
				}
			}
		} else {
			for _, co := range z.Countries {
				if !old.Exists(z.Offset, co.Name, "") {
					old.Insert(z.Offset, co.Name, "")
				}
			}
		}
		data, err = json.MarshalIndent(old, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		data = sorter(data)
		os.WriteFile("tznew.json", data, 0644)
		os.WriteFile("tz.go.new", []byte(fmt.Sprintf(template, data)), 0644)
	}
	// loop end
	sort.Sort(sort.Reverse(old))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for _, k := range old {
		tmp := len(k.String())
		if tmp > 0 {
			fmt.Fprintf(w, "%v\t%d\t%d\t\n", k.Offset, tmp, tmp-396)
		}
	}
	w.Flush()
	fmt.Println(len(old))
}

// END POST PROCESSING

var template = `package nyb

// Zones contains time zone information in JSON format
var Zones = []byte(` + "`%s`)\n"

type places struct {
	Places []struct {
		Name    string
		State   string
		Country string
	}
}

// Set target year
var target = func() time.Time {
	tmp := time.Now().UTC()
	if tmp.Month() == time.January && tmp.Day() < 2 {
		return time.Date(tmp.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(tmp.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
}()

func fixes(s string) string {
	if s == "US Minor Outlying Islands" {
		return "United States"
	}
	if s == "Midway" {
		return "Midway Atoll"
	}
	if s == "Congo Dem. Rep." {
		return "DRC"
	}
	if s == "South Georgia/Sandwich Is." {
		return "South Georgia and the South Sandwich Islands"
	}
	if s == "USA" {
		return "United States"
	}
	if s == "Central African Republic" {
		return "CAR"
	}
	if s == "Vatican City State" {
		return "Vatican"
	}
	s = strings.Split(s, "(")[0]
	s = strings.TrimSpace(s)
	return s
}

// Get Timezone Offset
func timeZone(country, city string) (float64, error) {
	mapj, err := nyb.NominatimFetcherLong(*email, *nominatim, country, city, "")
	if len(mapj) == 0 || err != nil {
		mapj, err = nyb.NominatimFetcher(*email, *nominatim, country+", "+city)
		if err != nil {
			return 0, err
		}
		if len(mapj) == 0 {
			return 0, fmt.Errorf("no results")
		}
	}
	point := tz.Point{
		Lat: mapj[0].Lat,
		Lon: mapj[0].Lon,
	}
	tzid, err := tz.GetZone(point)
	if err != nil {
		return 0, err
	}
	zone, err := time.LoadLocation(tzid[0])
	if err != nil {
		return 0, err
	}
	offset := zoneOffset(target, zone)
	return float64(offset) / 60 / 60, nil
}

func zoneOffset(target time.Time, zone *time.Location) int {
	_, offset := time.Date(target.Year(), target.Month(), target.Day(),
		target.Hour(), target.Minute(), target.Second(),
		target.Nanosecond(), zone).Zone()
	return offset
}

// Sorting stuff
// Cities ...
type Cities []string

// TZ holds infor for Time Zone
// TZ holds info for Time Zone
type TZ struct {
	Countries []struct {
		Name   string `json:"name"`
		Cities Cities `json:"cities"`
	} `json:"countries"`
	Offset float64 `json:"offset"`
}

// TZS is a slice of TZ
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
