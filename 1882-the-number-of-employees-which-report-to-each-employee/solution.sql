-- Write your PostgreSQL query statement below
SELECT 
    e.reports_to as employee_id, 
    (SELECT e2.name FROM Employees e2 WHERE e2.employee_id = e.reports_to) as name, 
    COUNT(*) as reports_count, 
    ROUND(AVG(e.age)) as average_age
FROM Employees e
WHERE e.reports_to IS NOT NULL
GROUP BY e.reports_to
ORDER BY e.reports_to ASC;

