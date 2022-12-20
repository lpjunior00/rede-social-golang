CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS posts;

CREATE TABLE users (
    id int auto_increment primary key,
    name varchar(50) not null,
    nickname varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    creationDate timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers (
    userId int not null, FOREIGN KEY (userid) REFERENCES users(id) ON DELETE CASCADE,
    followerId int not null, FOREIGN KEY (followerId) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (userId, followerId)
) ENGINE=INNODB;

CREATE TABLE posts (
    id int auto_increment primary key,
    title varchar(100) not null,
    content varchar(4000) not null,
    authorId int not null, FOREIGN KEY (authorId) REFERENCES users(id) ON DELETE CASCADE,
    likes int not null default(0),
    creationDate timestamp default current_timestamp()
) ENGINE=INNODB

GRANT ALL PRIVILEGES ON devbook.* TO 'golang'@'localhost';