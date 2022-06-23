package iteration

func Repeat(s string, n int) string {
	repeater := ""
	for i := 0; i < n; i++ {
		repeater += s
	}
	return repeater
}
