CREATE TABLE stages (
  id bigserial not null primary key,
  title varchar not null,
  is_deleted boolean not null,
  creator_user_id int null,
  created_at timestamp with time zone default current_timestamp not null,
  updated_at timestamp with time zone default current_timestamp not null
);
