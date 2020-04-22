package slack

import (
	"encoding/json"
	"log"
)

func toStringSlice(s []interface{}) []string {
	r := make([]string, len(s))
	for i, c := range s {
		r[i] = c.(string)
	}
	return r
}

func pp(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		log.Printf(string(b))
	}
	return
}
