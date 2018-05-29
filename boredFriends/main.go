/*
Two friends Davit and Tigran were bored. So they decided to play a game.

Rules of the game are as follows: First, they choose an integer n. Then they take successive turns, starting with Davit. In each turn, a player chooses an integer from 1 to 8 (inclusive) and subtracts it from n.

The player wins if after his turn, n becomes 0.

Your task is to determine whether Davit will win or not, assuming that both players play optimally.

Examples

For n = "5", the output should be boredFriends(n) = true. Davit can make n = 0 after his first move.

For n = "9", the output should be boredFriends(n) = false. Davit can't win during his first move, but no matter what number he chooses, Tigran can win on the next move.
*/

func boredFriends(n string) bool {
	if len(n) < 2 {
		if n == "9" {
			return false
		}
		return true
	}
	r := 0
	for _, v := range n {
		a, _ := strconv.Atoi(string(v))
		r += a
	}
	return boredFriends(strconv.Itoa(r))
}
