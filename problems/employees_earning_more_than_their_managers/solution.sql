# Write your MySQL query statement below
SELECT a.name AS Employee FROM Employee a INNER JOIN Employee b ON a.ManagerId = b.Id WHERE a.Salary > b.Salary;