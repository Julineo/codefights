/*
https://app.codesignal.com/challenge/rjNJ2CEF6wiXzXawb/solutions/wPtL9NtaYttNHcEnF

Alex lives in a small town that only has one store that sells school supplies. On the first day of school, she made a list of all the supplies she would need for each class. For every new item on the list, she made a note of the price, on a separate list.

Given the arrays supplies and prices, your task is to calculate how much Alex's school supplies will cost for this term.

NOTES:

The same item may appear more than once in the supplies list, but its price will only be logged in prices the first time it appears (see example 2).
The names of the items may or may not be pluralized, depending on how many are required.
To avoid floating point values, all prices will be listed in terms of cents (eg: 325 means $3.25).
Example

For supplies = ["4 pencils", "2 notebooks", "1 textbook"] and prices = [50, 100, 1500], the output should be schoolSupplies(supplies, prices) = 1900.

From the inputs, we can infer the following:

Item		Price	Amount required
Pencil		$0.50	4
Notebook	$1.00	2
Textbook	$15.00	1

So the total is 4 * 50 + 2 * 100 + 1 * 1500 = 1900.

For supplies = ["2 binders", "5 pens", "3 notebooks", "1 calculator", "4 notebooks", "1 binder", "3 folders"] and prices = [750, 150, 200, 1875, 250], the output should be schoolSupplies(supplies, prices) = 7025.

In this case, some items appear more than once in supplies (but only once in prices). Also note that "binders" and "binder" are considered the same item.

Item		Price	Amount required
Binder		$7.50	2 + 1 = 3
Pen		$1.50	5
Notebook	$2.00	3 + 4 = 7
Calculator	$18.75	1
Folder		$2.50	3

So the total is 2 * 750 + 5 * 150 + 3 * 200 + 1 * 1875 + 4 * 200 + 1 * 750 + 3 * 250 = 7025.
*/
import (
    "strings"
)

func schoolSupplies(supplies []string, prices []int) (res int) {
    supply := []string{}
    for _, s := range supplies {
        supply = append(supply, strings.Replace(s, "s", "", -1))
    }

    supplyC := []string{}
    m := make(map[string]int)
    for _, s := range supply {
        nItem := strings.Fields(s)
        ns, Item := nItem[0], nItem[1]
        _, ok := m[Item]
        if ! ok {
            supplyC = append(supplyC, Item)
        }
        n, _ := strconv.Atoi(ns)
        m[Item] += n
    }

    for i, s := range supplyC {
        res += m[s] * prices[i]
    }
    return
}
