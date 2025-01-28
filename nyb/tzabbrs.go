package nyb

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"math"
	"strings"
	"time"
)

// source: https://timezonedb.com/download
//
//go:embed time_zone.csv
var tzCSV []byte

var tzAbbrs = parseZoneInfo(tzCSV, target)

func parseZoneInfo(data []byte, target time.Time) (abbrs map[string]int) {
	type zoneInfo struct {
		abbr   string
		offset int
	}
	type zoneMap map[string]zoneInfo
	zones := make(zoneMap)

	entries := bytes.Split(data, []byte{'\n'})
	var zoneName string
	var country2iso string
	var abbr string
	var unixTime int64
	var offsetSec int
	var dst bool
	for _, entry := range entries {
		entry = bytes.TrimSpace(entry)
		if len(entry) == 0 {
			break
		}
		entry = bytes.ReplaceAll(entry, []byte{','}, []byte{' '})

		_, err := fmt.Sscanf(
			string(entry),
			"%s %s %s %d %d %t",
			&zoneName,
			&country2iso,
			&abbr,
			&unixTime,
			&offsetSec,
			&dst,
		)
		if err != nil {
			log.Fatal("zone abbr parse", err)
		}

		abbr = strings.ToUpper(abbr)
		if unixTime <= target.Unix() {
			zones[zoneName] = zoneInfo{
				abbr:   abbr,
				offset: offsetSec,
			}
		}
	}

	counter := make(map[string]map[int]int)
	for _, zone := range zones {
		if counter[zone.abbr] == nil {
			counter[zone.abbr] = make(map[int]int)
		}
		counter[zone.abbr][zone.offset]++
	}

	abbrs = make(map[string]int)
	var max int
	for abbr, offsets := range counter {
		max = 0
		for offset, count := range offsets {
			if count > max {
				max = count
				abbrs[abbr] = offset
			}
		}
	}
	abbrs["UTC"] = 0
	abbrs["Z"] = 0
	milneg := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "K", "L", "M"}
	milpos := []string{"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y"}
	for i := 1; i <= 12; i++ {
		abbrs[milneg[i-1]] = -i * 3600
		abbrs[milpos[i-1]] = i * 3600
	}
	tzCSV = nil
	return abbrs
}

func parseUTC(in string) (int, error) {
	var offset int64
	formats := []string{
		"UTC+%d:%d",
		"UTC-%d:%d",
		"GMT+%d:%d",
		"GMT-%d:%d",
	}
	for i := range formats {
		var hours int64
		var minutes int64
		_, err := fmt.Sscanf(in, formats[i], &hours, &minutes)
		if err == nil {
			if i%2 == 0 {
				offset += hours * 3600
				offset += minutes * 60
				if offset < 0 {
					return 0, fmt.Errorf("overflow")
				}
			} else {
				offset -= hours * 3600
				offset -= minutes * 60
				if offset > 0 {
					return 0, fmt.Errorf("underflow")
				}
			}
			if i > 1 {
				offset = -offset
			}
			if offset > int64(time.Duration(math.MaxInt64)/time.Second) || offset < int64(time.Duration(math.MinInt64)/time.Second) {
				return 0, fmt.Errorf("too big")
			}
			return int(offset), nil
		}
	}
	formatsShort := []string{
		"UTC+%d",
		"UTC-%d",
		"GMT+%d",
		"GMT-%d",
	}
	for i := range formatsShort {
		var hours int64
		_, err := fmt.Sscanf(in, formatsShort[i], &hours)
		if err == nil {
			if i%2 == 0 {
				offset += hours * 3600
				if offset < 0 {
					return 0, fmt.Errorf("overflow")
				}
			} else {
				offset -= hours * 3600
				if offset > 0 {
					return 0, fmt.Errorf("underflow")
				}
			}
			if i > 1 {
				offset = -offset
			}
			if offset > int64(time.Duration(math.MaxInt64)/time.Second) || offset < int64(time.Duration(math.MinInt64)/time.Second) {
				return 0, fmt.Errorf("too big")
			}
			return int(offset), nil
		}
	}
	return 0, fmt.Errorf("zone not found")
}
