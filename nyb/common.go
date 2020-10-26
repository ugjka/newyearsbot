package nyb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

//TZ holds info for Time Zone
type TZ struct {
	Countries []struct {
		Name   string   `json:"name"`
		Cities []string `json:"cities"`
	} `json:"countries"`
	Offset float64 `json:"offset"`
}

func (t TZ) String() (x string) {
	for i, country := range t.Countries {
		x += fmt.Sprintf("%s", country.Name)
		for i, city := range country.Cities {
			if i == 0 {
				x += " ("
			}
			x += fmt.Sprintf("%s", city)
			if i < len(country.Cities)-1 {
				x += ", "
			}
			if i == len(country.Cities)-1 {
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
	return t[i].Offset < t[j].Offset
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
	for _, dt := range strings.Split(value, ",") {
		*i = append(*i, strings.TrimSpace(dt))
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

//NominatimResult ...
type NominatimResult struct {
	Lat         float64
	Lon         float64
	DisplayName string `json:"Display_name"`
}

//UnmarshalJSON ...
func (n *NominatimResult) UnmarshalJSON(data []byte) (err error) {
	v := struct {
		Lat         string
		Lon         string
		DisplayName string `json:"Display_name"`
	}{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		return
	}
	n.Lat, err = strconv.ParseFloat(v.Lat, 64)
	if err != nil {
		return
	}
	n.Lon, err = strconv.ParseFloat(v.Lon, 64)
	if err != nil {
		return
	}
	n.DisplayName = v.DisplayName
	return
}

//NominatimResults ...
type NominatimResults []NominatimResult

//NominatimEndpoint api url
const NominatimEndpoint = "/search?"

//cache and client
var nominatim = struct {
	cache map[string][]byte
	sync.RWMutex
	http.Client
}{
	cache: make(map[string][]byte),
}

//NominatimFetcher make an api request
func NominatimFetcher(url string) (data []byte, err error) {
	nominatim.RLock()
	if v, ok := nominatim.cache[url]; ok {
		nominatim.RUnlock()
		return v, nil
	}
	nominatim.RUnlock()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	get, err := nominatim.Do(req)
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
	nominatim.Lock()
	nominatim.cache[url] = data
	nominatim.Unlock()
	return
}
