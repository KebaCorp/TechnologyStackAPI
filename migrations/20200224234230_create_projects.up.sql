CREATE TABLE projects (
  id bigserial not null primary key,
  title varchar not null,
  code varchar not null,
  is_active boolean default false not null,
  creator_user_id int null,
  created_at timestamp with time zone default current_timestamp not null,
  updated_at timestamp with time zone default current_timestamp not null
);

CREATE INDEX code on projects (code);
