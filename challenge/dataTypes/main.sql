/* https://codefights.com/challenge/YNefPAfkrBHy2Sfu5
You have a flat file that needs to be imported into your database but it looks like there was a problem with parsing column types. As a result, the flat file was imported as a table flat_data with two columns:

id - the unique data ID;
data - a string containing the data which originally had one of the following types: (STRING, DECIMAL, DATETIME, INT, BIT).
You want to recover the original types of the values from the data column, but you'll need to figure it out according to how the data looks as a string. Here are some rules to keep in mind:

The types have the following precedence: DATETIME > DECIMAL > BIT > INT > STRING (i.e. 0 or 1 should be considered BIT since its precedence is higher than INT)
DATETIME format is yyyy-mm-dd.
You can assume that all any month can have up to 31 days.
All decimals must contain exactly one . and all numbers shouldn't have any leading zeros.
Decimals can't start with .
All numbers can be prefixed with single + or -.
Given the flat_data table, your task is to compose the resulting table with two columns id and data_type for the data id and its type respectively. The output should be sorted by id in ascending order.

Example

For given table flat_data

id	data
1	2 is better than 1
2	42.00
3	2011-11-11
4	1532
5	1
6	0
7	11/11/2011
the output should be

id	data_type
1	STRING
2	DECIMAL
3	DATETIME
4	INT
5	BIT
6	BIT
7	STRING
*/
CREATE PROCEDURE dataTypes()
BEGIN
	SELECT id, 
        CASE
            WHEN    DATA REGEXP '^[01]{1}$' THEN "BIT"
            WHEN    DATA REGEXP '^[+-]?[1-9]{1}[0-9]*$' THEN "INT"
            WHEN    DATA IN ('-0', '+0') THEN "INT"
            WHEN    DATA REGEXP '^[+-]?([1-9]{1}[0-9]*|[0-9]{1})[.]{1}[0-9]*$' THEN "DECIMAL"
            WHEN    DATA REGEXP '[0-9]{1}[0-9]{3}[-]{1}(0[1-9]{1}|1[012]{1})[-]{1}(0[1-9]{1}|[12]{1}[0-9]{1}|3[01]{1})' THEN "DATETIME"
            ELSE
                "STRING"
        END AS data_type
    FROM flat_data;
END
