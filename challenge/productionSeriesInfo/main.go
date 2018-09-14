/*
https://app.codesignal.com/challenge/XanxRWkdG5vKryKxo/solutions/CBANSF4nDp54kZDBF

To make it easier to calculate the remaining stock of ingredients in your factory, you decided to create a small algorithm based on the following information:

Knowing the current stock (ingredients), the amount of ingredients used in each recipe you produce in the factory (recipe1 and recipe2), and the production order (productionSeries), your task is to determine whether you'll have sufficient ingredients to start the machines for the next production series.

Important note: If productionSeries holds a different number than 1 or 2, it means it's a special recipe that uses the exact amount of ingredients the number shows. Otherwise, it's referring to recipe1 or recipe2.

The expected output is a 1 or 2-element array with the following format:

ingredients	output[0]	output[1]
enough for 1 series	"Ok"	(no output[1])
enough for more than 1	"Ok"	"Ingredients for x more series"
not enough	"Out of ingredients!"	"Missing x ingredients"
Example

For ingredients = 16, recipe1 = 2, recipe2 = 5, and productionSeries = [1, 1, 2, 1, 2], the output should be productionSeriesInfo(ingredients, recipe1, recipe2, productionSeries) = ["Ok"].

You have ingredients for just one production series:
2 + 2 + 5 + 2 + 5 = 16

For ingredients = 25, recipe1 = 1, recipe2 = 2, and productionSeries = [1, 1, 3, 2], the output should be productionSeriesInfo(ingredients, recipe1, recipe2, productionSeries) = ["Ok", "Ingredients for 2 more series"].

You have ingredients for the first production and the remaining amount will allow you to produce two more.
1 + 1 + 3 + 2 = 7 * 3 = 21

For ingredients = 10, recipe1 = 1, recipe2 = 10, and productionSeries = [1, 2], the output should be productionSeriesInfo(ingredients, recipe1, recipe2, productionSeries) = ["Out of ingredients!", "Missing 1 ingredients"].

You're just 1 short on ingredients.
1 + 10 = 11

Input / Output

[execution time limit] 4 seconds (go)

[input] integer ingredients

The quantity of ingredients remaining in the factory.

Guaranteed constraints:
0 ≤ ingredients ≤ 107

[input] integer recipe1

The amount of ingredients used to produce recipe number 1.

Guaranteed constraints:
1 ≤ recipe1 ≤ 106

[input] integer recipe2

The amount of ingredients used to produce recipe number 2.

Guaranteed constraints:
1 ≤ recipe2 ≤ 106

[input] array.integer productionSeries

The production order of a single serie of production. If productionSeries[i] is 1, you produce recipe number 1; if productionSeries[i] is 2, you produce recipe 2. For other numbers, you use productionSeries[i] as the ingredient amount.

Guaranteed constraints:
1 ≤ productionSeries.length ≤ 50
1 ≤ productionSeries[i] ≤ 106

[output] array.string

An array of strings representing your report on the production. See the note above for formatting specifics.
*/
func productionSeriesInfo(ingredients int, recipe1 int, recipe2 int, productionSeries []int) []string {
    ingredientsOrig := ingredients
    for _, v := range productionSeries {
        switch v {
            case 1:
                ingredients = ingredients - recipe1
            case 2:
                ingredients = ingredients - recipe2
            default:
                ingredients = ingredients - v
        }
    }

    cycleSp := ingredientsOrig - ingredients

    if ingredients < 0 {
        return []string{"Out of ingredients!", "Missing " + strconv.Itoa(-ingredients) + " ingredients"}
    }

    if ingredients/cycleSp == 0 {
        return []string{"Ok"}
    }

    return []string{"Ok", "Ingredients for " + strconv.Itoa(ingredients/cycleSp) +" more series"}
}
