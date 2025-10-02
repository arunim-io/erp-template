-- migrate:up
CREATE TABLE users (
    id integer PRIMARY KEY AUTOINCREMENT,

    password varchar(128),
    last_login timestamp,
    is_active boolean DEFAULT true,

    username varchar(150) NOT NULL UNIQUE,
    first_name varchar(150),
    last_name varchar(150),
    email varchar(254) NOT NULL UNIQUE,
    date_joined timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE VIEW user_view AS
SELECT
    id,
    last_login,
    is_active,
    username,
    first_name,
    last_name,
    email,
    date_joined
FROM users;

-- migrate:down
DROP TABLE IF EXISTS users;
DROP VIEW IF EXISTS user_view;
