-- DDL.sql
CREATE TABLE bank_account (
    id SERIAL PRIMARY KEY,
    owner VARCHAR(255) NOT NULL,
    money INT NOT NULL
    );

INSERT INTO bank_account (owner, money) VALUES
('user1', 1000),
('user2', 1500),
('user3', 2000);
