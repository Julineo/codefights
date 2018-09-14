/*
https://app.codesignal.com/arcade/db/a-table-of-desserts/Yd8P4ebwe2RQJFxeo

You found a really old computer at your local library, and it looks like there's still some data about old library books left on its hard drive.

After closer examination, you've discovered that there is a book_chapters table, which has the following structure:

chapter_name: the name of the book's chapter;
chapter_number: the unique number of the chapter written as a Roman numeral.
Note that there could be gaps in chapter numbers. Also, every library book has fewer than 1000 chapters.

Here is the table of Roman numeral symbols and their values:

Symbol	Value
I	1
V	5
X	10
L	50
C	100
D	500
M	1000
After looking through the library, you've found what seems to be the remnants of this same book! Unfortunately its chapters are not numbered on the pages.

You decide to find the order in which you should read the books chapters by using the old computer and the information stored in the book_chapters table.

Given the book_chapters table, compose a results table that consists of a single column chapter_name that contains the names of the book's chapters. The table should be sorted in ascending order by the chapter's actual numbers (i.e. chapter number III should come before chapter number V).

Example

For the following table book_chapters

chapter_name	chapter_number
A Dead Man		LVI
Behaviour in General	I
Cast Up			XLIX
Imitation		IX
Nemesis			L
Paste			XXIII
The Cub			XXI
The Oxenham Arms	XXIV
Two Bequests		XLVII

the output should be

chapter_name
Behaviour in General
Imitation
The Cub
Paste
The Oxenham Arms
Two Bequests
Cast Up
Nemesis
A Dead Man

Here are the chapter number converted to Hindu-Arabic numerals:

chapter_name	chapter_number
A Dead Man		56
Behaviour in General	1
Cast Up			49
Imitation		9
Nemesis			50
Paste			23
The Cub			21
The Oxenham Arms	24
Two Bequests		47
*/
CREATE FUNCTION fromRoman (inRoman varchar(15)) RETURNS integer DETERMINISTIC
BEGIN

    DECLARE numeral CHAR(7) DEFAULT 'IVXLCDM';

    DECLARE digit TINYINT;
    DECLARE previous INT DEFAULT 0;
    DECLARE current INT;
    DECLARE sum INT DEFAULT 0;

    SET inRoman = UPPER(inRoman);

    WHILE LENGTH(inRoman) > 0 DO
        SET digit := LOCATE(RIGHT(inRoman, 1), numeral) - 1;
        SET current := POW(10, FLOOR(digit / 2)) * POW(5, MOD(digit, 2));
        SET sum := sum + POW(-1, current < previous) * current;
        SET previous := current;
        SET inRoman = LEFT(inRoman, LENGTH(inRoman) - 1);
    END WHILE;

    RETURN sum;
END;

CREATE PROCEDURE sortBookChapters()
BEGIN
	SELECT chapter_name
    FROM book_chapters
    ORDER BY fromRoman(chapter_number);
END
