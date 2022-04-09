select t.customer_number from Orders o, (
    select customer_number, COUNT(*) as num from Orders group by customer_number ORDER BY num DESC LIMIT 1
) t where t.customer_number = o.customer_number LIMIT 1;
