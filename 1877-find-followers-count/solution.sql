-- Write your PostgreSQL query statement below
SELECT f.user_id, COUNT(*) as followers_count FROM Followers f
GROUP BY f.user_id
ORDER BY f.user_id ASC;
