SELECT t1.id
FROM Weather t1
INNER JOIN Weather t2
ON TO_DAYS(t1.recordDate) = TO_DAYS(t2.recordDate) + 1
WHERE t1.temperature > t2.temperature
