package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"time"

	c "github.com/ugjka/newyearsbot/common"
)

var target = time.Date(2017, time.December, 31, 0, 0, 0, 0, time.UTC)

func main() {
	var zones c.TZS
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
	client := http.Client{}
	maps := url.Values{}
	maps.Add("address", loc)
	maps.Add("sensor", "false")
	maps.Add("language", "en")
	req, err := http.NewRequest("GET", c.Geocode+maps.Encode(), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla")
	get, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer get.Body.Close()
	text, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return "", err
	}
	var mapj c.Gmap
	json.Unmarshal(text, &mapj)
	if mapj.Status != "OK" {
		return "", errors.New(loc + " Status not OK")
	}
	tmzone := url.Values{}
	location := fmt.Sprintf("%.6f,%.6f", mapj.Results[0].Geometry.Location.Lat, mapj.Results[0].Geometry.Location.Lng)
	tmzone.Add("location", location)
	tmzone.Add("timestamp", fmt.Sprintf("%d", target.Unix()))
	tmzone.Add("sensor", "false")

	req2, err := http.NewRequest("GET", c.Timezone+tmzone.Encode(), nil)
	if err != nil {
		return "", err
	}
	req2.Header.Set("User-Agent", "Mozilla")
	get2, err := client.Do(req2)
	if err != nil {
		return "", err
	}
	defer get2.Body.Close()
	text, err = ioutil.ReadAll(get2.Body)
	if err != nil {
		return "", err
	}
	var timej c.Gtime
	json.Unmarshal(text, &timej)
	if timej.Status != "OK" {
		return "", errors.New(loc + " Couldn't get timezone info.")
	}
	var offset float64
	offset = (float64(timej.RawOffset) + float64(timej.DstOffset)) / 3600.0
	return fmt.Sprintf("%f", offset), nil
}
