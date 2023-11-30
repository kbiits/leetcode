-- Write your PostgreSQL query statement below
SELECT p.product_name, s.year, s.price
FROM sales s
INNER JOIN product p on s.product_id = p.product_id;
