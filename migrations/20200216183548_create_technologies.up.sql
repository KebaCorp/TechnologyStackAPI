CREATE TABLE technologies (
  id bigserial not null primary key,
  type_id int not null,
  stage_id int not null,
  title varchar not null,
  is_deprecated boolean not null,
  creator_user_id int null,
  created_at timestamp with time zone default current_timestamp not null,
  updated_at timestamp with time zone default current_timestamp not null
);

CREATE INDEX type_id	on technologies (type_id);
CREATE INDEX stage_id	on technologies (stage_id);
