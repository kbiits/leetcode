# Write your MySQL query statement below
select t.class FROM (
    select c.class, count(*) as num from Courses c group by c.class
) t where t.num > 4;