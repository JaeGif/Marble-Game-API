CREATE TABLE measurements (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  weight FLOAT NOT NULL,
  height FLOAT NOT NULL,
  body_fat FLOAT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
