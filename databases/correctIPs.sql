/*
https://app.codesignal.com/arcade/db/regular-paradise/XYXE3mXLuZ243yqmb

The company you work for has a database with the IP addresses of all the company's computers in it. When a new computer is purchased, an employee adds its IP address to this database. Unfortunately, you've just discovered that there is no address validation, so some of the records are incorrect.

Now your boss wants you to write a program that will find and retain only the correct records from this table. A record is correct if the IP it contains is a valid IPv4 address, and either the first or the second component in the host portion is a two-digit number.

The ips table contains the following columns:

id: the unique ID of the computer;
ip: the unique IP address of the computer.
Your task is to create a new table that contains only the correct records from the given ips table. This table should be sorted by id.

Example

For the given table ips

id	ip
4	"1.1.1.1"
3	"1.111.111.11"
2	"11.11.11.11"
1	"11.11.11.11.11.11"
5	"11.11.11.111"
the output should be

id	ip
2	"11.11.11.11"
3	"1.111.111.11"
5	"11.11.11.111"
*/

CREATE PROCEDURE correctIPs()
BEGIN
    select * 
    from ips 
    where ip regexp "^[1-2]{0,1}[0-9]{1,2}\.[1-2]{0,1}[0-9]{1,2}\.([0-9]{2}\.[1-2]{0,1}[0-9]{1,2}|[1-2]{0,1}[0-9]{1,2}[.][0-9]{2})$"
    order by id;
END
