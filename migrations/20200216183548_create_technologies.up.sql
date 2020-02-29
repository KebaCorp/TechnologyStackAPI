CREATE TABLE technologies (
  id bigserial not null primary key,
  type_id bigint not null,
  stage_id bigint not null,
  title varchar not null,
  description text default '' not null,
  image text default '' not null,
  is_deprecated boolean default false not null,
  creator_user_id bigint null,
  created_at timestamp with time zone default current_timestamp not null,
  updated_at timestamp with time zone default current_timestamp not null,
  is_deleted boolean default false not null
);

CREATE INDEX type_id on technologies (type_id);
CREATE INDEX stage_id	on technologies (stage_id);
