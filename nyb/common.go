package nyb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

var userAgent = fmt.Sprintf("NYE IRC party bot, v%d: https://github.com/ugjka/newyearsbot", now().Year()+1)

// TZ holds time zone data
type TZ struct {
	Countries []Country `json:"countries"`
	Offset    float64   `json:"offset"`
}

type Country struct {
	Name   string   `json:"name"`
	Cities []string `json:"cities"`
}

func (t TZ) String() (x string) {
	for i, country := range t.Countries {
		x += country.Name
		for i, city := range country.Cities {
			if i == 0 {
				x += " ("
			}
			x += city
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

func (t TZ) Split(max int) (arr []string) {
	var x string
	var prev int
	var total int = max
	for i, country := range t.Countries {
		prev = len(x)
		//x += fmt.Sprintf("\x02%s\x0f", country.Name)
		x += country.Name
		for i, city := range country.Cities {
			if i == 0 {
				x += " ("
			}
			x += city
			if i < len(country.Cities)-1 {
				x += ", "
			}
			if i == len(country.Cities)-1 {
				x += ")"
			}
		}
		if len(x) > total {
			x = x[:prev-1] + "\n" + x[prev:]
			total += max
		}
		if i < len(t.Countries)-1 {
			x += ", "
		}
	}
	return strings.Split(x, "\n")
}

// TZS is a slice of timezones
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

func (t TZS) Exists(offset float64, country, city string) bool {
	for _, tz := range t {
		if tz.Offset == offset {
			for _, c := range tz.Countries {
				if c.Name == country {
					if city == "" {
						return true
					}
					for _, cit := range c.Cities {
						if cit == city {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func (t TZS) Insert(offset float64, country, city string) TZS {
	for i, tz := range t {
		if tz.Offset == offset {
			for j, c := range tz.Countries {
				if c.Name == country {
					t[i].Countries[j].Cities = append(t[i].Countries[j].Cities, city)
					return t
				}
			}
			if country == city || city == "" {
				t[i].Countries = append(t[i].Countries, Country{
					Name:   country,
					Cities: []string{},
				})
				return t
			}
			t[i].Countries = append(t[i].Countries, Country{
				Name:   country,
				Cities: []string{city},
			})
			return t
		}
	}
	if country == city || city == "" {
		t = append(t, TZ{
			Offset: offset,
			Countries: []Country{
				{
					Name:   country,
					Cities: []string{},
				},
			},
		})
		return t
	}
	t = append(t, TZ{
		Offset: offset,
		Countries: []Country{
			{
				Name:   country,
				Cities: []string{city},
			},
		},
	})
	return t
}

// Channels is a flag that parses a list of IRC channels
type Channels []string

func (i *Channels) String() string {
	return fmt.Sprint(*i)
}

// Set satisfies the flag Interface
func (i *Channels) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("channel flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		*i = append(*i, strings.TrimSpace(dt))
	}
	return nil
}

// Timer struct
type Timer struct {
	C      chan bool
	Target time.Time
	stop   chan bool
	ticker *time.Ticker
}

// NewTimer returns a ticker based timer.
// We need this to take into account time taken in suspend, hibernation or if system time is changed.
func NewTimer(dur time.Duration) *Timer {
	t := &Timer{}
	t.C = make(chan bool)
	t.stop = make(chan bool)
	t.Target = now().UTC().Add(dur)
	t.ticker = time.NewTicker(time.Millisecond * 100)
	go func(t *Timer) {
		defer t.ticker.Stop()
		for range t.ticker.C {
			select {
			case <-t.stop:
				return
			default:
				if now().UTC().After(t.Target) {
					close(t.C)
					return
				}
			}
		}
	}(t)
	return t
}

// Stop stops the timer
func (t *Timer) Stop() {
	select {
	case <-t.stop:
		return
	default:
		close(t.stop)
	}
}

// NominatimResult ...
type NominatimResult struct {
	Lat         float64
	Lon         float64
	DisplayName string `json:"Display_name"`
}

// NominatimResults ...
type NominatimResults []NominatimResult

// UnmarshalJSON ...
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

// cache and client for NominatimFetcher
var nominatim = struct {
	cache map[string]NominatimResults
	sync.RWMutex
	http.Client
}{
	cache: make(map[string]NominatimResults),
}

func NominatimFetcher(email, server, query string) (res NominatimResults, err error) {
	return NominatimFetcherLong(email, server, "", "", query)
}

// NominatimFetcher makes Nominatim API request
func NominatimFetcherLong(email, server, country, city, query string) (res NominatimResults, err error) {
	maps := url.Values{}
	if country != "" {
		maps.Add("country", country)
	}
	if city != "" {
		maps.Add("city", city)
	}
	if query != "" {
		maps.Add("q", query)
	}
	maps.Add("format", "json")
	maps.Add("accept-language", "en")
	maps.Add("limit", "1")
	maps.Add("email", email)
	url := server + "/search?" + maps.Encode()
	nominatim.RLock()
	if v, ok := nominatim.cache[url]; ok {
		nominatim.RUnlock()
		return v, nil
	}
	nominatim.RUnlock()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := nominatim.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	nominatim.Lock()
	nominatim.cache[url] = res
	nominatim.Unlock()
	return
}
