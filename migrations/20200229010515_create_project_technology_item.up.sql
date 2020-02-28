CREATE TABLE project_technology_item (
    id bigserial not null primary key,
    project_id bigint not null,
    technology_item_id bigint not null,
    started_at timestamp with time zone default current_timestamp not null,
    ended_at timestamp with time zone default current_timestamp not null,
    creator_user_id bigint null,
    created_at timestamp with time zone default current_timestamp not null,
    updated_at timestamp with time zone default current_timestamp not null
);
