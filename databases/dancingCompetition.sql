/*
https://app.codesignal.com/arcade/db/selecting-what-to-select/YPLqWmfgWJDKTjPAw

During the dance contest, each judge evaluates the dancers' performances based on three criteria, using a score from 1 to 10 for each of the criteria. You are given a table of the scores awarded to a specific dancer by each judge. Recently, the dance contest made the decision to drop scores awarded by a certain judge if that judge gave an extreme (either minimum or maximum) score for at least two criteria.

For example, if the judge awarded one of the minimum scores for musicality (i.e., there may be similar scores for musicality, but there may not be smaller scores for that criterion) and one of the maximum scores for floorcraft, all three of the scores given by that judge should not be taken into account.

Return a table that consists of only the scores that should be considered after this filtering, sorted by arbiter_id.

The scores table contain the following columns:

arbiter_id - the unique ID of the judge;
first_criterion - the score given for the first criterion;
second_criterion - the score given for the second criterion;
third_criterion - the score given for the third criterion.
Example

For given table scores

arbiter_id	first_criterion	second_criterion	third_criterion
1		3		10			10
2		2		3			4
3		5		6			7
4		2		5			9
5		2		2			2

the output should be

arbiter_id	first_criterion	second_criterion	third_criterion
2		2		3			4
3		5		6			7
4		2		5			9

The first judge gave the maximal scores for the second and third criteria, so his scores aren't included in the answer. The fifth judge's given scores are all minimal, so her scores aren't included to the answer either.
*/

CREATE PROCEDURE dancingCompetition()
BEGIN
    SET @min_f = (SELECT MIN(first_criterion) FROM scores);
    SET @max_f = (SELECT MAX(first_criterion) FROM scores);
    SET @min_s = (SELECT MIN(second_criterion) FROM scores);
    SET @max_s = (SELECT MAX(second_criterion) FROM scores);
    SET @min_t = (SELECT MIN(third_criterion) FROM scores);
    SET @max_t = (SELECT MAX(third_criterion) FROM scores);

    SELECT s.* FROM scores s
    JOIN
    (
    SELECT arbiter_id, CASE WHEN (first_criterion <= @min_f OR first_criterion >= @max_f) THEN 1 ELSE 0 END
        + CASE WHEN (second_criterion <= @min_s OR second_criterion >= @max_s) THEN 1 ELSE 0 END 
        + CASE WHEN (third_criterion <= @min_t OR third_criterion >= @max_t) THEN 1 ELSE 0 END AS max0
    FROM scores
    ) b ON s.arbiter_id = b.arbiter_id AND max0 < 2;
END
