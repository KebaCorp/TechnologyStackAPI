CREATE TABLE tokens (
    id bigserial not null primary key,
    user_id bigint not null,
    user_agent varchar not null,
    ip varchar not null,
    expires_at timestamp with time zone default null null,
    updated_at timestamp with time zone default current_timestamp not null
);
