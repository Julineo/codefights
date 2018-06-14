import (
    "sort"
    "math"
)

func maxSalePrice(mass int, marketValues []int) (maxV int) {
    type tps struct {
        price float64
        mv int
        ms int
    }
    ps := []tps{}
    for s, v := range marketValues {
        p := float64(v)/float64(s)
        if p == 0 || math.IsNaN(p) {
            continue
        }
        ps = append(ps, tps{p, v, s})
    }
    
    /*sort.Slice(ps, func(i, j int) bool {
		return ps[i].mv >= ps[j].mv && ps[i].ms >= ps[j].ms
	})
*/
    
    ids := []int{}
    
    for i := 0; i < len(ps); i++ {
        ids = append(ids, i)
    }
    
    fmt.Println(ps)
     
    idss := [][]int{{}}
    idss = permutations(ids)
    for _, ii := range idss {
        tmpMaxV := 0
        lms := 0
        for _, i := range ii {
            if ps[i].ms + lms > mass {
                break
            }
            tmpMaxV += ps[i].mv
            lms += ps[i].ms
        }
        if tmpMaxV > maxV {
            maxV = tmpMaxV
        }
    }
   
    return  maxV
}

func permutations(arr []int)[][]int{
    var helper func([]int, int)
    res := [][]int{}

    helper = func(arr []int, n int){
        if n == 1{
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
}
