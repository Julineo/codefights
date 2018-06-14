/*
All tunnels meet underground at a common point, and for each tunnel, we know the coordinates of the entry point and the distance to this junction. For example, we could represent the following map by tunnels = [[1, 2, 9], [2, 7, 5], [8, 3, 4]]:

In this case, the underground distance from (1, 2) to (2, 7) would be 9 + 5 = 14. The underground distance from (2, 7) to (8, 3) would be 5 + 4 = 9.

Given the startPoint (coordinates of where you live), the endPoint (coordinates of your destination), and tunnels (a list of coordinates and distances for each tunnel), find the shortest total distance you'll need to travel. All above-ground travel is measured according to Manhattan distance.

Note: you can pass through the coordinates of a tunnel entrance without entering the tunnel, and it's also possible that there's a tunnel entrance at your start or end point.
*/
func tunnelPath(sp []int, ep []int, t [][]int) (min int) {
    min = dist(sp[0], sp[1], ep[0], ep[1])
    path := 0
    for k := 0; k < len(t); k++ {
        for l := 0; l < len(t); l++ {
            if l == k {
                continue
            }
            //distance to tunnel
            path += dist(sp[0], sp[1], t[k][0], t[k][1])

            //distance in tunnel
            path += t[k][2]+t[l][2]

            //distance from tunnel
            path += dist(t[l][0], t[l][1], ep[0],ep[1])

            if path < min {
                min = path
            }
            path = 0
        }
    }
    return
}

func dist(x1 , y1 , x2, y2 int) int {
    return abs(x2 - x1) + abs(y2 - y1)
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}
