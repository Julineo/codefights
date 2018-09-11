/*
https://app.codesignal.com/arcade/db/exotic-dishes/qwgoYcGh7SKipb5D3

You need to connect to a remote database and extract some information from it. The problem is, you don't know the name of the most important table in this database! You were told that it should start with the letter e and end with the letter s. Now you want to find all possible candidates, i. e. tables with such names, as well as their column names and their datatypes. It is guaranteed that at least one such table exists in the database.

You have already connected to the database named ri_db. Now you just need to implement a procedure that will find the information about the most important table in it, as described above. The resulting table should have the following structure:

tab_name: the name of the found table;
col_name: the name of the column in the found table;
data_type: the type of this column.
The rows in the output should be sorted first by the name of the table, then by the column names.

Example

For the following tables events

event_date	event_name
2016-10-08	Mum's birthday
2016-10-31	Halloween

and guestlist in the database

id	first_name	surname
1	John		Miller
2	Evelyn		Ross

the output should be

tab_name	col_name	data_type
events		event_date	date
events		event_name	varchar
*/
CREATE PROCEDURE findTable()
BEGIN
    DROP TABLE IF EXISTS tmp;
    CREATE TEMPORARY TABLE tmp
    SELECT TABLE_NAME
    FROM information_schema.TABLES 
    WHERE TABLE_NAME REGEXP "^e.*s$"
    AND TABLE_SCHEMA = "ri_db";

    SELECT TABLE_NAME tab_name, COLUMN_NAME col_name, DATA_TYPE data_type
    FROM information_schema.COLUMNS c
    WHERE TABLE_NAME IN (SELECT TABLE_NAME FROM tmp)
    AND TABLE_SCHEMA = "ri_db"
    ORDER BY 1, 2;
END
