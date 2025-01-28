package main

import (
	_ "embed"
	"fmt"
	"math"
	"time"
)

func main() {
	secs := int(time.Duration(math.MaxInt64) / time.Hour)
	fmt.Println(secs)
}
