CREATE TABLE tokens (
    id bigserial not null primary key,
    user_id bigint not null,
    token varchar not null,
    refresh_token varchar not null,
    user_agent varchar not null,
    ip varchar not null,
    expires_at timestamp with time zone default null null,
    created_at timestamp with time zone default current_timestamp not null
);

COMMENT ON COLUMN tokens.user_id IS 'Token attachment to user id';
COMMENT ON COLUMN tokens.user_agent IS 'User client UserAgent';
COMMENT ON COLUMN tokens.ip IS 'User client IP adress';
COMMENT ON COLUMN tokens.expires_at IS 'Token expires date';
COMMENT ON COLUMN tokens.created_at IS 'Token created date';
