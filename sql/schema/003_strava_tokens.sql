-- +goose up
CREATE TABLE STRAVA_TOKENS (
    id UUID PRIMARY KEY,
    strava_user_id UUID NOT NULL REFERENCES STRAVA_USER(id) ON DELETE CASCADE,
    access_token VARCHAR(255) NOT NULL,
    token_type VARCHAR(50) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    refresh_token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose down
DROP TABLE STRAVA_ACCESS_TOKENS;    


