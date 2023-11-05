CREATE TABLE Author (
    Id int primary key,
    FullName varchar(255),
    Pseudonym varchar(255),
    Spesialization varchar(255)
);
CReate table Book(
    Book_id int primary key ,
    Title varchar(255),
    Genre varchar(255),
    ISBN int,
    Author_id int references Author(Id)
);
CREATE TABLE Members(
    Id int primary key,
    FullName varchar(255)
);

CREATE TABLE BorrowedBooks(
    Member_id int references Members(Id),
    Book_id int references Book(book_id)
);

INSERT INTO author (id, fullname,
                    spesialization, pseudonym) values (1, 'Adem', 'Dem', 'love');

INSERT INTO members(id, fullname) values (1, 'Nurbek');
INSERT INTO book(book_id, title, genre, isbn, author_id) values (1, 'ogenri', 'horror', 123, 1)