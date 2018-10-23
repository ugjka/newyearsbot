//This utility validates Time Zone dataset against osm database and tz shapefile,
//double check the results, sometimes false positives
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"sort"
	"time"

	"github.com/ugjka/go-tz"
	"github.com/ugjka/newyearsbot/nyb"
)

//Set target year
var target = func() time.Time {
	tmp := time.Now().UTC()
	if tmp.Month() == time.January && tmp.Day() < 2 {
		return time.Date(tmp.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(tmp.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
}()

var ircEmail *string
var ircNominatim *string

func main() {
	ircEmail = flag.String("email", "", "Email for Open Street Map")
	ircNominatim = flag.String("nominatim", "http://nominatim.openstreetmap.org", "Nominatim server to use")
	flag.Parse()
	if *ircEmail == "" {
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
				remoteOffset, err := getTimeZone(country.Name)
				if err != nil {
					log.Println(country.Name, err)
				} else {
					if remoteOffset != zone.Offset {
						fmt.Printf("%s: Offset mismatch Loc: %v, Rem: %v\n",
							country.Name, zone.Offset, remoteOffset)
					}
				}
			}
			for _, city := range country.Cities {
				remoteOffset, err := getTimeZone(city + " " + country.Name)
				if err != nil {
					log.Println(city+" "+country.Name, err)
				} else {
					if remoteOffset != zone.Offset {
						fmt.Printf("%s, %s: Offset mismatch Loc: %v, Rem: %v\n",
							city, country.Name, zone.Offset, remoteOffset)
					}
				}
			}
		}
	}
}

//Get Timezone Offset
func getTimeZone(loc string) (float64, error) {
	maps := url.Values{}
	maps.Add("q", loc)
	maps.Add("format", "json")
	maps.Add("accept-language", "en")
	maps.Add("limit", "1")
	maps.Add("email", *ircEmail)
	var data []byte
	var err error
	time.Sleep(time.Second * 2)
	data, err = nyb.NominatimGetter(*ircNominatim + nyb.NominatimEndpoint + maps.Encode())
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
	location := gotz.Point{
		Lat: mapj[0].Lat,
		Lon: mapj[0].Lon,
	}
	zone, err := gotz.GetZone(location)
	if err != nil {
		return 0, err
	}
	offset := getOffset(target, zone)
	return float64(offset) / 60 / 60, nil
}

func getOffset(target time.Time, zone *time.Location) int {
	_, offset := time.Date(target.Year(), target.Month(), target.Day(),
		target.Hour(), target.Minute(), target.Second(),
		target.Nanosecond(), zone).Zone()
	return offset
}
