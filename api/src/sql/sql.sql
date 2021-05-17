CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS publications;

DROP TABLE IF EXISTS followers;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    PASSWORD VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
) ENGINE = INNODB;

CREATE TABLE followers (
    user_id int NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    follower_id int NOT NULL,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, follower_id)
) ENGINE = INNODB;

CREATE TABLE publications (
    id int AUTO_INCREMENT PRIMARY KEY,
    title varchar(50) NOT NULL,
    content varchar(300) NOT NULL,
    author_id int NOT NULL,
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE,
    likes int DEFAULT 0,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
) ENGINE = INNODB;