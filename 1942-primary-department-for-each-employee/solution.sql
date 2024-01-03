-- Write your PostgreSQL query statement below
SELECT temp.employee_id, e.department_id FROM (
    SELECT e.employee_id, COUNT(*) as cnt FROM Employee e
    GROUP BY e.employee_id
) temp 
INNER JOIN employee e ON e.employee_id = temp.employee_id
WHERE temp.cnt = 1 OR e.primary_flag = 'Y';

