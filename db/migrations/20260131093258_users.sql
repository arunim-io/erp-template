-- migrate:up
CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  username VARCHAR(150) NOT NULL UNIQUE,
  password_hash VARCHAR(128)
);

-- migrate:down
DROP TABLE IF EXISTS users;
