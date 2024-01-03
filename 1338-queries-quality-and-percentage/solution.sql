-- Write your PostgreSQL query statement below
SELECT 
    q.query_name, 
    ROUND(SUM(q.rating::numeric / q.position::numeric) / COUNT(*), 2) as quality,  
    ROUND(COUNT(*) FILTER (WHERE q.rating < 3)::numeric * 100 / COUNT(*)::numeric, 2) as poor_query_percentage 
FROM Queries q
WHERE q.query_name IS NOT NULL
GROUP BY q.query_name;
