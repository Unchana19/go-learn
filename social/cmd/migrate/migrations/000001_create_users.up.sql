create extension if not exists citext;

create table if not exists users (
  id bigserial primary key,
  email citext unique not null,
  username varchar(255) unique not null,
  password bytea not null,
  created_at timestamp with time zone default now() not null
);