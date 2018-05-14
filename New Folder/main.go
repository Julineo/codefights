import "strings"

//stack struct
type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	// what do we do if the stack is empty?
	//fmt.Println(s)

	l := len(s)
	return s[:l-1], s[l-1]
}

//function to reverse inside the string from i to j
func reverse(s string, i int, j int) string {
    runes := []rune(s)
    for ; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func reverseParentheses(s string) string {
	//stak for "(" positions
	ss := make(stack, 0)

	//start index
	var is int

	//go throug
	for i, v := range s {
		if v == '(' {
			//fmt.Println("(", i)
			ss = ss.Push(i)
		}
		if v == ')' {
			//fmt.Println(")", i)
			ss, is = ss.Pop()
			//fmt.Println("is: ", is)
			//fmt.Println("reverse from ", is, " till ", i)
			s = reverse(s, is, i)
			//fmt.Println(s)
		}
	}

	return strings.Replace(strings.Replace(s, "(", "", -1), ")", "", -1)
}
