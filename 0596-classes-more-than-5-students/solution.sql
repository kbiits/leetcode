-- Write your PostgreSQL query statement below
SELECT c.class FROM Courses c
GROUP BY c.class
HAVING COUNT(*) >= 5;
