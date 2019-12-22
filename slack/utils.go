package slack

func toStringSlice(s []interface{}) []string {
	r := make([]string, len(s))
	for i, c := range s {
		r[i] = c.(string)
	}
	return r
}
