CREATE TABLE customer (
    id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY(id)
)ENGINE = InnoDb;

DELETE FROM customer;

ALTER TABLE customer
    ADD COLUMN email VARCHAR(100),
    ADD COLUMN balance INT DEFAULT 0,
    ADD COLUMN rating DOUBLE DEFAULT 0.0,
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married BOOLEAN DEFAULT false;

DESC customer;

INSERT INTO customer (id,name,email,balance, rating, birth_date, married)
VALUES ('budi', 'Budi', 'budi@gmail.com', 1000000, 90.0, '1999-10-10', true),
 ('hendrik', 'Hendrik', 'hendrik@gmail.com', 500000, 85.5, '1995-07-21', false),
 ('jamal', 'Jamalk', 'jamal@gmail.com', 7500000, 85.5, '1996-06-18', false);

INSERT INTO customer (id,name,email,balance, rating, birth_date, married)
VALUES ('array', 'Array', NULL, 1000000, 90.0, NULL, true);

CREATE TABLE user (
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    PRIMARY KEY (username)
) ENGINE = InnoDB;

select * from user;

INSERT INTO user(username,password) VALUES ('admin', 'admin');

CREATE TABLE comments(
    id INT NOT NULL AUTO_INCREMENT ,
    email VARCHAR(100)  NOT NULL,
    comment TEXT,
    PRIMARY KEY(id)
) ENGINE = InnoDB;

