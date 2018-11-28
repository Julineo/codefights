/*

For [[24,7], [31,15], [50,41]], the output should be 82.

You start at position 0, so it takes 24 seconds to reach the first stoplight (which is 24 metres away). This first stoplight has a frequency of 7, which means it's green at 0 seconds, it turns red at 7 seconds, back to green at 14 seconds, then red again st 21 seconds, then it won't turn green again until 28 seconds. Since you arrive at 24 seconds, you'll have to wait 4 seconds for the light to turn green again.

After the first light turns green, you travel another 7 metres to the seconds stoplight at position 31. By this time, 35 seconds have passed, so the light is already green, and you don't have to stop (this light has a frequency of 15, so it would have started green at 0, turned red at 15, then back to green at 30, until 45, so at 35 seconds it's green).

You then travel another 19 metres to the final stoplight, which you arrive at after a total of 54 seconds. Since this light has a frequency of 41, it would have turned red at 41 seconds, and will stay red until 82 seconds have passed. After 82 seconds, you're able to make it through the last light, so this is the final answer.

For [[23,23], [37,30]], the output should be 60.

You arrive at the first stoplight at the exact moment it turns red 😡. This means you'll have to wait until another 23 seconds pass, so the time will be 46 by the time you can go.

After travelling another 14 metres, the time is now 60, which means the second stoplight will have just turned green the instant you arrive at it (it was green until 30 seconds, then back to green right at 60). So you can pass through this final stoplight after 60 seconds.

*/
func trafficLights1D(r [][]int) int {
    t := 0
    d := 0
    for _, v := range r {
        t += v[0] - d
        d = v[0]
        if t/v[1] % 2 > 0 {
            t += v[1] - t%v[1]
        }
    }
    return t
}
