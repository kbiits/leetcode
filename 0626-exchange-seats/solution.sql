-- -- Write your PostgreSQL query statement below
-- SELECT 
--     s.id, 
--     COALESCE(
--         (SELECT s2.student 
--             FROM seat s2 
--             WHERE 
--                 s2.id = (CASE WHEN s.id % 2 = 1 THEN (s.id + 1) ELSE (s.id - 1) END)
--         ),
--         s.student
--     ) as student
-- FROM seat s
-- ORDER BY s.id ASC;

SELECT 
    CASE 
        WHEN id = (SELECT MAX(id) FROM seat) AND id % 2 = 1
            THEN id 
        WHEN id % 2 = 1
            THEN id + 1
        ELSE id - 1
    END AS id,
    student
FROM seat
ORDER BY id
