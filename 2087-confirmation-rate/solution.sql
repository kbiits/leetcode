-- Write your PostgreSQL query statement below
SELECT s.user_id, ROUND(COUNT(c.user_id) filter (where c.action = 'confirmed')::numeric / COUNT(*)::numeric, 2) as confirmation_rate 
FROM Signups s
LEFT JOIN Confirmations c
ON c.user_id = s.user_id
GROUP BY s.user_id;
