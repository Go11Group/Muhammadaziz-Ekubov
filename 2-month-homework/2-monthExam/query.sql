CREATE TABLE Users (
    user_id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    birthday TIMESTAMP NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at Bigint default 0
);


CREATE TABLE Courses (
    course_id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at bigint default 0
);

e
CREATE TABLE Lessons (
    lesson_id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at bigint default 0,
    FOREIGN KEY (course_id) REFERENCES Courses(course_id)
);


CREATE TABLE Enrollments (
    enrollment_id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    course_id UUID NOT NULL,
    enrollment_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at bigint default 0,
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (course_id) REFERENCES Courses(course_id)
);

INSERT INTO Users (name, email, birthday, password)
VALUES
    ('John Doe', 'johndoe1@example.com', '1990-01-01', 'password1'),
    ('Jane Smith', 'janesmith2@example.com', '1992-02-02', 'password2'),
    ('Jim Brown', 'jimbrown3@example.com', '1994-03-03', 'password3'),
    ('Alice Johnson', 'alicejohnson4@example.com', '1988-04-04', 'password4'),
    ('Bob Davis', 'bobdavis5@example.com', '1986-05-05', 'password5'),
    ('Charlie White', 'charliewhite6@example.com', '1991-06-06', 'password6'),
    ('Diana King', 'dianaking7@example.com', '1993-07-07', 'password7'),
    ('Edward Harris', 'edwardharris8@example.com', '1989-08-08', 'password8'),
    ('Fiona Martin', 'fionamartin9@example.com', '1995-09-09', 'password9'),
    ('George Clark', 'georgeclark10@example.com', '1987-10-10', 'password10'),
    ('Hannah Lewis', 'hannahlewis11@example.com', '1990-11-11', 'password11'),
    ('Ian Walker', 'ianwalker12@example.com', '1991-12-12', 'password12'),
    ('Jane Lee', 'janelee13@example.com', '1985-01-13', 'password13'),
    ('Kevin Hall', 'kevinhall14@example.com', '1992-02-14', 'password14'),
    ('Laura Allen', 'lauraallen15@example.com', '1993-03-15', 'password15');

INSERT INTO Courses (title, description)
VALUES
    ('Course 1', 'Description for course 1'),
    ('Course 2', 'Description for course 2'),
    ('Course 3', 'Description for course 3'),
    ('Course 4', 'Description for course 4'),
    ('Course 5', 'Description for course 5');

INSERT INTO Lessons (course_id, title, content)
VALUES
    ((SELECT course_id FROM Courses WHERE title='Course 1'), 'Lesson 1 for Course 1', 'Content for lesson 1 of course 1'),
    ((SELECT course_id FROM Courses WHERE title='Course 1'), 'Lesson 2 for Course 1', 'Content for lesson 2 of course 1'),
    ((SELECT course_id FROM Courses WHERE title='Course 2'), 'Lesson 1 for Course 2', 'Content for lesson 1 of course 2'),
    ((SELECT course_id FROM Courses WHERE title='Course 2'), 'Lesson 2 for Course 2', 'Content for lesson 2 of course 2'),
    ((SELECT course_id FROM Courses WHERE title='Course 3'), 'Lesson 1 for Course 3', 'Content for lesson 1 of course 3'),
    ((SELECT course_id FROM Courses WHERE title='Course 3'), 'Lesson 2 for Course 3', 'Content for lesson 2 of course 3'),
    ((SELECT course_id FROM Courses WHERE title='Course 4'), 'Lesson 1 for Course 4', 'Content for lesson 1 of course 4'),
    ((SELECT course_id FROM Courses WHERE title='Course 4'), 'Lesson 2 for Course 4', 'Content for lesson 2 of course 4'),
    ((SELECT course_id FROM Courses WHERE title='Course 5'), 'Lesson 1 for Course 5', 'Content for lesson 1 of course 5'),
    ((SELECT course_id FROM Courses WHERE title='Course 5'), 'Lesson 2 for Course 5', 'Content for lesson 2 of course 5');


INSERT INTO Enrollments (user_id, course_id, enrollment_date)
VALUES
    ((SELECT user_id FROM Users WHERE email='johndoe1@example.com'), (SELECT course_id FROM Courses WHERE title='Course 1'), '2023-01-01'),
    ((SELECT user_id FROM Users WHERE email='janesmith2@example.com'), (SELECT course_id FROM Courses WHERE title='Course 2'), '2023-01-02'),
    ((SELECT user_id FROM Users WHERE email='jimbrown3@example.com'), (SELECT course_id FROM Courses WHERE title='Course 3'), '2023-01-03'),
    ((SELECT user_id FROM Users WHERE email='alicejohnson4@example.com'), (SELECT course_id FROM Courses WHERE title='Course 4'), '2023-01-04'),
    ((SELECT user_id FROM Users WHERE email='bobdavis5@example.com'), (SELECT course_id FROM Courses WHERE title='Course 5'), '2023-01-05'),
    ((SELECT user_id FROM Users WHERE email='charliewhite6@example.com'), (SELECT course_id FROM Courses WHERE title='Course 1'), '2023-01-06'),
    ((SELECT user_id FROM Users WHERE email='dianaking7@example.com'), (SELECT course_id FROM Courses WHERE title='Course 2'), '2023-01-07'),
    ((SELECT user_id FROM Users WHERE email='edwardharris8@example.com'), (SELECT course_id FROM Courses WHERE title='Course 3'), '2023-01-08'),
    ((SELECT user_id FROM Users WHERE email='fionamartin9@example.com'), (SELECT course_id FROM Courses WHERE title='Course 4'), '2023-01-09'),
    ((SELECT user_id FROM Users WHERE email='georgeclark10@example.com'), (SELECT course_id FROM Courses WHERE title='Course 5'), '2023-01-10'),
    ((SELECT user_id FROM Users WHERE email='hannahlewis11@example.com'), (SELECT course_id FROM Courses WHERE title='Course 1'), '2023-01-11'),
    ((SELECT user_id FROM Users WHERE email='ianwalker12@example.com'), (SELECT course_id FROM Courses WHERE title='Course 2'), '2023-01-12'),
    ((SELECT user_id FROM Users WHERE email='janelee13@example.com'), (SELECT course_id FROM Courses WHERE title='Course 3'), '2023-01-13'),
    ((SELECT user_id FROM Users WHERE email='kevinhall14@example.com'), (SELECT course_id FROM Courses WHERE title='Course 4'), '2023-01-14'),
    ((SELECT user_id FROM Users WHERE email='lauraallen15@example.com'), (SELECT course_id FROM Courses WHERE title='Course 5'), '2023-01-15');



