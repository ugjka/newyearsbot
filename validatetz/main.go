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
)

var target = time.Date(2017, time.December, 31, 0, 0, 0, 0, time.UTC)

type tz struct {
	Countries []struct {
		Name   string   `json:"name"`
		Cities []string `json:"cities"`
	} `json:"countries"`
	Offset string `json:"offset"`
}

type tzs []tz

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

const geocode = "http://maps.googleapis.com/maps/api/geocode/json?"
const timezone = "https://maps.googleapis.com/maps/api/timezone/json?"

type gmap struct {
	Results []struct {
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			}
		}
	}
	Status string
}

type gtime struct {
	Status    string
	RawOffset int
	DstOffset int
}

func getTimeZone(loc string) (string, error) {
	client := http.Client{}
	maps := url.Values{}
	maps.Add("address", loc)
	maps.Add("sensor", "false")
	maps.Add("language", "en")
	req, err := http.NewRequest("GET", geocode+maps.Encode(), nil)
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
	var mapj gmap
	json.Unmarshal(text, &mapj)
	if mapj.Status != "OK" {
		return "", errors.New(loc + " Status not OK")
	}
	tmzone := url.Values{}
	location := fmt.Sprintf("%.6f,%.6f", mapj.Results[0].Geometry.Location.Lat, mapj.Results[0].Geometry.Location.Lng)
	tmzone.Add("location", location)
	tmzone.Add("timestamp", fmt.Sprintf("%d", target.Unix()))
	tmzone.Add("sensor", "false")

	req2, err := http.NewRequest("GET", timezone+tmzone.Encode(), nil)
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
	var timej gtime
	json.Unmarshal(text, &timej)
	if timej.Status != "OK" {
		return "", errors.New(loc + " Couldn't get timezone info.")
	}
	var offset float64
	offset = (float64(timej.RawOffset) + float64(timej.DstOffset)) / 3600.0
	return fmt.Sprintf("%f", offset), nil
}

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
