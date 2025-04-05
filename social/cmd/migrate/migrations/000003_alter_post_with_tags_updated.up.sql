alter table posts
add column tags varchar(100) [];

alter table posts
add column updated_at timestamp with time zone default now();