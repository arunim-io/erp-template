-- migrate:up
CREATE TABLE users (
  id int PRIMARY KEY
);

-- migrate:down
DROP TABLE users;
