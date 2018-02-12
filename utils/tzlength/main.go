//This utility prints lenghts of zones string representation. Useful for seeing if some zone exceeds irc limit
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	c "github.com/ugjka/newyearsbot/common"
	nyb "github.com/ugjka/newyearsbot/nyb"
)

func main() {
	var zones c.TZS
	err := json.Unmarshal([]byte(nyb.TZ), &zones)
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(sort.Reverse(zones))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for _, k := range zones {
		tmp := len([]byte(k.String()))
		//if tmp > 396 {
		if tmp > 0 {
			fmt.Fprintf(w, "%s\t%d\t%d\t\n", k.Offset, tmp, tmp-396)
		}
	}
	w.Flush()
}
