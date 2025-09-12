-- +goose up
CREATE TABLE USERS (
    id uuid PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    hashed_password varchar default 'unset'
);

-- +goose down
DROP TABLE USERS;

