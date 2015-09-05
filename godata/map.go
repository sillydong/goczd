package godata

//map转slice
func Map2Slice(params map[string]string, sep string) []string {
	sparam := make([]string, len(params))
	for key, val := range params {
		sparam = append(sparam, key+sep+val)
	}
	return sparam
}
