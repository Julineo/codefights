/*
https://app.codesignal.com/challenge/TgCWyhmmB9rBnf2uY/solutions/NBPvAibS5qHisCGYb

You have an exam coming up tomorrow morning, and you're trying to decide how much longer you should keep studying before going to bed. Your uncle, reflecting on his experience in med school, once told you that if you have the choice between sleep and study, it's always better to get the sleep. But since this is your first semester, you suspect you might not be as well prepared for your exams as he would have been.

Given the integers familiarity and hoursRemaining, your task is to return how many more hours you should study for (assuming you're trying to maximize your expected grade). familiarity represents how confident you are in the material; basically the grade you would get if you were to take the exam right now, while hoursRemaining represents the number of hours between now and your exam.

Notes:

For every hour you spend studying, you diminish your 'unfamiliarity' by 20% (eg: if your familiarity is 70, then your unfamiliarity is 30, so after 1 hour of studying, it'll be down to 24, which means your familiarity will be 76).
You haven't slept yet, so to be fully alert at the exam, you're depending on getting 8 hours of sleep beforehand. If you get fewer than 8 hours, your performance will suffer - it's estimated that for every hour less than 8, your grade will take a 5 point penalty (eg: with a familiarity of 87, if you only sleep 5 hours, then since you're missing 3 hours, your expected grade will be 87 - 3 * 5 = 72).
Each hour will be uniquely committed to one activity (studying or sleeping); you have a regimented schedule and you'd prefer not to split an hour into half study and half sleep.
If your familiarity is already at 100, you probably shouldn't spend any more time studying.
Example

For familiarity = 64 and hoursRemaining = 8, the output should be studyOrSleep(familiarity, hoursRemaining) = 2.

After 1 hour, your familiarity rises to 71.2, but considering the 5 point penalty for losing the hours of sleep, it's effectively 66.2. That's more than you started with, so maybe it's worth it to stay up another hour.

After 2 hours, your familiarity would be 76.96, but now the penalty would be 10 points, so it's basically 66.96. That's a little bit higher still, so it could even be worth it to stay up 2 more hours.

After 3 hours, your familiarity would be at about 81.57, but with 15 points deducted, the resulting 66.57 is less than the previous result, so it's not worth it to lose any more sleep.

For familiarity = 38 and hoursRemaining = 9, the output should be studyOrSleep(familiarity, hoursRemaining) = 5.

In this case, you have an hour to study before it'll cost you any sleep time, so we should expect the result to be at least 1. How many more hours is it worth it to stay up?

Hours of study	Penalty for lost sleep	Familiarity	Expected result
0		0			38		38
1		0			50.4		50.4
2		5			60.32		55.32
3		10			68.26		58.26
4		15			74.6		59.6
5		20			79.68		59.68
6		25			83.75		58.74
7		30			86.99		56.99
The best result is expected with 5 hours of studying.

For familiarity = 75 and hoursRemaining = 8, the output should be studyOrSleep(familiarity, hoursRemaining) = 0.

By studying another hour, your familiarity would become 80, but since there would be a 5 point penalty on losing the hour of sleep, it would remain at 75, so it's not worth it. Get the 8 hours of sleep.
*/

func studyOrSleep(familiarity int, hoursRemaining int) int {
    familiarityF := float64(familiarity)
    hoursNightStudy := 0
    hoursStudy := 0
    max := familiarityF
    maxH := 0
    for hoursRemaining > 0 {
        pnlty := 0
        hoursRemaining--
        hoursStudy++
        familiarityF = fChange(familiarityF)
        if hoursRemaining < 8 {
            hoursNightStudy++
            pnlty = hoursNightStudy * 5
        }
        if familiarityF - float64(pnlty) > max {
            max = familiarityF - float64(pnlty)
            maxH = hoursStudy
        }
    }
    return maxH
}

func fChange(f float64) float64 {
    uf := 100 - f
    perc := uf * 0.2
    impr := uf - perc
    return 100 - impr
}
