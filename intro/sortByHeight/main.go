import "sort"

func sortByHeight(a []int) []int {
	//indexes with -1
    imo := make(map[int]bool)
    for i, v := range(a) {
        if v == -1 {
            imo[i] = true
        }
    }
    fmt.Println(imo)
    sort.Ints(a)
    for i := 0; i < len(a); i++ {
        if imo[i] && a[i] == -1 {
            continue
        }
        if a[i] != -1 {
            continue
        }
        for j := i + 1; j < len(a); j++ {
            if a[j] > -1 {
                fmt.Println(a[i], a[j])
                a[i], a[j] = a[j], a[i]
                break
            }
        }
    }
    
    return a
}
