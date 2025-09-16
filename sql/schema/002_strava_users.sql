-- +goose up
CREATE TABLE STRAVA_USER (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES USERS(id) ON DELETE CASCADE,
    strava_id BIGINT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    username  VARCHAR(100) NOT NULL,
    firstname VARCHAR(100),
    lastname VARCHAR(100),
    city VARCHAR(100),
    state VARCHAR(100),
    country VARCHAR(100),
    sex CHAR(2),
    premuim BOOLEAN,
    weight FLOAT
);

-- +goose down
DROP TABLE STRAVA_USER;
