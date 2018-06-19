/*
Check if all digits of the given integer are even.
*/

func evenDigitsOnly(n int) bool {
    for {
        d := n % 10
        n = n / 10
        if d % 2 != 0 {
            return false
        }
        if n == 0 {
            return true
        }
    }
}
