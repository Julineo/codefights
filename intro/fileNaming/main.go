/*
https://app.codesignal.com/arcade/intro/level-12/sqZ9qDTFHXBNrQeLC

You are given an array of desired filenames in the order of their creation. Since two files cannot have equal names, the one which comes later will have an addition to its name in a form of (k), where k is the smallest positive integer such that the obtained name is not used yet.

Return an array of names that will be given to the files.

Example

For names = ["doc", "doc", "image", "doc(1)", "doc"], the output should be
fileNaming(names) = ["doc", "doc(1)", "image", "doc(1)(1)", "doc(2)"].

Input/Output

[execution time limit] 4 seconds (go)

[input] array.string names

Guaranteed constraints:
5 ≤ names.length ≤ 15,
1 ≤ names[i].length ≤ 15.
*/
import . "strconv"

func fileNaming(names []string) []string {
    mList := make(map[string]int)
    res := []string{}
    for _, v := range names {
        if _, ok := mList[v]; !ok {
            mList[v]++
            res = append(res, v)
        } else {
            for i := 1; i < 16; i++ {
                if _, ok := mList[v + "(" + Itoa(i) + ")"]; ok { continue }
                mList[v + "(" + Itoa(i) + ")"]++
                res = append(res, v + "(" + Itoa(i) + ")")
                break
            }
        }
    }
    return res
}
