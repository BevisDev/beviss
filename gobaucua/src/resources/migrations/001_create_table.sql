CREATE TABLE users
(
    id    INT IDENTITY PRIMARY KEY,
    username  VARCHAR(100),
    email VARCHAR(100) NOT NULL
);