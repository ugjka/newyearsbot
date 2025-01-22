package nyb

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
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
		abbrs[fmt.Sprintf("UTC+%d", i)] = i * 3600
		abbrs[fmt.Sprintf("UTC-%d", i)] = -i * 3600
		abbrs[fmt.Sprintf("GMT+%d", i)] = i * 3600
		abbrs[fmt.Sprintf("GMT-%d", i)] = -i * 3600
	}
	tzCSV = nil
	return abbrs
}
