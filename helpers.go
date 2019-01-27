package go4th

func in(s string, ls []string) bool {
	for _, b := range ls {
		if b == s {
			return true
		}
	}
	return false
}
