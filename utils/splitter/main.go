// This utility tests TZ string splitting
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/ugjka/newyearsbot/nyb"
)

func main() {
	var zones nyb.TZS
	err := json.Unmarshal(nyb.Zones, &zones)
	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(sort.Reverse(zones))

	for _, k := range zones {
		fmt.Println("**********************")
		fmt.Println(k.Split(200))
	}
}
