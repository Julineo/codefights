/*
https://app.codesignal.com/arcade/intro/level-9/xHvruDnQCx7mYom3T/description

Each day a plant is growing by upSpeed meters. Each night that plant's height decreases by downSpeed meters due to the lack of sun heat. Initially, plant is 0 meters tall. We plant the seed at the beginning of a day. We want to know when the height of the plant will reach a certain level.

Example

For upSpeed = 100, downSpeed = 10, and desiredHeight = 910, the output should be
growingPlant(upSpeed, downSpeed, desiredHeight) = 10.
*/
func growingPlant(upSpeed int, downSpeed int, desiredHeight int) (res int) {
    height := 0
    for {
        height += upSpeed
        res++
        if height >= desiredHeight {
            return
        }
        height -= downSpeed
    }
}
