/*
https://app.codesignal.com/challenge/QTiyM5zk2mHmC8aox

Input:
time: "1:23PM"

Expected Output:
"12:23PM"

*/
import . "time"

func fallBack(t string) string {
	r, _ := Parse("3:04PM", t)
    return fmt.Sprint(r.Add(-Hour).Format(Kitchen))
}
