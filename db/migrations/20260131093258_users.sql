-- migrate:up
create table if not exists users (
  id bigserial primary key,

  password_hash varchar(128),
  last_login timestamp,
  is_active boolean default true,

  username varchar(150) not null unique,
  email varchar(254) unique,
  first_name varchar(150),
  last_name varchar(150),
  date_joined timestamp default current_timestamp
);

-- migrate:down
drop table if exists users;
