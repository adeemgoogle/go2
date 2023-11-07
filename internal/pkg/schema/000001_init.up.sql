CREATE TABLE Author (
                        ID int primary key ,
                        FullName varchar(255),
                        Pseudonym varchar(255),
                        Specialization varchar(255) not null
);

CREATE TABLE Book (
                      ID int primary key not null ,
                      Title varchar(255) not null ,
                      Genre varchar(255) not null ,
                      ISBN int not null,
                      AuthorId int references Author(ID)
);

CREATE TABLE Members(
                        MemberId int primary key,
                        FullName varchar(255) not null

);
CREATE TABLE BorrowedBooks(
                              MemberID int references Members(MemberId),
                              BookId int references Book(ID)
);

