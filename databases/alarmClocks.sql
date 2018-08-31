/*
https://app.codesignal.com/arcade/db/time-river-revisited/G9zoZjTzk6JpJGFbd

You are developing an alarm clock app that works as follows: the user can set a date and a time, and the app will ring every week at the given time starting at the given date until the end of the current year.

The starting date is the only record in the userInput table of the following structure:

input_date: the date and time of the first alarm (of a DATETIME type).
Given the table, your task is to compose the resulting table with a single column alarm_date. This column should contain all dates (including time) when the alarm clock will ring in ascending order.

Example

For the following table userInput

input_date
2016-10-23 22:00:00
the output should be

alarm_date
2016-10-23 22:00:00
2016-10-30 22:00:00
2016-11-06 22:00:00
2016-11-13 22:00:00
2016-11-20 22:00:00
2016-11-27 22:00:00
2016-12-04 22:00:00
2016-12-11 22:00:00
2016-12-18 22:00:00
2016-12-25 22:00:00
*/

CREATE PROCEDURE alarmClocks()
BEGIN
  set @x = (select input_date from userInput);
  set @y = (select year(input_date) from userInput);
  drop table if exists t;
  create temporary table t (select @x);
  while week(@x) <= 51 and year(@x) = @y do
    set @x = (select date_add(@x, interval 7 day));
    insert t select @x;
  end while;

  select `@x` alarm_date from t;
END
