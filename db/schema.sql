CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(255) primary key);
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
CREATE TABLE sqlite_sequence(name,seq);
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
FROM users
/* user_view(id,last_login,is_active,username,first_name,last_name,email,date_joined) */;
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20250916133209');
