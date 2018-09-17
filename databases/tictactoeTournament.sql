/*
https://app.codesignal.com/arcade/db/a-table-of-desserts/p9HAu8mLsR46uqWBs

The latest Tic-Tac-Toe World Championship has just concluded. It was a big success and attracted a lot of participants! Now everyone is waiting for the results. As a member of the judging committee, you've been tasked with creating a championship leaderboard.

The results of all the tic-tac-toe matches are stored in the results table, which has the following structure:

timestamp: the date and time of the game;
name_naughts: the name of the player that played with naughts (O);
name_crosses: the name of the player that played with crosses (X);
board: the state of the tic-tac-toe board at the end of the game, in the format described below.
The board is a string of 9 characters that represent the state of board's 9 cells at the end of the match. The first 3 characters represent the first (upper) row, the second 3 characters represent the second row, etc. The character is O if the respective board cell contained a naught, X if it contained a cross, or . if the cell was empty at the end of the game.

For example, this board

 | |O
 |O| 
X|X|X
is represented by the string ..O.O.XXX.

It is guaranteed that the opposing players have different names. It's also guaranteed that each board represents a valid terminal state for a tic-tac-toe game, meaning that either of the players wins or it's a draw.

Players are awarded points based on their performance: They get 2 points for each game they win, 1 point for a draw, and 0 points if they lose.

Create a leaderboard with the following format:

Given the results table, compose a results table that has the following six columns: name, points, played, won, draw, and lost , containing the player names, their points calculated as described above, the number of games they have played, the number of games they have won, and the number of games they have lost, respectively. The table should be sorted in descending order by the points, then in ascending order by the total number of played games, then in descending order by the number of victories, and then in ascending order by player names.

Example

For the following table results

timestamp		name_naughts		name_crosses		board
2017-05-01 08:06:00	Georgine Greely		Earnestine Kunzman	XO.X.OXXO
2017-05-01 11:20:00	Earnestine Kunzman	Georgine Greely		.O.OOXXXX
2017-05-01 16:48:00	Renee Fortenberry	Georgine Greely		XOOXXO..X
2017-05-02 10:57:00	Renee Fortenberry	Georgine Greely		OXXXOOXOX
2017-05-03 01:55:00	Georgine Greely		Renee Fortenberry	.X.OX.OX.
2017-05-03 04:29:00	Earnestine Kunzman	Renee Fortenberry	OOXXXOXXO
2017-05-04 14:29:00	Renee Fortenberry	Earnestine Kunzman	OOX.X.X..
2017-05-04 16:00:00	Earnestine Kunzman	Renee Fortenberry	OXOOXXXOX
the output should be

name			points	played	won	draw	lost
Renee Fortenberry	6	6	2	2	2
Earnestine Kunzman	5	5	2	1	2
Georgine Greely	5	5	2	1	2
*/

CREATE PROCEDURE tictactoeTournament()
BEGIN
    SELECT name, SUM(points) points, SUM(played) played, SUM(won) won, SUM(draw) draw, SUM(lost) lost
    FROM (
    SELECT name, CASE WHEN won = 1 THEN 2 WHEN draw = 1 THEN 1 ELSE 0 END points, 1 played, won, draw, lost
    FROM (
    SELECT name_naughts name, leftWin won, IF(rightWin = leftWin, 1, 0) draw,    IF(rightWin = 1, 1, 0) lost
    FROM
    (SELECT *,
        CASE 
            WHEN board REGEXP "O..O..O.."
            OR board REGEXP ".O..O..O."
            OR board REGEXP "..O..O..O"
            OR board REGEXP "OOO......"
            OR board REGEXP "...OOO..."
            OR board REGEXP "......OOO"
            OR board REGEXP "O...O...O"
            OR board REGEXP "..O.O.O.."
            THEN 1
            ELSE 0
        END leftWin,
        CASE 
            WHEN board REGEXP "X..X..X.."
            OR board REGEXP ".X..X..X."
            OR board REGEXP "..X..X..X"
            OR board REGEXP "XXX......"
            OR board REGEXP "...XXX..."
            OR board REGEXP "......XXX"
            OR board REGEXP "X...X...X"
            OR board REGEXP "..X.X.X.."
            THEN 1
            ELSE 0
        END rightWin
        FROM results) a
    UNION ALL
    SELECT name_crosses name, rightWin won, IF(rightWin = leftWin, 1, 0) draw, IF(leftWin = 1, 1, 0) lost
    FROM
    (SELECT *,
        CASE 
            WHEN board REGEXP "O..O..O.."
            OR board REGEXP ".O..O..O."
            OR board REGEXP "..O..O..O"
            OR board REGEXP "OOO......"
            OR board REGEXP "...OOO..."
            OR board REGEXP "......OOO"
            OR board REGEXP "O...O...O"
            OR board REGEXP "..O.O.O.."
            THEN 1
            ELSE 0
        END leftWin,
        CASE 
            WHEN board REGEXP "X..X..X.."
            OR board REGEXP ".X..X..X."
            OR board REGEXP "..X..X..X"
            OR board REGEXP "XXX......"
            OR board REGEXP "...XXX..."
            OR board REGEXP "......XXX"
            OR board REGEXP "X...X...X"
            OR board REGEXP "..X.X.X.."
            THEN 1
            ELSE 0
        END rightWin
        FROM results) b
        ) ab
        ) c
        GROUP BY name
        ORDER BY points DESC, played, won DESC, name;
END
