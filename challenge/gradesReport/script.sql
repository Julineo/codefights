/*
https://app.codesignal.com/challenge/ECxw2gmfYWtdjLL33

You are helping your teacher at the university to generate a report containing three columns: name, grade and mark. You are given a table of grades and students. You should create a report, satisfying your teacher's conditions. He doesn't want the names of those students who received a grade lower than C. Higher grades in report should be entered first. If there is more than one student with the same grade (C-A) assigned to them, order those particular students by their name alphabetically. Finally, if the grade is lower than C, use "NULL" as their name and list them by their grades in descending order. If there is more than one student with the same grade (F-D) assigned to them, order those particular students by their marks in ascending order.

The grades table contains the following columns:

grade - letter from A to F;
min_mark - minimal mark, starting from which student got this particular grade;
max_mark - maximal mark, having which student still has this particular grade.
It's guaranteed that table of grades covers all possible marks (from 0 to 100).

The students table contains the following columns:

name - first and last names of student, all names are unique;
mark - mark he got, integer from 0 to 100.
Your task is to return table with columns name, grade and mark, satisfying the conditions.

Example

For given tables grades

grade	min_mark	max_mark
A	91	100
B	82	90
C	70	81
D	64	69
E	60	63
F	0	59

and students

name	mark
Drake Porter	78
Ewan Black	100
Jamie Miller	76
Joe Allen	43
Rhys Kelly	54
Rowen Little	67
Ryan Richards	90

the output should be

name	grade	mark
Ewan Black	A	100
Ryan Richards	B	90
Drake Porter	C	78
Jamie Miller	C	76
NULL		D	67
NULL		F	43
NULL		F	54
*/

CREATE PROCEDURE gradesReport()
BEGIN
    SELECT ELT(g.grade < 'D', s.name) name, g.grade, s.mark
    FROM students s
    JOIN grades g ON s.mark BETWEEN g.min_mark AND g.max_mark
    ORDER BY 2, 1, 3;
END
