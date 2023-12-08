-- Write your PostgreSQL query statement below
SELECT 
    a.machine_id, ROUND(AVG(b.timestamp - a.timestamp)::numeric, 3) as processing_time
FROM Activity a INNER JOIN Activity b 
ON
    a.machine_id = b.machine_id AND 
    a.process_id = b.process_id AND 
    a.activity_type = 'start' AND 
    b.activity_type = 'end'
GROUP BY a.machine_id;
