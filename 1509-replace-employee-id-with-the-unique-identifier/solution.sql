-- Write your PostgreSQL query statement below
SELECT eu.unique_id, e.name
FROM employees e
LEFT JOIN employeeUNI eu ON e.id = eu.id;
