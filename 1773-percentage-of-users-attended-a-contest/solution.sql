-- Write your PostgreSQL query statement below
select r.contest_id, ROUND(COUNT(*)::numeric / (SELECT COUNT(*) FROM users)::numeric * 100, 2) as percentage from register r
INNER JOIN users u ON r.user_id = u.user_id
GROUP BY r.contest_id
ORDER BY percentage DESC, r.contest_id ASC;
