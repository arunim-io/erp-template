CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(255) primary key);
CREATE TABLE users (
  id int PRIMARY KEY
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20250916133209');
