package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//Geocode is Google maps api
const Geocode = "http://maps.googleapis.com/maps/api/geocode/json?"

//Timezone is Google maps time api
const Timezone = "https://maps.googleapis.com/maps/api/timezone/json?"

//Gmap holds map json
type Gmap struct {
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

//Gtime holds location tz info
type Gtime struct {
	Status       string
	RawOffset    int
	DstOffset    int
	TimeZoneId   string
	TimeZoneName string
}

//TZ holds infor for Time Zone
type TZ struct {
	Countries []struct {
		Name   string   `json:"name"`
		Cities []string `json:"cities"`
	} `json:"countries"`
	Offset string `json:"offset"`
}

func (t TZ) String() (x string) {
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

//TZS is a slice of TZ
type TZS []TZ

func (t TZS) Len() int {
	return len(t)
}

func (t TZS) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TZS) Less(i, j int) bool {
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

//Getter Gets "GET" DATA
func Getter(url string) (data []byte, err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "github.com/ugjka/newyearsbot")
	get, err := client.Do(req)
	if err != nil {
		return
	}
	defer get.Body.Close()
	if get.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status: %d", get.StatusCode)
	}
	data, err = ioutil.ReadAll(get.Body)
	if err != nil {
		return
	}
	return
}

//IrcChans is a custom flag
type IrcChans []string

func (i *IrcChans) String() string {
	return fmt.Sprint(*i)
}

//Set satisfies flag Interface?
func (i *IrcChans) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ", ") {
		*i = append(*i, dt)
	}
	return nil
}

//Timer struct
type Timer struct {
	C      chan bool
	Target time.Time
	stop   chan bool
	ticker *time.Ticker
}

//NewTimer returns ticker based timer
//We need this to take into account time taken in suspend and hibernation
//Golang time.Timers suck
func NewTimer(dur time.Duration) *Timer {
	t := &Timer{}
	t.C = make(chan bool)
	t.stop = make(chan bool)
	t.Target = time.Now().UTC().Add(dur)
	t.ticker = time.NewTicker(time.Millisecond * 100)
	go func(t *Timer) {
		defer t.ticker.Stop()
		for range t.ticker.C {
			select {
			case <-t.stop:
				return
			default:
				if time.Now().UTC().After(t.Target) {
					close(t.C)
					return
				}
			}
		}
	}(t)
	return t
}

//Stop stops the timer
func (t *Timer) Stop() {
	select {
	case <-t.stop:
		return
	default:
		close(t.stop)
	}
}
