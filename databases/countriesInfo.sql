/*
Your task is to return a new table that contains the number of countries in the given list, along with their average and total population, in columns titled number, average and total.

Example

For the following table countries

name	continent	population
Grenada	North America	103328
Monaco	Europe	38400
San Marino	Europe	33005
Vanuatu	Australia	277500
the output should be

number	average	total
4	113058.2500	452233
*/
CREATE PROCEDURE countriesInfo()
BEGIN
	SELECT COUNT(*) number, AVG(population) average, SUM(population) total
    FROM countries;
END
