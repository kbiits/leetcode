-- Write your PostgreSQL query statement below
SELECT 
    p.product_id, 
    ROUND(
        COALESCE(SUM(p.price * us.units )::numeric / SUM(us.units)::numeric, 0)
    , 2) as average_price 
FROM Prices p
LEFT JOIN UnitsSold us ON us.product_id = p.product_id
WHERE us.purchase_date BETWEEN p.start_date AND p.end_date
OR us.product_id IS NULL
GROUP BY p.product_id;
