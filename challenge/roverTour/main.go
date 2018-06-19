/*
https://codefights.com/challenge/4oRAnJXGqnL9K7Rdf

It's a beautiful day, and you've decided to take your solar-powered rover out for a tour of the Martian landscape.

From the Hellas impact basin to the Elysium volcanic province, there's quite a variety of terrain on Mars, so each kilometre of your trip may cost a different amount of solar energy to travel through.

Given an array terrainDifficulty, representing the energy cost for each kilometre of the trip, you'd like to find the longest strip of land you can travel across with your current energy. Return the length of this trip (in kilometres).

Example

For terrainDifficulty = [3, 5, 4, 2, 3, 1, 5, 3, 4] and energy = 11, the output should be roverTour(terrainDifficulty, energy) = 4.
*/
func roverTour(td []int, energy int) (max int) {
    i, j := 0, 0
    tmp := 0
    for j < len(td) {
        tmp += td[j]
        if tmp > energy {
            tmp -= td[i]
            i++
        }
        if max < j - i + 1 {
            max = j - i + 1
        }
        j++
    }
    return
}
