CREATE TABLE stages (
  id bigserial not null primary key,
  title varchar not null,
  creator_user_id bigint null,
  created_at timestamp with time zone default current_timestamp not null,
  updated_at timestamp with time zone default current_timestamp not null,
  is_deleted boolean default false not null
);
