-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE tweets (
  id INT AUTO_INCREMENT PRIMARY KEY,
  body VARCHAR(300) NOT NULL,
  title VARCHAR(100) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);