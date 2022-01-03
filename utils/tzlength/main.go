//This utility prints lengths of zones string representation. Useful for seeing if some zone exceeds irc limit
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/rhinosf1/newyearsbot/nyb"
)

func main() {
	var zones nyb.TZS
	err := json.Unmarshal(nyb.Zones, &zones)
	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(sort.Reverse(zones))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight|tabwriter.Debug)

	for _, k := range zones {
		tmp := len([]byte(k.String()))
		if tmp > 0 {
			fmt.Fprintf(w, "%v\t%d\t%d\t\n", k.Offset, tmp, tmp-396)
		}
	}
	w.Flush()
}
