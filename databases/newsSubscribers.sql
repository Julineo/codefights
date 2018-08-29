/*
Given the full_year and half_year tables, compose the result as follows: The resulting table should have one column subscriber that contains all the distinct names of anyone who is subscribed to a newspaper with the word Daily in its name. The table should be sorted in ascending order by the subscribers' first names.

Example

The following tables full_year

id	newspaper	subscriber
1	The Paragon Herald	Crissy Sepe
2	The Daily Reporter	Tonie Moreton
3	Morningtide Daily	Erwin Chitty
4	Daily Breakfast	Tonie Moreton
5	Independent Weekly	Lavelle Phu
and half_year

id	newspaper	subscriber
1	The Daily Reporter	Lavelle Phu
2	Daily Breakfast	Tonie Moreton
3	The Paragon Herald	Lia Cover
4	The Community Gazette	Lavelle Phu
5	Nova Daily	Lia Cover
6	Nova Daily	Joya Buss
the output should be

subscriber
Erwin Chitty
Joya Buss
Lavelle Phu
Lia Cover
Tonie Moreton
*/
CREATE PROCEDURE newsSubscribers()
BEGIN
    SELECT DISTINCT subscriber FROM
	(SELECT *
    FROM full_year
    UNION
    SELECT *
    FROM half_year) a
    WHERE newspaper REGEXP '.*daily.*'
    ORDER BY 1;
END

/*
CREATE PROCEDURE newsSubscribers()
BEGIN
    SELECT DISTINCT subscriber FROM
	(SELECT *
    FROM full_year
    UNION
    SELECT *
    FROM half_year) a
    WHERE newspaper LIKE '%daily%'
    ORDER BY 1;
END
*/
