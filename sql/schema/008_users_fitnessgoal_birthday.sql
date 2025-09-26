-- +goose up
ALTER TABLE users
ADD COLUMN fitness_goal TEXT,
ADD COLUMN birthday DATE;

-- +goose down
ALTER TABLE users
DROP COLUMN fitness_goal,
DROP COLUMN birthday;