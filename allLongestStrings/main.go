func allLongestStrings(inputArray []string) []string {
    var max int
    for _, a := range(inputArray) {
        if len(a) > max {
            max = len(a)
        }
    }
    var outputArray []string
    for _, a := range(inputArray) {
        if len(a) == max {
            outputArray = append(outputArray, a)
        }    
    }
    return outputArray
}
