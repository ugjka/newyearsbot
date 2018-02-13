package nyb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hako/durafmt"
	gotz "github.com/ugjka/go-tz"
	c "github.com/ugjka/newyearsbot/common"
)

//Func for querying newyears in specified location
func getNewYear(loc string, email string, server string) (string, error) {
	var adress string
	log.Println("Querying location:", loc)
	maps := url.Values{}
	maps.Add("q", loc)
	maps.Add("format", "json")
	maps.Add("accept-language", "en")
	maps.Add("limit", "1")
	maps.Add("email", email)
	data, err := c.NominatimGetter(server + c.NominatimGeoCode + maps.Encode())
	if err != nil {
		log.Println(err)
		return "", err
	}
	var mapj c.NominatimResults
	if err = json.Unmarshal(data, &mapj); err != nil {
		log.Println(err)
		return "", err
	}
	if len(mapj) == 0 {
		return "Couldn't find that place.", nil
	}
	adress = mapj[0].DisplayName
	lat, err := strconv.ParseFloat(mapj[0].Lat, 64)
	if err != nil {
		return "", err
	}
	lon, err := strconv.ParseFloat(mapj[0].Lon, 64)
	if err != nil {
		return "", err
	}
	p := gotz.Point{
		Lat: lat,
		Lng: lon,
	}
	zone, err := gotz.GetZone(p)
	if err != nil {
		return "Couldn't get the timezone for that location.", nil
	}
	//RawOffset
	offset, err := time.ParseDuration(fmt.Sprintf("%ds", getOffset(target, zone)))
	if err != nil {
		log.Println(err)
		return "", err
	}
	//Check if past target
	if time.Now().UTC().Add(offset).Before(target) {
		humandur, err := durafmt.ParseString(target.
			Sub(time.Now().UTC().Add(offset)).String())
		if err != nil {
			log.Println(err)
			return "", err
		}
		return fmt.Sprintf("New Year in %s will happen in %s", adress,
			removeMilliseconds(humandur.String())), nil
	}
	humandur, err := durafmt.ParseString(time.Now().UTC().Add(offset).Sub(target).String())
	if err != nil {
		log.Println(err)
		return "", err
	}
	return fmt.Sprintf("New Year in %s happened %s ago",
		adress, removeMilliseconds(humandur.String())), nil
}

func removeMilliseconds(dur string) string {
	arr := strings.Split(dur, " ")
	if len(arr) < 3 {
		return dur
	}
	return strings.Join(arr[:len(arr)-2], " ")
}

func getOffset(target time.Time, zone *time.Location) int {
	_, offset := time.Date(target.Year(), target.Month(), target.Day(),
		target.Hour(), target.Minute(), target.Second(),
		target.Nanosecond(), zone).Zone()
	return offset
}

func pingpong(c chan bool) {
	select {
	case c <- true:
	default:
		return
	}
}
