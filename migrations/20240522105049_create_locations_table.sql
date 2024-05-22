-- +goose Up

CREATE TABLE locations (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(300) NOT NULL,
  state INT NOT NULL,
  FOREIGN KEY (state) REFERENCES state(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down

DROP TABLE locations;