# Write your MySQL query statement below
SELECT 
    t.*, 
    CASE 
        WHEN t.x + t.y > t.z AND t.y + t.z > t.x AND t.z + t.x > t.y THEN 'Yes'    
        ELSE 'No'
    END as triangle
FROM Triangle t;
