create table Problems (
    id uuid primary key not null default gen_random_uuid(),
    description varchar(225) not null ,
    difficulty varchar(12),
    IsSolved bool default false
);

create table Users (
    id uuid primary key not null default gen_random_uuid(),
    progress int default 0,
    username varchar(50),
    email varchar(30),
    password varchar(50)
);

create table Solved_problems (
    user_id uuid references Users(id),
    problem_id uuid references Problems(id)
);




