package jsonx

import (
	"encoding/json"
)

//ToJSON struct to json
func ToJSON(v ...interface{}) string {
	vbyte, _ := json.MarshalIndent(v, " ", "    ")
	return string(vbyte)
}
