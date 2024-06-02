create table book
(
    id          int primary key,
    name        varchar not null,
    page        int4,
    author_name varchar
);

create table author
(
    id   int primary key,
    name varchar not null
);

INSERT INTO book (id,name, page, author_name) VALUES
    (1,'To Kill a Mockingbird', NULL, 'Harper Lee'),
    (2,'1984', NULL, 'George Orwell'),
    (3,'The Great Gatsby', NULL, 'F. Scott Fitzgerald'),
    (4,'Pride and Prejudice', NULL, 'Jane Austen'),
    (5,'The Catcher in the Rye', NULL, 'J.D. Salinger'),
    (6,'The Lord of the Rings', NULL, 'J.R.R. Tolkien'),
    (7,'Harry Potter and the Sorcerer''s Stone', NULL, 'J.K. Rowling'),
    (8,'The Hobbit', NULL, 'J.R.R. Tolkien'),
    (9,'The Da Vinci Code', NULL, 'Dan Brown'),
    (10,'The Alchemist', NULL, 'Paulo Coelho'),
    (11,'The Hunger Games', NULL, 'Suzanne Collins'),
    (12,'The Hitchhiker''s Guide to the Galaxy', NULL, 'Douglas Adams'),
    (13,'The Chronicles of Narnia', NULL, 'C.S. Lewis'),
    (14,'Animal Farm', NULL, 'George Orwell');

INSERT INTO author (id ,name) VALUES
    (1,'Harper Lee'),
    (2,'George Orwell'),
    (3,'F. Scott Fitzgerald'),
    (4,'J.D. Salinger'),
    (5,'J.R.R. Tolkien'),
    (6,'J.K. Rowling');

SELECT book.name AS book_name, book.page, author.name AS author_name
FROM book
INNER JOIN author ON book.author_name = author.name;
