-- Write your PostgreSQL query statement below
SELECT e.name FROM Employee e
INNER JOIN Employee e2 ON e.id = e2.managerId
GROUP BY e.id, e.name
HAVING COUNT(e2.id) >= 5;
