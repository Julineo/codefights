/*You want to make a tabular view of your favorite football players with a couple of stats from each game. You got some data however it is not complete. Now in order to make the view better you decided to order it.

The player_statistics table contains the following columns:

player_name - the unique name of the player;
pass - the number of passes;
run - the number of kilometers player ran;
catch - the number of catches.
Your task is to order it by the completeness of data in columns, firstly goes most filled column and so on, lastly by name. If columns have same number of filled values, priority go in following order pass > run > catch > player_name. Columns pass, run, catch should be ordered in descending order, column player_name - is ascending order.

Example

For given table player_statistics

player_name	pass	run	catch
Gareth Bale	30	500	20
Ryan Giggs	35	NULL	30
Sergio Ramos	28	200	40
Wayne Rooney	30	400	40
the output should be

player_name	pass	run	catch
Ryan Giggs	35	NULL	30
Wayne Rooney	30	400	40
Gareth Bale	30	500	20
Sergio Ramos	28	200	40
Columns pass and catch are equally complete, so we sort first by pass, then by catch, and later by run, because it's less complete.

*/

CREATE PROCEDURE playerStatistics()
BEGIN
    set @1 = (select COUNT(pass) from player_statistics where pass is not null);
    set @2 = (select COUNT(run) from player_statistics where run is not null);
    set @3 = (select COUNT(catch) from player_statistics where catch is not null);

    set @o =
    (case
        when @2 > @3 and @3 > @1 then 'run desc, catch desc, pass desc'
        when @3 > @2 and @2 > @1 then 'catch desc, run desc, pass desc'
        when @3 > @1 and @1 > @2 then 'catch desc, pass desc, run desc'
        when @2 > @1 and @1 > @3 then 'run desc, pass desc, catch desc'
        when @1 > @3 and @3 > @2 then 'pass desc, catch desc, run desc'    
        when @1 > @2 and @2 > @3 then 'pass desc, run desc, catch desc'
     
        when @3 = @1 and @1 > @2 then 'pass desc, catch desc, run desc'
        when @2 = @3 and @3 > @1 then 'run desc, catch desc, pass desc'
        when @1 = @3 and @3 > @2 then 'pass desc, catch desc, run desc'
     
        when @3 = @1 and @1 < @2 then 'run desc, pass desc, catch desc'
        when @2 = @3 and @3 < @1 then 'pass desc, run desc, catch desc'
        when @1 = @3 and @3 < @2 then 'run desc, pass desc, catch desc'
     
        when @1 = @2 and @2 = @3 then 'pass desc, run desc, catch desc'
        else 'pass desc, run desc, catch desc'
    end);

    set @q = concat('SELECT * FROM player_statistics ORDER BY ', @o, ', player_name asc;');
    /*select @query;*/

    PREPARE stmt FROM @q;
    EXECUTE stmt;
END
