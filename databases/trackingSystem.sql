/*
https://app.codesignal.com/arcade/db/selecting-what-to-select/NQ5LyGx2X57AmmJ2G

Your web application tracks the activities of its users using a tracking system. While a user hasn't logged in or signed up, all the user's actions are tracked using anonymous_id and the user_id is null, and afterwards they are tracked using the same anonymous_id and user_id. It is known that after a user logs in or signs up, the user_id is no longer null.

You are given the table tracks, which contains the following columns:

received_at - the unique timestamp of action;
event_name - the name of the action that was performed at this time;
anonymous_id - the anonymous ID of user;
user_id - the user ID, which can be null.
Your task is to find two events for each anonymous_id, which will be the column anonym_id in the returned table. Find the last event where the user was tracked only by anonymous_id (column last_null) and the first event that was tracked by user_id (column first_notnull). The resulting table should be sorted by anonym_id.

Note: It is not guaranteed that a user ever logs in or signs up. In this case, the column first_notnull should have a value of null. However, it is guaranteed that for each anonymous_id, there is at least one event where user_id equals null.

Example

For given table tracks

received_at	event_name	anonymous_id	user_id
2016-01-01 12:13:12	buttonClicked	1	NULL
2016-01-02 12:14:15	pageReloaded	3	NULL
2016-02-02 13:15:13	pageRendered	2	NULL
2016-02-03 13:15:23	commentWritten	3	NULL
2016-03-03 14:15:15	avatarUpdated	2	2
2016-03-04 14:15:24	statusUpdated	1	1

the output should be

anonym_id	last_null	first_notnull
1	buttonClicked	statusUpdated
2	pageRendered	avatarUpdated
3	commentWritten	NULL
*/

CREATE PROCEDURE trackingSystem()
BEGIN
    SELECT DISTINCT anonymous_id anonym_id,
    (SELECT event_name FROM tracks t1 WHERE t1.anonymous_id = t0.anonymous_id AND t1.user_id IS NULL ORDER BY received_at DESC LIMIT 1) last_null,
    (SELECT event_name FROM tracks t1 WHERE t1.anonymous_id = t0.anonymous_id AND t1.user_id IS NOT NULL ORDER BY received_at LIMIT 1) first_null
    FROM tracks t0
    ORDER BY 1
;
END

/*
CREATE PROCEDURE trackingSystem()
BEGIN
    SELECT a.*, b.first_notnull  FROM
	(SELECT t1.anonymous_id AS anonym_id, t1.event_name last_null
    FROM tracks t1
    WHERE t1.received_at IN (SELECT MAX(t0.received_at) FROM tracks t0 WHERE t0.anonymous_id = t1.anonymous_id AND t0.user_id IS NULL)) a
    LEFT JOIN
	(SELECT t2.anonymous_id AS anonym_id, t2.event_name first_notnull
    FROM tracks t2
    WHERE t2.received_at IN (SELECT MIN(t0.received_at) FROM tracks t0 WHERE t0.anonymous_id = t2.anonymous_id AND t0.user_id IS NOT NULL)) b
    on a.anonym_id = b.anonym_id
    ORDER BY 1
;
END
*/
