CREATE TABLE customer (
    id varchar(100) NOT NULL PRIMARY KEY,
    name varchar(100) NOT NULL
);

ALTER TABLE
    customer
ADD
    COLUMN email varchar(100);

ALTER TABLE
    customer
ADD
    COLUMN balance int DEFAULT 0,
ADD
    COLUMN rating double DEFAULT 0.0,
ADD
    COLUMN created_at timestamp DEFAULT CURRENT_TIMESTAMP,
ADD
    COLUMN birth_data DATE;

ALTER TABLE
    customer
ADD
    married boolean DEFAULT false;

DESC customer;

INSERT INTO
    customer (
        id,
        name,
        email,
        balance,
        rating,
        birth_data,
        married
    )
VALUES
    (
        'dani',
        'dani',
        'dani@gmail.com',
        1000000,
        90.0,
        '2003-08-12',
        false
    );

INSERT INTO
    customer (
        id,
        name,
        email,
        balance,
        rating,
        birth_data,
        married
    )
VALUES
    (
        'oukenzeumasio',
        'oukenzuemasio',
        'oukenzuemasio@gmail.com',
        1000000,
        90.0,
        '2004-05-12',
        false
    ),
    (
        'akrida',
        'akrida',
        'akrida@gmail.com',
        1000000,
        90.0,
        '2004-05-18',
        false
    );

SELECT
    id,
    name,
    email,
    balance,
    rating,
    birth_data,
    married,
    created_at
FROM
    customer;

DELETE FROM
    customer
WHERE
    id = 'robbani';

INSERT INTO
    customer (
        id,
        name,
        email,
        balance,
        rating,
        birth_data,
        married
    )
VALUES
    (
        'lofta',
        'lofta',
        NULL,
        1000000,
        90.0,
        NULL,
        TRUE
    );

CREATE TABLE user(
    username varchar(100) NOT NULL PRIMARY KEY,
    PASSWORD varchar(100) NOT NULL
) ENGINE = InnoDB;

DESC user;

SELECT
    *
FROM
    user;

INSERT INTO
    user(username, PASSWORD)
VALUES
    ("akhdan", "akhdanganteng");

CREATE TABLE comment (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL,
    comment TEXT
);

DESC comment;

SELECT
    *
FROM
    comment;

SELECT
    count(*)
FROM
    comment;