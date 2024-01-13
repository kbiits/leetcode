-- Write your PostgreSQL query statement below

SELECT ROUND(
    COUNT(*)::numeric / 
    (SELECT COUNT(distinct player_id) FROM activity)::numeric
, 2) as fraction
FROM Activity a
INNER JOIN activity b
ON
    a.player_id = b.player_id AND
    a.event_date = (b.event_date + 1)
WHERE
    (a.player_id, least(a.event_date, b.event_date)) IN (
        SELECT player_id, min(event_date) FROM activity group by player_id
    );

-- SELECT player_id, min(event_date) FROM activity x group by player_id;
