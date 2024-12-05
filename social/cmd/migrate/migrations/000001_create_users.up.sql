CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users(
  id bigserial PRIMARY KEY,
  user_name varchar(255) NOT NULL,
  email citext UNIQUE NOT NULL,
  username varchar(255) UNIQUE NOT NULL,
  password bytea NOT NULL,
  create_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
)