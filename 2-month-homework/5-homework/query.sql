create table if not exists student
(
    id   uuid primary key default gen_random_uuid() not null,
    name varchar                                    not null,
    age int                                         not null
);


create table if not exists course
(
    id            uuid primary key default gen_random_uuid() not null,
    name          varchar not null,
);


create table if not exists student_course
(
    id uuid primary key not null default gen_random_uuid(),
    student_id uuid references student (id),
    course_id uuid references course (id)
);

create table if not exists grade
(
    id uuid primary key not null default gen_random_uuid(),
    stc_id uuid references student_course (id),
    mark int not null
);

--3

SELECT
    s.name AS student_name,
    AVG(g.mark) AS average_grade
FROM
    student s
        JOIN
    student_course sc ON s.id = sc.student_id
        JOIN
    grade g ON sc.id = g.stc_id
GROUP BY
    s.name;

--4
WITH youngest_students AS (
    SELECT
        c.name AS course_name,
        s.name AS student_name,
        s.age,
        RANK() OVER (PARTITION BY c.id ORDER BY s.age ASC) AS age_rank
    FROM
        student s
            JOIN
        student_course sc ON s.id = sc.student_id
            JOIN
        course c ON sc.course_id = c.id
)
SELECT
    course_name,
    student_name,
    age
FROM
    youngest_students
WHERE
    age_rank = 1
ORDER BY
    course_name, student_name;

--5
WITH course_avg_grades AS (
    SELECT
        c.id AS course_id,
        c.name AS course_name,
        AVG(g.mark) AS average_grade
    FROM
        course c
            JOIN
        student_course sc ON c.id = sc.course_id
            JOIN
        grade g ON sc.id = g.stc_id
    GROUP BY
        c.id, c.name
)
SELECT
    course_name,
    average_grade
FROM
    course_avg_grades
ORDER BY
    average_grade DESC
LIMIT 1;

--1
WITH course_avg_grades AS (
    SELECT
        c.id AS course_id,
        c.name AS course_name,
        AVG(g.mark) AS average_grade
    FROM
        course c
            JOIN
        student_course sc ON c.id = sc.course_id
            JOIN
        grade g ON sc.id = g.stc_id
    GROUP BY
        c.id, c.name
)
SELECT
    course_name,
    average_grade
FROM
    course_avg_grades
ORDER BY
    average_grade DESC
LIMIT 1;

-- 2
WITH max_grades_per_course AS (
    SELECT
        sc.course_id,
        MAX(g.mark) AS max_grade
    FROM
        student_course sc
            JOIN
        grade g ON sc.id = g.stc_id
    GROUP BY
        sc.course_id
),
     best_students AS (
         SELECT
             c.name AS course_name,
             s.name AS student_name,
             g.mark AS grade
         FROM
             max_grades_per_course mg
                 JOIN
             student_course sc ON mg.course_id = sc.course_id
                 JOIN
             grade g ON sc.id = g.stc_id AND g.mark = mg.max_grade
                 JOIN
             student s ON sc.student_id = s.id
                 JOIN
             course c ON sc.course_id = c.id
     )
SELECT
    course_name,
    student_name,
    grade
FROM
    best_students
ORDER BY
    course_name, student_name;


