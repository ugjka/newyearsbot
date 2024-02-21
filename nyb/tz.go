package nyb

import _ "embed"

// Zones contains time zone information in JSON format
//
//go:embed tz.json
var Zones []byte
