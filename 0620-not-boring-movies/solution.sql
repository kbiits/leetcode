SELECT c.* FROM Cinema c
WHERE mod(c.id, 2) = 1 
AND c.description != 'boring'
ORDER BY c.rating DESC;
