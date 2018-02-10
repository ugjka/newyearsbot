//This utility validates Time Zone dataset against google maps api, double check the results, sometimes false positives
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
	"strconv"
	"time"

	"github.com/ugjka/go-tz"

	c "github.com/ugjka/newyearsbot/common"
	nyb "github.com/ugjka/newyearsbot/nyb"
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

func main() {
	ircEmail = flag.String("email", "", "Email for Open Street Map")
	flag.Parse()
	if *ircEmail == "" {
		fmt.Fprintf(os.Stderr, "%s", "provide email with -email flag")
		return
	}
	var zones c.TZS
	if err := json.Unmarshal([]byte(nyb.TZ), &zones); err != nil {
		log.Fatal(err)
	}
	//print target to be sure
	fmt.Println("Target:", target)
	sort.Sort(sort.Reverse(zones))
	for _, k := range zones {
		fmt.Println("Zone:", k.Offset)
		for _, k2 := range k.Countries {
			if len(k2.Cities) == 0 {
				res, err := getTimeZone(k2.Name)
				time.Sleep(time.Second * 5)
				if err != nil {
					log.Println(err)
				} else {
					res, _ := strconv.ParseFloat(res, 64)
					koff, _ := strconv.ParseFloat(k.Offset, 64)
					if res != koff {
						fmt.Println(k2.Name, k.Offset, res)
					}
				}
			}
			for _, k3 := range k2.Cities {
				res, err := getTimeZone(k2.Name + " " + k3)
				time.Sleep(time.Second * 5)
				if err != nil {
					log.Println(err)
				} else {
					res, _ := strconv.ParseFloat(res, 64)
					koff, _ := strconv.ParseFloat(k.Offset, 64)
					if res != koff {
						fmt.Println(k2.Name, k3, k.Offset, res)
					}
				}
			}
		}
	}
}

func getTimeZone(loc string) (string, error) {
	maps := url.Values{}
	maps.Add("q", loc)
	maps.Add("format", "json")
	maps.Add("accept-language", "en")
	maps.Add("limit", "1")
	maps.Add("email", *ircEmail)
	data, err := c.OSMGetter(c.OSMGeocode + maps.Encode())
	if err != nil {
		return "", err
	}
	var mapj c.OSMmapResults
	if err = json.Unmarshal(data, &mapj); err != nil {
		return "", err
	}
	if len(mapj) == 0 {
		return "", errors.New(loc + " Status not OK")
	}
	lat, err := strconv.ParseFloat(mapj[0].Lat, 64)
	if err != nil {
		return "", err
	}
	lon, err := strconv.ParseFloat(mapj[0].Lon, 64)
	if err != nil {
		return "", err
	}
	location := gotz.Point{
		Lat: lat,
		Lng: lon,
	}
	zone, err := gotz.GetZone(location)
	if err != nil {
		return "", err
	}
	_, offset := time.Date(target.Year(), target.Month(), target.Day(), target.Hour(), target.Minute(),
		target.Second(), target.Nanosecond(), zone).Zone()
	return fmt.Sprintf("%f", float64(offset)/60/60), nil
}
