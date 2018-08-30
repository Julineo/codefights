/*
https://app.codesignal.com/arcade/db/when-was-it-the-case/CZBMSWZTZCHc4rcFo

You have a table scores that contains information about a series of soccer games. Your goal is to determine the winner of the series. A team is declared the winner if it won more games than the other team. If both teams had the same number of wins, then the winner is determined by the better goal difference (the difference between the goals that a team scores and the goals that the opposing team scores on them over all the games). If the goal differences are equal, the winner is the team that scored more goals during away games (i.e. games when it was not the host team). The result is the index of the winning team. If the above criteria are not sufficient for determining the winner, return 0.

The scores table contains the following columns:

match_id - the unique ID of the match;
first_team_score - the score of the 1st team in the current match;
second_team_score - the score of the 2nd team in the current match;
match_host - the team that is the host of the match (can be 1 or 2).
Your task is to return a new table with a single column winner, which can contain 1, 2, or 0.

Example

For given table scores

match_id first_team_score second_team_score match_host
1	 3		  2			1
2	 2		  1			2
3	 1		  2			1
4	 2		  1			2

the output should be

winner
1

*/

CREATE PROCEDURE soccerGameSeries()
BEGIN
    select case
            when w1 > w2 then 1
            when w1 < w2 then 2
            when g1 > g2 then 1
            when g1 < g2 then 2
            when o1 > o2 then 1
            when o2 < o1 then 2
            else 0
           end
    as winner
    from (
        select
            sum( first_team_score > second_team_score )   as w1,
            sum( first_team_score < second_team_score )   as w2,
            sum( first_team_score )                       as g1,
            sum( second_team_score )                      as g2,
            sum(if(match_host = 2, first_team_score, 0))  as o1,
            sum(if(match_host = 1, second_team_score, 0)) as o2
       from scores
    )b;
END

/*
CREATE PROCEDURE soccerGameSeries()
BEGIN
    set @wins = (select case when sum(first) > sum(second) then 1
                when sum(first) < sum(second) then 2
                else 0
            end as wins
    from (
        select if (first_team_score > second_team_score, 1, 0) first,
            if (first_team_score < second_team_score, 1, 0) second
        from scores
    )b);

	set @goals = (select case when sum(first_team_score) > sum(second_team_score) then 1
                when sum(first_team_score) < sum(second_team_score) then 2
                else 0
            end as goals
    from scores);

    set @goalsOp = (select case when sum(first) > sum(second) then 1
                when sum(first) < sum(second) then 2
                else 0
            end as goalsOp
    from ( 
        select sum(if ( match_host = 2, first_team_score, 0)) first,
                sum(if ( match_host = 1, second_team_score, 0)) second
        from scores
    )b);

    select
    case
        when @goals > 0     then @goals
        when @wins > 0      then @wins
        when @goalsOp > 0   then @goalsOp
        else 0
    end as winner;
END
*/
