SELECT s.name as name
FROM salesperson as s
LEFT JOIN orders as o
using(sales_id)
LEFT JOIN company as c
using(com_id)
GROUP BY s.sales_id
HAVING sum(c.name = 'RED') = 0 or sum(c.name is null) > 0