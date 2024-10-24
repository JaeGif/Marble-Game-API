CREATE TABLE scores {
    id SERIAL PRIMARY KEY,
    user_name STRING NOT NULL,
    score INTEGER NOT NULL
    created_at TIMESTAMP NOT NULL
}