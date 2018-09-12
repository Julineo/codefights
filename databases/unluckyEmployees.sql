/*
https://app.codesignal.com/arcade/db/between-join-and-select/W9abjqqFan2Y5NGwg

Your company has fallen on hard times, and you have to let some of your employees go. You figure it will be easier to fire an entire department all at once, so now you want to determine which department it's going to be.

Information about your employees and departments is stored in two tables, Employee and Department, respectively. Here are their structures:

Department:
id: unique department id
name: department name
Employee:
id: unique employee id
full_name: employee's full name
department: foreign key referencing Department.id
salary: employee's salary
To choose the unfortunate department, you set a number of criteria: you are willing to get rid of any department that has no more than 5 employees. Among these smaller departments, you will consider those where the total salary of all its employees is maximal. Lastly, to make a tough situation more fair, you decide to make the final choice from the remaining departments at random. Thus, you'd like to build the following table of departments:

select all departments with less than 6 employees;
sort these departments by the total salary of its workers in descending order (in the case of a tie, the department with the greatest number of employees should go first; if it's still not enough to break a tie, the department with the smallest id should go first);
cross out the departments at the even rows and leave only those in the odd positions, to consider them more thoroughly afterwards.
Given tables Employee and Department, build a table as described below. The resulting table should have columns dep_name, emp_number and total_salary and be sorted as described above.

Example

For the following tables Department

id	name
1	IT
2	HR
3	Sales

and Employee

id	full_name		salary	department
1	James Smith		20	1
2	John Johnson		13	1
3	Robert Jones		15	1
4	Michael Williams	15	1
5	Mary Troppins		17	1
8	Penny Old		14	2
9	Richard Young		17	2
10	Drew Rich		50	3

the output should be

dep_name	emp_number	total_salary
IT			5	80
HR			2	31

All three departments have 5 or fewer employees, so they are all candidates to be fired. When sorted in descending order by total_salary, the Sales department becomes the second (i.e. is located at an even row), so it's not present in the resulting table.
*/
CREATE PROCEDURE unluckyEmployees()
BEGIN
    DROP TEMPORARY TABLE IF EXISTS tmp;
    CREATE TEMPORARY TABLE tmp (
        ID INT NOT NULL AUTO_INCREMENT,
        dep_name VARCHAR(255),
        emp_number INT,
        total_salary INT,
        PRIMARY KEY (ID) 
    );

    INSERT tmp (dep_name, emp_number, total_salary)
	SELECT d.name dep_name, COUNT(salary) emp_number, SUM(COALESCE(salary, 0)) total_salary
    FROM Department d
    LEFT JOIN Employee e ON d.id = e.department
    GROUP BY dep_name
    HAVING COUNT(salary) < 6
    ORDER BY total_salary DESC, emp_number DESC, d.id;

    SELECT dep_name, emp_number, total_salary FROM tmp
    WHERE ID % 2 != 0;
END
