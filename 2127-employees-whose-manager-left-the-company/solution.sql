-- Write your PostgreSQL query statement below
SELECT e.employee_id FROM Employees e
LEFT JOIN employees e2 ON e.manager_id = e2.employee_id
WHERE e.manager_id IS NOT NULL AND e2.employee_id IS NULL AND e.salary < 30000
ORDER BY e.employee_id;

-- SELECT employee_id
-- FROM Employees
-- WHERE salary < 30000
-- AND manager_id NOT IN (
--   SELECT employee_id FROM Employees
-- )
-- ORDER BY employee_id;
