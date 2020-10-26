// Interactive utility for querying location info
// Reads lines from stdin prints to stdout
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/ugjka/newyearsbot/nyb"
	"gopkg.in/ugjka/go-tz.v2/tz"
)

var email *string
var ircNominatim *string
var ext *string

//Set target year
var target = func() time.Time {
	tmp := time.Now().UTC()
	if tmp.Month() == time.January && tmp.Day() < 2 {
		return time.Date(tmp.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(tmp.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
}()

func main() {
	ext = flag.String("ext", "", "external geojson")
	email = flag.String("email", "", "Email for Open Street Map")
	ircNominatim = flag.String("nominatim", "http://nominatim.openstreetmap.org", "Nominatim server to use")
	flag.Parse()
	if *email == "" {
		fmt.Fprintf(os.Stderr, "%s", "provide email with -email flag\n")
		return
	}
	if *ext != "" {
		f, err := os.OpenFile(*ext, os.O_RDONLY, 0655)
		if err != nil {
			panic(err)
		}
		tz.LoadGeoJSON(f)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result, err := locationInfo(scanner.Text())
		if err == nil {
			fmt.Printf("%s\n", result)
		} else {
			fmt.Printf("%s, Error: %s\n", scanner.Text(), err)
		}
	}
}

func locationInfo(loc string) (string, error) {
	maps := url.Values{}
	maps.Add("q", loc)
	maps.Add("format", "json")
	maps.Add("accept-language", "en")
	maps.Add("limit", "1")
	maps.Add("email", *email)
	var data []byte
	var err error
	data, err = nyb.NominatimFetcher(*ircNominatim + nyb.NominatimEndpoint + maps.Encode())
	if err != nil {
		return "", err
	}

	var mapj nyb.NominatimResults
	if err = json.Unmarshal(data, &mapj); err != nil {
		return "", err
	}
	if len(mapj) == 0 {
		return "", errors.New("status not OK")
	}
	location := tz.Point{
		Lat: mapj[0].Lat,
		Lon: mapj[0].Lon,
	}
	now := time.Now()
	tzid, err := tz.GetZone(location)
	lookup := time.Now().Sub(now)
	if err != nil {
		return "", err
	}
	zone, err := time.LoadLocation(tzid[0])
	if err != nil {
		return "", err
	}
	offset := zoneOffset(target, zone)
	return fmt.Sprintf("%s, Offset %v, zone: %s, time now: %s, tz dur %s", mapj[0].DisplayName, float64(offset)/60/60, zone, time.Now().In(zone), lookup), nil
}

func zoneOffset(target time.Time, zone *time.Location) int {
	_, offset := time.Date(target.Year(), target.Month(), target.Day(),
		target.Hour(), target.Minute(), target.Second(),
		target.Nanosecond(), zone).Zone()
	return offset
}
