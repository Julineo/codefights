/*
https://app.codesignal.com/arcade/db/selecting-what-to-select/ZWiqurZME6zN9argE

There are many employees at your company, with a wide range of salaries. You've decided to find out the difference in salary between the employees who make the most and the employees who make the least.

You store information about your employees in the table employees, which has the structure:

id: the unique employee ID;
name: the employee's name;
salary: the employee's salary - a positive integer.
The difference between the sum of the highest salaries and the sum of lowest salaries will give you the information you want. So, given the table employees, build the resulting table as follows: The table should contain only one column difference and only one row, which will contain the difference between sum of the highest and the sum of lowest salaries.

Example

For the following table employees

id	name	salary
1	John	1200
2	Bill	1000
3	Mike	1300
4	Katy	1200
5	Brendon	1300
6	Amanda	900

the output should be

difference
1700

There are two highest salaries (1300) and one lowest salary (900). 1300 * 2 - 900 * 1 = 1700.
*/
CREATE PROCEDURE salaryDifference()
BEGIN
    SELECT COALESCE(SUM(a), 0) difference FROM(
        SELECT SUM(salary) a FROM employees
        WHERE salary IN (
            SELECT MAX(salary) FROM employees
            )
        UNION
        SELECT -SUM(salary) FROM employees
        WHERE salary IN (
            SELECT MIN(salary) FROM employees
            )
    ) t;
END
