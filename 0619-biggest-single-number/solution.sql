-- -- Write your PostgreSQL query statement below
-- WITH CTE AS (
--     SELECT mn.num FROM MyNumbers mn
--     GROUP BY mn.num
--     HAVING COUNT(*) = 1
-- ) SELECT MAX(c.num) as num FROM CTE c;

SELECT (SELECT mn.num as num 
FROM MyNumbers mn
GROUP BY mn.num
HAVING COUNT(*) = 1
ORDER BY mn.num DESC LIMIT 1);
