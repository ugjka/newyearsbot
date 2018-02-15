package nyb

import (
	"encoding/json"
	"testing"
	"time"
)

func TestZones(t *testing.T) {
	var zones TZS
	err := json.Unmarshal([]byte(Zones), &zones)
	if err != nil {
		t.Errorf("Corrupted zone data: %s", err)
	}
	for _, zone := range zones {
		_, err := time.ParseDuration(zone.Offset + "h")
		if err != nil {
			t.Errorf("Could not parse offset %s: Err: %s", zone.Offset, err)
		}
		for _, country := range zone.Countries {
			if country.Name == "" {
				t.Errorf("Empty country name in zone %s", zone.Offset)
			}
			for _, city := range country.Cities {
				if city == "" {
					t.Errorf("Empty city name in %s in zone %s", country.Name, zone.Offset)
				}
			}
		}
	}
}
