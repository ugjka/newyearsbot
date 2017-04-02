//This utility prints lenghts of zones string representation. Useful for seeing if some zone exceeds irc limit
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	c "github.com/ugjka/newyearsbot/common"
)

func main() {
	tzdatapath := flag.String("tzpath", "../tz.json", "path to tz.json")
	//Check if tz.json exists
	if _, err := os.Stat(*tzdatapath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: file %s does not exist\n", *tzdatapath)
		os.Exit(1)
	}
	var zones c.TZS
	file, err := os.Open(*tzdatapath)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &zones)
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
