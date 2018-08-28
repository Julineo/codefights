/*
https://app.codesignal.com/arcade/db/would-you-like-the-second-meal/TfCxoCqHXShx7Hhzt

Example

For the following table expressions

id	a	b	operation	c
1	2	3	+		5
2	2	3	+		6
3	3	2	/		1
4	4	7	*		28
5	54	2	-		27
6	3	0	/		0


the output should be

id	a	b	operation	c
1	2	3	+		5
4	4	7	*		28
Explanation:

2 + 3 = 5 - correct;
2 + 3 = 6 - incorrect;
3 / 2 = 1 - incorrect;
4 * 7 = 28 - correct;
54 - 2 = 27 - incorrect;
3 / 0 = 0 - incorrect.
*/
CREATE PROCEDURE expressionsVerification()
BEGIN
	SELECT * FROM expressions
    WHERE CASE operation 
        WHEN '+' THEN a + b = c
        WHEN '-' THEN a - b = c
        WHEN '*' THEN a * b = c
        WHEN '/' THEN a / b = c
        END
;
END
