create table if not exists posts (
  id bigserial primary key,
  title varchar(255) not null,
  content text not null,
  user_id bigint not null references users(id),
  created_at timestamp with time zone default now() not null
);
