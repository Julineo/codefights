/*
1
["-"]

2
[" - ", 
 "-|-", 
 " - "]

3
["  -  ", 
 " -|- ", 
 "-|-|-", 
 " -|- ", 
 "  -  "]

4
["   -   ", 
 "  -|-  ", 
 " -|-|- ", 
 "-|-|-|-", 
 " -|-|- ", 
 "  -|-  ", 
 "   -   "]
*/
import "strings"

func dashes(n int) []string {
	n--
	var r, e []string
		for i := 0; i <= n; i++ {
			s := ""
			s = s + strings.Repeat(" ", n-i)
			s = s + strings.Repeat("-|", i)
			s = s + "-"
			s = s + strings.Repeat(" ", n-i)
			r = append(r, s)
			e = append([]string{s}, e...)
		}
	r = append(r, e[1:]...)
	return r
}

