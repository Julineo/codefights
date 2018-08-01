/*
https://app.codesignal.com/challenge/aCvD2B4rqZvziDZuy
There's a popular sequence in mathematics that starts with a given number, and proceeds according to the following rules:

If the number is odd, multiply by 3 and add 1.
If the number is even, divide by 2.
A well-known conjecture suggests that when you keep doing this, you will always reach 1 eventually.

For example, if you start with 4, you reach 1 in 2 steps:
4 -> 2 -> 1

If you started with 3, you reach 1 in 7 steps:
3 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1

David, a talented CodeSignal user, has studied this sequence for a long time, and he's discovered an algorithm to find the total number of positive integers which require exactly K steps to reach 1 for the first time. Because of his curiosity, he'd like to extend his algorithm so that it can work for any positive integer (not just 1), but he's having some trouble.

Given an integer steps and an integer number, let's help David find the total number of positive integers which require exactly steps to reach number for the first time.

Examples

For steps = 1 and number = 1, the output should be IsaacRule(steps, number) = 1.

Only the number 2 will reach 1 after exactly 1 step:
2 -> 1
So the answer is 1.

For steps = 2 and number = 5, the output should be IsaacRule(steps, number) = 2 .

3 and 20 are the only two numbers that will reach 5 after 2 steps:
3 -> 10 -> 5 and 20 -> 10 -> 5
So the answer is 2.
*/
function IsaacRule(steps, number) {
    return collatzM(steps+1, number)
}

function collatzM(maxLevels, n){
    let levels = [[n]]
    while (levels.length < maxLevels) {
        let level = [];
        levels[levels.length-1].forEach(function(num){
            const mod6 = ((num-4)%6);
            const odd = (num-1)/3;
            if (!mod6 && odd != 1) {
                level.push(odd)
            }
            level.push(num*2);
        })
        levels.push(level);
    }
    return levels[maxLevels - 1].length;
}
