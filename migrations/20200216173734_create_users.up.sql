CREATE TABLE users (
  id bigserial not null primary key,
  email varchar not null unique,
  username varchar not null unique,
  first_name varchar not null,
  last_name varchar not null,
  middle_name varchar not null,
  is_active boolean default true not null,
  encrypted_password varchar not null,
  creator_user_id bigint null,
  created_at timestamp with time zone default current_timestamp not null,
  updated_at timestamp with time zone default current_timestamp not null
);
