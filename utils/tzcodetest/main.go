package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "embed"
)

func main() {
	target := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)
	abbrs := parseTimezoneAbbrs(tzcsv, target)
	fmt.Println(abbrs)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		v, ok := abbrs[sc.Text()]
		if ok {
			fmt.Println(time.Now().In(time.FixedZone(sc.Text(), v)))
		}
	}
}

// source: https://timezonedb.com/download
//
//go:embed time_zone.csv
var tzcsv []byte

func parseTimezoneAbbrs(data []byte, target time.Time) (abbrs map[string]int) {
	type zoneinfo struct {
		abbr   string
		offset int
	}
	type zonemap map[string]zoneinfo
	zones := make(zonemap)

	entries := bytes.Split(data, []byte{'\n'})
	var zonename string
	var country2iso string
	var abbr string
	var unixtime int64
	var offsetsec int
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
			&zonename,
			&country2iso,
			&abbr,
			&unixtime,
			&offsetsec,
			&dst,
		)
		if err != nil {
			log.Fatal("zone abbr parse", err)
		}

		abbr = strings.ToUpper(abbr)
		if unixtime <= target.Unix() {
			zones[zonename] = zoneinfo{
				abbr:   abbr,
				offset: offsetsec,
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
	tzcsv = nil
	return abbrs
}
