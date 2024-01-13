SELECT ROUND(SUM(temp.immediate)::numeric * 100 / COUNT(*)::numeric, 2) as immediate_percentage FROM (
    SELECT 
        CASE WHEN 
            MIN(d.order_date) = min(d.customer_pref_delivery_date)
        THEN
            1
        ELSE 0
        END as immediate
    FROM Delivery d
    GROUP BY d.customer_id
) temp;
