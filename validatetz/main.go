//This utility validates Time Zone dataset against google maps api, double check the results, sometimes false positives
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"sort"
	"strconv"
	"time"

	c "github.com/ugjka/newyearsbot/common"
)

//Set target year
var target = func() time.Time {
	tmp := time.Now().UTC()
	if tmp.Month() == time.January && tmp.Day() < 2 {
		return time.Date(tmp.Year()-1, time.January, 1, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(tmp.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
}()

func main() {
	tzdatapath := flag.String("tzpath", "../tz.json", "path to tz.json")
	//Check if tz.json exists
	if _, err := os.Stat(*tzdatapath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: file %s does not exist\n", *tzdatapath)
		os.Exit(1)
	}
	var zones c.TZS
	file, err := os.Open(*tzdatapath)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(content, &zones); err != nil {
		log.Fatal(err)
	}
	//print target to be sure
	fmt.Println("Target:", target)
	sort.Sort(sort.Reverse(zones))
	for _, k := range zones {
		for _, k2 := range k.Countries {
			if len(k2.Cities) == 0 {
				res, err := getTimeZone(k2.Name)
				time.Sleep(time.Second * 2)
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
				time.Sleep(time.Second * 2)
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
	maps.Add("address", loc)
	maps.Add("sensor", "false")
	maps.Add("language", "en")
	data, err := c.Getter(c.Geocode + maps.Encode())
	if err != nil {
		return "", err
	}
	var mapj c.Gmap
	if err = json.Unmarshal(data, &mapj); err != nil {
		return "", err
	}
	if mapj.Status != "OK" {
		return "", errors.New(loc + " Status not OK")
	}
	location := fmt.Sprintf("%.7f,%.7f", mapj.Results[0].Geometry.Location.Lat, mapj.Results[0].Geometry.Location.Lng)
	tmzone := url.Values{}
	tmzone.Add("location", location)
	tmzone.Add("timestamp", fmt.Sprintf("%d", target.Unix()))
	tmzone.Add("sensor", "false")
	data, err = c.Getter(c.Timezone + tmzone.Encode())
	if err != nil {
		return "", err
	}
	var timej c.Gtime
	if err = json.Unmarshal(data, &timej); err != nil {
		return "", err
	}
	if timej.Status != "OK" {
		return "", errors.New(loc + " Couldn't get timezone info.")
	}
	var offset float64
	offset = (float64(timej.RawOffset) + float64(timej.DstOffset)) / 3600.0
	return fmt.Sprintf("%f", offset), nil
}
