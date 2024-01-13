-- -- Write your PostgreSQL query statement below
SELECT  
    CONCAT(EXTRACT('year' FROM t.trans_date), '-', LPAD(EXTRACT('month' FROM t.trans_date)::text, 2, '0'::text)) as month,
    t.country,
    COUNT(*) as trans_count,
    COUNT(1) FILTER (WHERE t.state = 'approved') as approved_count,
    SUM(amount) as trans_total_amount,
    SUM(CASE WHEN state = 'approved' THEN amount ELSE 0 END) as approved_total_amount
FROM transactions t
GROUP BY month, country;
