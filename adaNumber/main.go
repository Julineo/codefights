import (
    "strings"
)

func adaNumber(line string) bool {
    
    if len(line) == 0 {
		return false
	}
    
    if strings.HasPrefix(line, "_") {
        return false
    }
    
    if strings.HasSuffix(line, "_") {
        return false
    }
    
    if strings.HasPrefix(line, "#") {
        return false
    }
    line = strings.Replace(line, "_", "", -1)

    if strings.Count(line, "#") == 1 || strings.Count(line, "#") > 2 {
        return false
    }
    
    
    
    fmt.Println(strings.Split(line, "#"))
    n := strings.Split(line, "#")
    if len(n) == 3 {
        if n[2] != "" || n[1] == "" {
            return false
        }
        base, err := strconv.Atoi(n[0])
        if err != nil {
            return false
        }
        if base < 2 || base > 16 {
            return false
        }
        fmt.Println(base, n[1])
        for _, v := range(n[1]) {
            if _, err := strconv.ParseUint(string(v), base, 64); err != nil {
            fmt.Println(err)
            return false
        }
        }

    }
    
    if ! strings.ContainsAny(line, "#") {
        for _, v := range(line) {
            if _, err := strconv.ParseUint(string(v), 10, 64); err != nil {
                fmt.Println(err)
                return false
            }
        }
    }


    /*
    2 01
    3 012
    4 0123
    5 01234
    6 012345
    7 0123456
    8 01234567
    9 012345678
    10 0123456789
    11 0123456789a
    12 0123456789ab
    13 0123456789abc
    14 0123456789abcd
    15 0123456789abcde
    16 0123456789abcdef
    */
    return true
}
