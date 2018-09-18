func rome(n int) string {
    return Roman(n)
}

func Roman(arg int)(string) {

    figure := []int{1000,100,10,1}
    roman_digitA := []string{
    1 : "I",
	10 : "X",
	100 : "C",
	1000 : "M",
    }
    roman_digitB := []string{
	1 : "V",
	10 : "L",
	100 : "D",
	1000 : "MMMMM",
    }

    if arg < 0 || arg > 4000 {
	return "ROMAN_OUT_OF_RANGE"
    }

    roman_slice := []string{}
    x := ""

    for _, f := range figure {
        digit, i, v := int(arg / f), roman_digitA[f], roman_digitB[f]
        switch digit {
        case 1: roman_slice = append(roman_slice, string(i) )
        case 2: roman_slice = append(roman_slice, string(i) + string(i))
        case 3: roman_slice = append(roman_slice, string(i) + string(i) + string(i) )
        case 4: roman_slice = append(roman_slice, string(i) + string(v) )
        case 5: roman_slice = append(roman_slice, string(v) )
        case 6: roman_slice = append(roman_slice, string(v) + string(i) )
        case 7: roman_slice = append(roman_slice, string(v) + string(i)+ string(i) )
        case 8: roman_slice = append(roman_slice, string(v) + string(i)+ string(i)+ string(i) )
        case 9: roman_slice = append(roman_slice, string(i) + string(x)  )
	}

        arg -= digit * f
        x = i
    }
    ret := ""
    for _, e := range roman_slice {
        ret += e
    }
    return ret
}
