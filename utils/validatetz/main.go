//This utility validates Time Zone dataset against osm database and tz shapefile,
//double check the results, sometimes false positives
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/rhinosf1/newyearsbot/nyb"
	"gopkg.in/ugjka/go-tz.v2/tz"
)

//Set target year
var target = func() time.Time {
	tmp := time.Now().UTC()
	if tmp.Month() == time.January && tmp.Day() < 2 {
		return time.Date(tmp.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(tmp.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
}()

var email *string
var nominatim *string

func main() {
	email = flag.String("email", "", "nominatim email")
	nominatim = flag.String("nominatim", "http://nominatim.openstreetmap.org", "nominatim server")
	flag.Parse()
	if *email == "" {
		fmt.Fprintf(os.Stderr, "%s", "provide email with -email flag\n")
		return
	}
	var zones nyb.TZS
	if err := json.Unmarshal(nyb.Zones, &zones); err != nil {
		log.Fatal(err)
	}
	//print target to be sure
	fmt.Println("Target:", target)
	sort.Sort(sort.Reverse(zones))
	for _, zone := range zones {
		fmt.Println("Zone:", zone.Offset)
		for _, country := range zone.Countries {
			if len(country.Cities) == 0 {
				remoteOffset, err := timeZone(country.Name)
				if err != nil {
					log.Println(country.Name, err)
				} else {
					if remoteOffset != zone.Offset {
						fmt.Printf("%s: Offset mismatch; Local: %v, Remote: %v\n",
							country.Name, zone.Offset, remoteOffset)
					}
				}
			}
			for _, city := range country.Cities {
				remoteOffset, err := timeZone(city + ", " + country.Name)
				if err != nil {
					log.Println(city+", "+country.Name, err)
				} else {
					if remoteOffset != zone.Offset {
						fmt.Printf("%s, %s: Offset mismatch; Local: %v, Remote: %v\n",
							city, country.Name, zone.Offset, remoteOffset)
					}
				}
			}
		}
	}
}

//Get Timezone Offset
func timeZone(location string) (float64, error) {
	time.Sleep(time.Second * 2)
	data, err := nyb.NominatimFetcher(email, nominatim, &location)
	if err != nil {
		return 0, err
	}
	var mapj nyb.NominatimResults
	if err = json.Unmarshal(data, &mapj); err != nil {
		return 0, err
	}
	if len(mapj) == 0 {
		return 0, errors.New("could not find location")
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
