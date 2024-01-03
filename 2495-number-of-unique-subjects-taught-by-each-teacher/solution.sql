-- -- Write your PostgreSQL query statement below
-- SELECT t.teacher_id, COUNT(DISTINCT t.subject_id) as cnt FROM Teacher t
-- GROUP BY t.teacher_id;

SELECT teacher_id, COUNT(*) AS cnt
FROM (
    SELECT teacher_id, subject_id
    FROM Teacher
    GROUP BY teacher_id, subject_id
) as Teacher
GROUP BY teacher_id
