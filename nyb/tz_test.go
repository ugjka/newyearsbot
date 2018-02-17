package nyb

import (
	"encoding/json"
	"testing"
)

func TestZones(t *testing.T) {
	var zones TZS
	err := json.Unmarshal([]byte(Zones), &zones)
	if err != nil {
		t.Errorf("Corrupted zone data: %s", err)
	}
	for _, zone := range zones {
		for _, country := range zone.Countries {
			if country.Name == "" {
				t.Errorf("Empty country name in zone %v", zone.Offset)
			}
			for _, city := range country.Cities {
				if city == "" {
					t.Errorf("Empty city name in %s in zone %v", country.Name, zone.Offset)
				}
			}
		}
	}
}
